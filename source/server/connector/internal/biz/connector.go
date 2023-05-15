package biz

import (
	"connector/api/v1/connector"
	"connector/api/v1/online"
	"connector/api/v1/universal"
	"connector/internal/components/broker"
	"connector/internal/components/cache"
	"connector/internal/components/redis"
	"connector/internal/conf"
	"connector/internal/data/orm/model"
	"connector/pkg"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
)

const (
	WsPrefix = "ws:uid:"
)

type ConnectorServiceBiz struct {
	repo         ConnectorServiceRepo
	helper       *log.Helper
	upgrader     *websocket.Upgrader
	cache        *cache.ConnectionCache
	onlineClient online.OnlineClient
	serverCf     *conf.Server
	kafkaBroker  *broker.KafkaBroker
	mqCf         *conf.MessageQueue
	redisCli     *redis.Redis
}

func NewConnectorServiceBiz(connectorServiceRepo ConnectorServiceRepo, helper *log.Helper, onlineClient online.OnlineClient, upgrader *websocket.Upgrader, cache *cache.ConnectionCache, cf *conf.Bootstrap, kafkaBroker *broker.KafkaBroker, redisCli *redis.Redis) *ConnectorServiceBiz {
	return &ConnectorServiceBiz{repo: connectorServiceRepo, helper: helper, onlineClient: onlineClient, upgrader: upgrader, cache: cache, serverCf: cf.Server, kafkaBroker: kafkaBroker, mqCf: cf.MessageQueue, redisCli: redisCli}
}

type ConnectorServiceRepo interface {
	FindUserByPhone(ctx context.Context, phone string) (*model.User, error)
	FindUserByAccountId(ctx context.Context, accountId string) (*model.User, error)
	UpdateLoginStatus(ctx context.Context, uid int64, status int) error
	FindAddressByCityId(ctx context.Context, cityId string) (*connector.UserAddress, error)
}

func (c *ConnectorServiceBiz) Login(ctx context.Context, t int64, username, password string) (*model.User, *connector.UserAddress, error) {
	var user *model.User
	var err error

	if t == pkg.Phone {
		user, err = c.repo.FindUserByPhone(ctx, username)
		if err != nil {
			return nil, nil, err
		}
		v, err := c.redisCli.Get("smscode:" + username)
		if err != nil {
			return nil, nil, pkg.InvalidArgumentError("验证码已失效")
		}
		var smsCode *model.SmsCode
		err = json.Unmarshal([]byte(v), &smsCode)
		if err != nil {
			return nil, nil, pkg.InternalError("解析json失败")
		}
		if smsCode.Code != password {
			return nil, nil, pkg.InvalidArgumentError("验证码错误")
		}
	} else if t == pkg.Account {
		user, err = c.repo.FindUserByPhone(ctx, username)
		if err != nil {
			user, err = c.repo.FindUserByAccountId(ctx, username)
			if err != nil {
				return nil, nil, err
			}
		}
		c.helper.Infof("user: %v", user)
		if err != nil {
			return nil, nil, err
		}
		if user == nil {
			return nil, nil, pkg.InvalidArgumentError("user not found")
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return nil, nil, pkg.InvalidArgumentError("password error")
		}
	}
	var userAddress = &connector.UserAddress{}
	if user.CityID != "" {
		userAddress, err = c.repo.FindAddressByCityId(ctx, user.CityID)
		if err != nil {
			return nil, nil, err
		}
	}
	err = c.repo.UpdateLoginStatus(ctx, user.UID, pkg.Login)
	if err != nil {
		return nil, nil, err
	}
	return user, userAddress, nil
}
func (c *ConnectorServiceBiz) Logout(ctx context.Context, uid string) error {
	err := pkg.ContextErr(ctx)
	if err != nil {
		return err
	}
	err = c.repo.UpdateLoginStatus(ctx, pkg.ParseInt64(uid), pkg.NotLogin)
	if err != nil {
		return err
	}
	conn, ok := c.cache.GetConn(WsPrefix + uid)
	if !ok {
		return nil
	}
	err = conn.Close()
	if err != nil {
		c.helper.Errorf("Logout error: %v", err)
		return err
	}
	c.cache.RemoveConn(WsPrefix + uid)
	_, err = c.onlineClient.UnregisterDevice(ctx, &online.UnregisterDeviceRequest{
		Uid: pkg.ParseInt64(uid),
	})
	if err != nil {
		c.helper.Errorf("unregisterDevice error: %v", err)
	}
	return nil
}
func (c *ConnectorServiceBiz) Serve(conn *websocket.Conn, uid string) {
	// 向mq发送消息，做连接后的初始化工作，比如说推送离线消息等
	msg := model.ConnectInitMessage{
		UserId: pkg.ParseInt64(uid),
	}
	err := c.kafkaBroker.Publish(c.mqCf.ConnectTopic, msg)
	if err != nil {
		c.helper.Errorf("Publish error: %v", err)
	}
	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err) {
				c.cache.RemoveConn(WsPrefix + uid)
				return
			} else {
				err := conn.Close()
				if err != nil {
					c.helper.Errorf("close error: %v", err)
				}
				c.cache.RemoveConn(WsPrefix + uid)
				return
			}
		}
		msg := &model.UniversalMessage{}
		json.Unmarshal(bytes, msg)
		if msg.T == "ping" {
			c.helper.Infof("HeartCheck: %v", msg)
			if err != nil {
				c.helper.Errorf("HeartCheck error: %v", err)
				c.cache.RemoveConn(WsPrefix + uid)
				return
			}
			message := &model.UniversalMessage{
				T:    "pong",
				Data: "pond",
			}
			bytes, err = json.Marshal(message)
			err := conn.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				c.cache.RemoveConn(WsPrefix + uid)
				return
			}
		}
	}

}

