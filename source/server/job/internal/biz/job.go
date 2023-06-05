package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/tx7do/kratos-transport/broker"
	"job/api/v1/connector"
	"job/api/v1/message"
	"job/api/v1/online"
	"job/api/v1/relationship"
	"job/api/v1/universal"
	"job/api/v1/user"
	"job/internal/components/redis"
	"job/internal/util/mq_kafka"
	"job/pkg"
)

// FriendRequestPrefix 好友请求前缀
const FriendRequestPrefix = "friend_request:uid:"

// FriendPrefix 好友前缀
const FriendPrefix = "friend:uid:"

// JobBiz job业务逻辑
type JobBiz struct {
	helper          *log.Helper
	redisCli        *redis.Redis
	onlineCli       online.OnlineClient
	relationshipCli relationship.RelationShipClient
	userCli         user.UserClient
	messageCli      message.MessageServiceClient
}

// NewConnectorClient  返回一个grpc连接
func NewConnectorClient(url string) (connector.ConnectorServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(url),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.Error("dial err: %v", err)
		return nil, err
	}
	return connector.NewConnectorServiceClient(conn), nil
}

// StoreFriendRequestMessage redis存储好友请求
func (b *JobBiz) StoreFriendRequestMessage(uid, requestId int64) error {
	return b.redisCli.SAdd(FriendRequestPrefix+pkg.FormatInt(uid), requestId)
}

// GetDevice 根据uid获取用户在线设备
func (b *JobBiz) GetDevice(ctx context.Context, uid int64) (*online.GetOnlineDeviceResponse, error) {
	device, err := b.onlineCli.GetOnlineDevice(ctx, &online.GetOnlineDeviceRequest{Uid: uid})
	if err != nil {
		b.helper.Errorf("GetOnlineDevice err: %s", err)
		return nil, err
	}
	return device, nil
}

// GetInt64Ids 将string数组转换成int64数组
func (b *JobBiz) GetInt64Ids(ids []string) []int64 {
	var res []int64
	for _, id := range ids {
		idInt64 := pkg.ParseInt64(id)
		res = append(res, idInt64)
	}
	return res
}

// StoreFriendMessages 将未推送的好友存入redis
func (b *JobBiz) StoreFriendMessages(uid int64, friendIds ...int64) error {
	return b.redisCli.SAdd(FriendPrefix+pkg.FormatInt(uid), friendIds)
}

// HandleFriendRequest 处理好友请求消息
func (b *JobBiz) HandleFriendRequest(ctx context.Context, event broker.Event, msg *mq_kafka.FriendRequestMessage) {
	if msg.RequestId == 0 || msg.UserId == 0 {
		b.helper.Errorf("HandleFriendRequest: %s", "request_id or user_id is empty")
	}
	b.helper.Infof("HandleFriendRequest: %s %s %v", msg.RequestId, msg.UserId, msg.PublishAt)
	if err := event.Ack(); err != nil {
		b.helper.Errorf("ack err: %s", err)
	}
	//获取用户在线设备
	device, err := b.GetDevice(ctx, msg.UserId)
	if err != nil {
		//用户不在线,存入redis
		if err = b.StoreFriendRequestMessage(msg.UserId, msg.RequestId); err != nil {
			return
		}
		return
	}
	reply, err := b.relationshipCli.GetFriendRequest(ctx, &relationship.GetFriendRequestRequest{RequestId: msg.RequestId})
	if err != nil {
		b.helper.Errorf("GetFriendRequest err: %s", err)
		return
	}
	b.helper.Info(b.userCli)
	res, err := b.userCli.GetProfiles(ctx, &user.GetProfilesRequest{
		Uids: []int64{
			reply.FriendRequest.RequesterId,
		},
	})
	if err != nil {
		b.helper.Errorf("get profiles error: %v", err)
		return
	}
	var NickName, Avatar string
	if len(res.Profiles) == 1 {
		NickName = res.Profiles[0].NickName
		Avatar = res.Profiles[0].Avatar
	}
	connectorClient, err := NewConnectorClient(device.DeviceUrl)
	if err != nil {
		b.helper.Errorf("NewConnectorClient err: %s", err)
		return
	}
	_, err = connectorClient.PushFriendRequests(ctx, &connector.PushFriendRequestsRequest{
		Uid: msg.UserId,
		Requests: []*universal.FriendRequest{
			{
				RequestId:   reply.FriendRequest.RequestId,
				RequesterId: reply.FriendRequest.RequesterId,
				ReceiverId:  reply.FriendRequest.ReceiverId,
				Desc:        reply.FriendRequest.Desc,
				Status:      reply.FriendRequest.Status,
				CreateTime:  reply.FriendRequest.CreateTime,
				UpdateTime:  reply.FriendRequest.UpdateTime,
				NickName:    NickName,
				Avatar:      Avatar,
			},
		},
	})
	if err != nil {
		b.helper.Errorf("SendFriendRequest err: %s", err)
		return
	}

}

