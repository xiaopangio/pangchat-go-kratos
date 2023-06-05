package biz

import (
	"connector/api/v1/connector"
	"connector/api/v1/message"
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
	"time"
)

const (
	WsPrefix = "ws:uid:"
)

type ConnectorServiceBiz struct {
	repo          ConnectorServiceRepo
	helper        *log.Helper
	upgrader      *websocket.Upgrader
	cache         *cache.ConnectionCache
	onlineClient  online.OnlineClient
	messageClient message.MessageServiceClient
	serverCf      *conf.Server
	kafkaBroker   *broker.KafkaBroker
	mqCf          *conf.MessageQueue
	redisCli      *redis.Redis
}

func NewConnectorServiceBiz(connectorServiceRepo ConnectorServiceRepo, messageClient message.MessageServiceClient, helper *log.Helper, onlineClient online.OnlineClient, upgrader *websocket.Upgrader, cache *cache.ConnectionCache, cf *conf.Bootstrap, kafkaBroker *broker.KafkaBroker, redisCli *redis.Redis) *ConnectorServiceBiz {
	return &ConnectorServiceBiz{repo: connectorServiceRepo, helper: helper, messageClient: messageClient, onlineClient: onlineClient, upgrader: upgrader, cache: cache, serverCf: cf.Server, kafkaBroker: kafkaBroker, mqCf: cf.MessageQueue, redisCli: redisCli}
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
	c.helper.Info("unregisterDevice")
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
	//维持心跳
	go c.HeartCheck(conn, uid)
}

func (c *ConnectorServiceBiz) Connect(ctx *gin.Context, uid string) error {
	c.helper.Infof("Connect: %v", uid)
	conn, ok := c.cache.GetConn(WsPrefix + uid) //判断是否已经存在连接
	if ok {
		c.CloseConn(conn, uid)
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

func (c *ConnectorServiceBiz) HeartCheck(conn *websocket.Conn, uid string) {
	for {
		//设置读取超时时间
		err := conn.SetReadDeadline(time.Now().Add(time.Second * 11))
		if err != nil {
			c.helper.Errorf("HeartCheck error: %v", err)
			c.CloseConn(conn, uid)
			return
		}
		//读取消息
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			//判断连接是否关闭
			if websocket.IsCloseError(err) {
				//关闭连接
				c.CloseConn(conn, uid)
				return
			} else {
				//关闭连接
				c.CloseConn(conn, uid)
				return
			}
		}
		//解析消息
		msg := &model.UniversalMessage{}
		err = json.Unmarshal(bytes, msg)
		if err != nil {
			c.helper.Errorf("HeartCheck error: %v", err)
			//关闭连接
			c.CloseConn(conn, uid)
			return
		}
		//判断消息类型
		if msg.T == "ping" {
			//回复心跳
			m := &model.UniversalMessage{
				T:    "pong",
				Data: "pong",
			}
			bytes, err = json.Marshal(m)
			err := conn.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				c.CloseConn(conn, uid)
				return
			}
		} else if msg.T == "single_message" {
			c.ProcessSingleMessage(conn, uid, msg)
		}
	}
}
func (c *ConnectorServiceBiz) ProcessSingleMessage(conn *websocket.Conn, uid string, msg *model.UniversalMessage) {
	c.helper.Infof("msg: %v", msg)
	res := msg.Data.(map[string]interface{})
	m := &universal.Message{
		MessageId:  res["message_id"].(string),
		Type:       pkg.ParseInt64(res["type"].(string)),
		Content:    res["content"].(string),
		SenderId:   res["sender_id"].(string),
		ReceiverId: res["receiver_id"].(string),
		SendAt:     res["send_at"].(string),
	}
	c.helper.Infof("ProcessSingleMessage: %v", m)
	if _, err := c.messageClient.DealSingleMessage(context.Background(), &message.DealSingleMessageRequest{
		Message: m,
	}); err != nil {
		//消息处理失败
		reply := &model.UniversalMessage{
			T:    "single_message_reply_error",
			Data: err.Error(),
		}
		bytes, err := json.Marshal(reply)
		err = conn.WriteMessage(websocket.TextMessage, bytes)
		if err != nil {
			c.CloseConn(conn, uid)
			return
		}
		return
	}
}
func (c *ConnectorServiceBiz) CloseConn(conn *websocket.Conn, uid string) {
	err := conn.Close()
	if err != nil {
		c.helper.Errorf("close error: %v", err)
	}
	c.cache.RemoveConn(WsPrefix + uid)
	_, err = c.onlineClient.UnregisterDevice(context.Background(), &online.UnregisterDeviceRequest{
		Uid: pkg.ParseInt64(uid),
	})
	if err != nil {
		c.helper.Errorf("UnregisterDevice error: %v", err)
	}
}

func (c *ConnectorServiceBiz) PushFriend(ctx context.Context, uid string, friends []*universal.Friend) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	conn, ok := c.cache.GetConn(WsPrefix + uid)
	if !ok {
		return pkg.InternalError("用户不在线")
	}
	msg := &model.UniversalMessage{
		T:    "friend",
		Data: friends,
	}
	if bytes, err := json.Marshal(msg); err != nil {
		return err
	} else {
		if err = conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
			c.helper.Errorf("PushFriend error: %v", err)
			return err
		}
		return nil
	}
}

func (c *ConnectorServiceBiz) PushMessage(ctx context.Context, uid string, m *universal.Message) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	conn, ok := c.cache.GetConn(WsPrefix + uid)
	if !ok {
		return pkg.InternalError("用户不在线")
	}
	msg := &model.UniversalMessage{
		T:    "message",
		Data: m,
	}
	if bytes, err := json.Marshal(msg); err != nil {
		c.helper.Errorf("marshal error: %v", err)
		return err
	} else {
		if err = conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
			c.helper.Errorf("PushMessage error: %v", err)
			return err
		}
	}

	return nil
}

func (c *ConnectorServiceBiz) ReplyMessage(ctx context.Context, uid string, m *universal.Message) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil
	}
	conn, ok := c.cache.GetConn(WsPrefix + uid)
	if !ok {
		return pkg.InternalError("用户不在线")
	}
	m.Content = ""
	m.Type = 0
	msg := &model.UniversalMessage{
		T:    "single_message_reply",
		Data: m,
	}
	if bytes, err := json.Marshal(msg); err != nil {
		c.helper.Errorf("marshal error: %v", err)
		return err
	} else {
		if err = conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
			c.helper.Errorf("ReplyMessage error: %v", err)
			return err
		}
		return nil
	}
}

func (c *ConnectorServiceBiz) PushUnreadMessageList(ctx context.Context, uid string, list []*universal.UnreadMessageInfo) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	conn, ok := c.cache.GetConn(WsPrefix + uid)
	if !ok {
		return pkg.InternalError("用户不在线")
	}
	data := &model.UnreadMessageResponse{
		List: list,
	}
	msg := &model.UniversalMessage{
		T:    "unread_message_list",
		Data: data,
	}
	if bytes, err := json.Marshal(msg); err != nil {
		c.helper.Errorf("marshal error: %v", err)
		return err
	} else {
		if err = conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
			c.helper.Errorf("PushUnreadMessageList error: %v", err)
			return err
		}
		return nil
	}
}
