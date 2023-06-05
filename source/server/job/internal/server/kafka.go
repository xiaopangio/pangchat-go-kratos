package server

import (
	"context"
	"github.com/tx7do/kratos-transport/transport/kafka"
	"job/internal/conf"
	"job/internal/service"
	"job/internal/util/mq_kafka"
)

func NewKafkaConsumerServer(cf *conf.Bootstrap, service *service.JobService) (*kafka.Server, func()) {
	ctx := context.Background()
	kafkaSrv := kafka.NewServer(
		kafka.WithAddress(cf.Kafka.Addrs),
		kafka.WithCodec("json"),
	)
	var err error
	if err = kafkaSrv.RegisterSubscriber(
		ctx,
		cf.MessageQueue.FriendRequestTopic,
		cf.MessageQueue.FriendRequestGroup,
		true,
		mq_kafka.RegisterFriendRequestHandler(service.JobFriendRequest),
		mq_kafka.FriendRequestDataCreator,
	); err != nil {
		return nil, nil
	}
	if err = kafkaSrv.RegisterSubscriber(
		ctx,
		cf.MessageQueue.ConnectTopic,
		cf.MessageQueue.ConnectGroup,
		true,
		mq_kafka.RegisterConnectHandler(service.AfterConnectInit),
		mq_kafka.ConnectDataCreator,
	); err != nil {
		return nil, nil
	}
	if err = kafkaSrv.RegisterSubscriber(
		ctx,
		cf.MessageQueue.FriendTopic,
		cf.MessageQueue.FriendGroup,
		true,
		mq_kafka.RegisterFriendHandler(service.JobFriend),
		mq_kafka.FriendDataCreator,
	); err != nil {
		return nil, nil
	}
	if err = kafkaSrv.RegisterSubscriber(
		ctx,
		cf.MessageQueue.MessageTopic,
		cf.MessageQueue.MessageGroup,
		true,
		mq_kafka.RegisterMessageHandler(service.JobMessage),
		mq_kafka.MessageDataCreator,
	); err != nil {
		return nil, nil
	}

	err = kafkaSrv.Start(ctx)
	if err != nil {
		panic(err)
	}
	return kafkaSrv, func() {
		if err := kafkaSrv.Stop(ctx); err != nil {
			panic(err)
		}
	}

}
