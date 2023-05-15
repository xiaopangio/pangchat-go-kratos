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
	err := kafkaSrv.RegisterSubscriber(
		ctx,
		cf.MessageQueue.FriendRequestTopic,
		cf.MessageQueue.FriendRequestGroup,
		true,
		mq_kafka.RegisterFriendRequestHandler(service.JobFriendRequest),
		mq_kafka.FriendRequestDataCreator,
	)
	if err != nil {
		panic(err)
	}
	err = kafkaSrv.RegisterSubscriber(
		ctx,
		cf.MessageQueue.ConnectTopic,
		cf.MessageQueue.ConnectGroup,
		true,
		mq_kafka.RegisterConnectHandler(service.AfterConnectInit),
		mq_kafka.ConnectDataCreator,
	)
	if err != nil {
		panic(err)
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
