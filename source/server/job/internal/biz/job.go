package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/tx7do/kratos-transport/broker"
	"job/api/v1/connector"
	"job/api/v1/online"
	"job/api/v1/relationship"
	"job/api/v1/universal"
	"job/api/v1/user"
	"job/internal/components/redis"
	"job/internal/util/mq_kafka"
	"strconv"
)

type JobBiz struct {
	helper          *log.Helper
	redisCli        *redis.Redis
	onlineCli       online.OnlineClient
	relationshipCli relationship.RelationShipClient
	userCli         user.UserClient
}

const FriendRequestPrefix = "friend_request:uid:"

func NewConnectorClient(url string) connector.ConnectorServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(url),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	return connector.NewConnectorServiceClient(conn)
}

func (b *JobBiz) HandleFriendRequest(ctx context.Context, event broker.Event, msg *mq_kafka.FriendRequestMessage) {
	if msg.RequestId == 0 || msg.UserId == 0 {
		b.helper.Errorf("HandleFriendRequest: %s", "request_id or user_id is empty")
	}
	b.helper.Infof("HandleFriendRequest: %s %s %v", msg.RequestId, msg.UserId, msg.PublishAt)
	if err := event.Ack(); err != nil {
		b.helper.Errorf("ack err: %s", err)
	}
	device, err := b.onlineCli.GetOnlineDevice(ctx, &online.GetOnlineDeviceRequest{Uid: msg.UserId})
	if err != nil {
		b.helper.Errorf("GetOnlineDevice err: %s", err)
		return
	}
	b.helper.Infof("device: %v", device)
	if device == nil || device.DeviceUrl == "" {
		err := b.redisCli.SAdd(FriendRequestPrefix+strconv.FormatInt(msg.UserId, 10), msg.RequestId)
		if err != nil {
			b.helper.Errorf("redis set err: %s", err)
			return
		}
	} else {
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
		connectorClient := NewConnectorClient(device.DeviceUrl)
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

}

func (b *JobBiz) HandleConnectInit(ctx context.Context, event broker.Event, msg *mq_kafka.ConnectInitMessage) {
	uid := msg.UserId
	b.helper.Infof("HandleConnectInit: %s", uid)
	//检查uid是否为空
	if uid == 0 {
		b.helper.Errorf("HandleConnectInit: %s", "user_id is empty")
	}
	//ack消息
	if err := event.Ack(); err != nil {
		b.helper.Errorf("ack err: %s", err)
	}
	//从redis中获取未推送的好友请求
	requestIds, err := b.redisCli.SMembers(FriendRequestPrefix + strconv.FormatInt(uid, 10))
	if err != nil {
		b.helper.Errorf("redis get err: %s", err)
		return
	}
	//如果没有未推送的好友请求，直接返回
	if len(requestIds) == 0 {
		return
	}
	//将requestIds转换成int64数组
	var requestIdsInt64 []int64
	for _, requestId := range requestIds {
		requestIdInt64, err := strconv.ParseInt(requestId, 10, 64)
		if err != nil {
			b.helper.Errorf("strconv.ParseInt err: %s", err)
			continue
		}
		requestIdsInt64 = append(requestIdsInt64, requestIdInt64)
	}
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
	// 获取在线设备所在的连接器地址
	device, err := b.onlineCli.GetOnlineDevice(ctx, &online.GetOnlineDeviceRequest{Uid: msg.UserId})
	if err != nil {
		b.helper.Errorf("GetOnlineDevice err: %s", err)
		return
	}
	b.helper.Infof("device: %v", device)
	//	构建连接器客户端
	connectorClient := NewConnectorClient(device.DeviceUrl)
	//推送好友请求
	_, err = connectorClient.PushFriendRequests(ctx, &connector.PushFriendRequestsRequest{
		Uid:      uid,
		Requests: requests.FriendRequests,
	})
	//删除redis中的好友请求
	err = b.redisCli.Del(FriendRequestPrefix + strconv.FormatInt(uid, 10))
	if err != nil {
		b.helper.Errorf("redis del err: %s", err)
		return
	}
}

func NewJobBiz(helper *log.Helper, redisCli *redis.Redis, onlineCli online.OnlineClient, relationshipCli relationship.RelationShipClient, userCli user.UserClient) *JobBiz {
	return &JobBiz{helper: helper, redisCli: redisCli, onlineCli: onlineCli, relationshipCli: relationshipCli, userCli: userCli}
}