// HandleConnectInit 处理用户连接消息
func (b *JobBiz) HandleConnectInit(ctx context.Context, event broker.Event, msg *mq_kafka.ConnectInitMessage) {
	//ack消息
	if err := event.Ack(); err != nil {
		b.helper.Errorf("ack err: %s", err)
	}
	// 获取在线设备所在的连接器地址
	device, err := b.GetDevice(ctx, msg.UserId)
	if err != nil {
		return
	}
	uid := msg.UserId
	b.helper.Infof("HandleConnectInit: %v", uid)
	//检查uid是否为空
	if uid == 0 {
		b.helper.Errorf("HandleConnectInit: %s", "user_id is empty")
	}
	//	构建连接器客户端
	connectorClient, err := NewConnectorClient(device.DeviceUrl)
	if err != nil {
		b.helper.Errorf("HandleConnectInit: %s", err)
		return
	}
	//	初始化好友请求
	b.InitFriendRequest(ctx, pkg.FormatInt(uid), connectorClient)
	//初始化好友
	b.InitFriend(ctx, pkg.FormatInt(uid), connectorClient)
	//初始化未读消息
	b.InitMessage(ctx, pkg.FormatInt(uid), connectorClient)
}

// InitFriendRequest 用户连接成功后，用来推送未推送的好友请求以及处理完毕后的好友请求
func (b *JobBiz) InitFriendRequest(ctx context.Context, uid string, conn connector.ConnectorServiceClient) {
	//从redis中获取未推送的好友请求
	requestIds, err := b.redisCli.SMembers(FriendRequestPrefix + uid)
	b.helper.Infof("requestIds: %v", requestIds)
	if err != nil {
		b.helper.Errorf("redis get err: %s", err)
		return
	}
	//如果没有未推送的好友请求，直接返回
	if len(requestIds) == 0 {
		return
	}
	//将requestIds转换成int64数组
	requestIdsInt64 := b.GetInt64Ids(requestIds)
	//获取好友请求详情
	requests, err := b.relationshipCli.GetFriendRequests(ctx, &relationship.GetFriendRequestsRequest{
		RequestIds: requestIdsInt64,
	})
	if err != nil {
		b.helper.Errorf("GetFriendRequests err: %s", err)
		return
	}
	b.helper.Infof("requests: %v", requests)
	if requests == nil {
		return
	}
	//获取请求方的部分信息
	var uids []int64
	for _, request := range requests.FriendRequests {
		uids = append(uids, request.RequesterId)
	}
	reply, err := b.userCli.GetProfiles(ctx, &user.GetProfilesRequest{
		Uids: uids,
	})
	if err != nil {
		b.helper.Errorf("GetProfiles err: %s", err)
		return
	}
	m := make(map[int64]*user.ShortProfile, len(reply.Profiles))
	for _, profile := range reply.Profiles {
		m[profile.Uid] = profile
	}
	//拼接参数
	for i := 0; i < len(requests.FriendRequests); i++ {
		requests.FriendRequests[i].Avatar = m[requests.FriendRequests[i].RequesterId].Avatar
		requests.FriendRequests[i].NickName = m[requests.FriendRequests[i].RequesterId].NickName
	}
	//推送好友请求
	_, err = conn.PushFriendRequests(ctx, &connector.PushFriendRequestsRequest{
		Uid:      pkg.ParseInt64(uid),
		Requests: requests.FriendRequests,
	})
	//删除redis中的好友请求
	err = b.redisCli.Del(FriendRequestPrefix + uid)
	if err != nil {
		b.helper.Errorf("redis del err: %s", err)
		return
	}
}

// InitFriend 用户连接成功后，用来推送新增好友
func (b *JobBiz) InitFriend(ctx context.Context, uid string, conn connector.ConnectorServiceClient) {
	friendIds, err := b.redisCli.SMembers(FriendPrefix + uid)
	if len(friendIds) == 0 {
		return
	}
	if err != nil {
		b.helper.Errorf("redis get err: %s", err)
		return
	}
	ids := b.GetInt64Ids(friendIds)
	res, err := b.relationshipCli.GetFriendsByIDS(ctx, &relationship.GetFriendsByIDSRequest{
		FriendIds: ids,
	})
	if err != nil {
		return
	}
	if _, err = conn.PushFriend(ctx, &connector.PushFriendRequest{
		Uid:     pkg.ParseInt64(uid),
		Friends: res.Friends,
	}); err != nil {
		if err = b.StoreFriendMessages(pkg.ParseInt64(uid), ids...); err != nil {
			return
		}
		return
	}
}