func (c *ConnectorServiceBiz) Connect(ctx *gin.Context, uid string) error {
	c.helper.Infof("Connect: %v", uid)
	conn, ok := c.cache.GetConn(WsPrefix + uid) //判断是否已经存在连接
	if ok {
		err := conn.Close() //关闭连接
		if err != nil {
			c.helper.Errorf("close error: %v", err)
			return err
		}
		return nil
	}
	ws, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil) //重新升级为websocket连接
	if err != nil {
		c.helper.Errorf("Connect error: %v", err)
		return err
	}
	c.cache.SetConn(WsPrefix+uid, ws) //将连接放入缓存
	c.helper.Infof("Connect success: %v", uid)
	//拿到grpc的地址
	addr := c.serverCf.Grpc.Addr
	_, err = c.onlineClient.RegisterDevice(ctx, &online.RegisterDeviceRequest{
		Uid:       pkg.ParseInt64(uid),
		DeviceUrl: addr,
	}) //注册设备
	if err != nil {
		c.helper.Errorf("RegisterDevice error: %v", err)
		c.cache.Remove(WsPrefix + uid)
		return err
	}
	msg := &model.UniversalMessage{
		T:    "pong",
		Data: "pong",
	}
	bytes, err := json.Marshal(msg)
	err = ws.WriteMessage(websocket.TextMessage, bytes) //发送心跳
	if err != nil {
		c.helper.Errorf("Connect error: %v", err)
		return err
	}
	go c.Serve(ws, uid) //开启协程监听心跳
	return nil
}

func (c *ConnectorServiceBiz) PushFriendRequests(ctx context.Context, uid string, requests []*universal.FriendRequest) error {
	err := pkg.ContextErr(ctx)
	if err != nil {
		return err
	}
	conn, ok := c.cache.GetConn(WsPrefix + uid)
	var mReqs []*model.FriendRequest
	for _, req := range requests {
		mReq := &model.FriendRequest{
			RequestId:   pkg.FormatInt(req.RequestId),
			RequesterId: pkg.FormatInt(req.RequesterId),
			ReceiverId:  pkg.FormatInt(req.ReceiverId),
			CreateTime:  req.CreateTime,
			UpdateTime:  req.UpdateTime,
			Avatar:      req.Avatar,
			NickName:    req.NickName,
			Desc:        req.Desc,
			Status:      req.Status,
		}
		mReqs = append(mReqs, mReq)
	}
	if !ok {
		return nil
	}
	msg := &model.UniversalMessage{
		T:    "friendRequest",
		Data: mReqs,
	}
	bytes, err := json.Marshal(msg)
	c.helper.Infof("PushFriendRequests: %v", msg)
	err = conn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		c.helper.Errorf("PushFriendRequests error: %v", err)
		return err
	}
	return nil
}
