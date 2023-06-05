package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"job/api/v1/universal"
	"job/internal/biz"
	"job/internal/util/mq_kafka"
)

type JobService struct {
	helper *log.Helper
	biz    *biz.JobBiz
}

func NewJobService(helper *log.Helper, biz *biz.JobBiz) *JobService {
	return &JobService{helper: helper, biz: biz}
}

func (j *JobService) JobFriendRequest(ctx context.Context, topic string, event broker.Event, headers broker.Headers, msg *mq_kafka.FriendRequestMessage) error {
	j.helper.Infof("JobFriendRequest: %v", msg.RequestId)
	j.biz.HandleFriendRequest(ctx, event, msg)
	return nil
}
func (j *JobService) AfterConnectInit(ctx context.Context, topic string, event broker.Event, headers broker.Headers, msg *mq_kafka.ConnectInitMessage) error {
	j.helper.Infof("AfterConnectInit: %v", msg.UserId)
	j.biz.HandleConnectInit(ctx, event, msg)
	return nil
}
func (j *JobService) JobFriend(ctx context.Context, topic string, event broker.Event, headers broker.Headers, msg *mq_kafka.FriendMessage) error {
	j.helper.Infof("JobFriend: %v", msg.UserId)
	j.biz.HandleFriend(ctx, event, msg)
	return nil
}
func (j *JobService) JobMessage(ctx context.Context, topic string, event broker.Event, headers broker.Headers, msg *universal.Message) error {
	j.helper.Infof("JobMessage: %v", msg.MessageId)
	j.biz.HandleMessage(ctx, event, msg)
	return nil
}