// InitMessage 用户连接成功后，用来推送未读消息列表
func (b *JobBiz) InitMessage(ctx context.Context, uid string, conn connector.ConnectorServiceClient) {
	b.helper.Info("InitMessage")
	if reply, err := b.messageCli.GetLatestUnreadMessageList(ctx, &message.GetLatestUnreadMessageListRequest{
		Uid: uid,
	}); err != nil {
		return
	} else {
		if len(reply.List) == 0 {
			return
		}
		if _, err = conn.PushUnreadMessageList(ctx, &connector.PushUnreadMessageListRequest{
			Uid:  pkg.ParseInt64(uid),
			List: reply.List,
		}); err != nil {
			return
		}
		// 更新redis中的ack
		var ackInfos []*message.AckMessageInfo
		for _, info := range reply.List {
			ackInfos = append(ackInfos, &message.AckMessageInfo{
				MessageId: info.LatestMessage.MessageId,
				SenderId:  info.LatestMessage.SenderId,
			})
		}
		if _, err = b.messageCli.UpdateAckMessages(ctx, &message.UpdateAckMessagesRequest{
			ReceiverId: uid,
			List:       ackInfos,
		}); err != nil {
			return
		}
	}
}

// HandleFriend 处理好友消息
func (b *JobBiz) HandleFriend(ctx context.Context, event broker.Event, msg *mq_kafka.FriendMessage) {
	//ack消息
	if err := event.Ack(); err != nil {
		b.helper.Errorf("ack err: %s", err)
	}
	device, err := b.GetDevice(ctx, msg.UserId)
	if err != nil {
		if err = b.StoreFriendMessages(msg.UserId, msg.FriendId); err != nil {
			return
		}
		return
	}
	uid := msg.UserId
	friendId := msg.FriendId
	res, err := b.relationshipCli.GetOneFriend(ctx, &relationship.GetOneFriendRequest{
		UserId:   uid,
		FriendId: friendId,
	})
	if err != nil {
		b.helper.Errorf("GetOneFriend err: %s", err)
		return
	}
	//构建连接器客户端
	connectorClient, err := NewConnectorClient(device.DeviceUrl)
	if err != nil {
		b.helper.Errorf("NewConnectorClient err: %s", err)
		return
	}
	//推送朋友信息
	if _, err = connectorClient.PushFriend(ctx, &connector.PushFriendRequest{
		Uid: uid,
		Friends: []*universal.Friend{
			res.Friend,
		},
	}); err != nil {
		b.helper.Errorf("PushFriend err: %s", err)
		if err = b.StoreFriendMessages(uid, friendId); err != nil {
			return
		}
		return
	}
}

// HandleMessage 处理消息消息
func (b *JobBiz) HandleMessage(ctx context.Context, event broker.Event, msg *universal.Message) {
	//ack消息
	if err := event.Ack(); err != nil {
		b.helper.Errorf("ack err: %s", err)
	}
	senderDevice, err := b.GetDevice(ctx, pkg.ParseInt64(msg.SenderId))
	if err != nil {
		return
	}
	//构建连接器客户端
	senderConnClient, err := NewConnectorClient(senderDevice.DeviceUrl)
	if err != nil {
		b.helper.Errorf("NewConnectorClient err: %s", err)
		return
	}
	//推送消息
	if _, err = senderConnClient.ReplyMessage(ctx, &connector.ReplyMessageRequest{
		Uid:     pkg.ParseInt64(msg.SenderId),
		Message: msg,
	}); err != nil {
		b.helper.Errorf("PushMessage err: %s", err)
		return
	}
	receiverDevice, err := b.GetDevice(ctx, pkg.ParseInt64(msg.ReceiverId))
	if err != nil {
		return
	}
	receiverConnClient, err := NewConnectorClient(receiverDevice.DeviceUrl)
	if err != nil {
		b.helper.Errorf("NewConnectorClient err: %s", err)
		return
	}
	if _, err = receiverConnClient.PushMessage(ctx, &connector.PushMessageRequest{
		Uid:     pkg.ParseInt64(msg.ReceiverId),
		Message: msg,
	}); err != nil {
		b.helper.Errorf("PushMessage err: %s", err)
		return
	}
	//更新消息状态
	if _, err = b.messageCli.UpdateAckMessage(ctx, &message.UpdateAckMessageRequest{
		MessageId:  msg.MessageId,
		SenderId:   msg.SenderId,
		ReceiverId: msg.ReceiverId,
	}); err != nil {
		return
	}
}

// NewJobBiz 创建job业务实例
func NewJobBiz(helper *log.Helper, redisCli *redis.Redis, messageCli message.MessageServiceClient, onlineCli online.OnlineClient, relationshipCli relationship.RelationShipClient, userCli user.UserClient) *JobBiz {
	return &JobBiz{helper: helper, redisCli: redisCli, messageCli: messageCli, onlineCli: onlineCli, relationshipCli: relationshipCli, userCli: userCli}
}
