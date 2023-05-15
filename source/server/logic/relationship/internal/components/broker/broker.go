package broker

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/kafka"
	"relationship/internal/conf"
)

type KafkaBroker struct {
	broker.Broker
	helper *log.Helper
	cf     *conf.Kafka
}

func NewKafkaBroker(helper *log.Helper, cf *conf.Bootstrap) *KafkaBroker {
	ctx := context.Background()
	helper.Info("init kafka broker")
	helper.Infof("init kafka broker addrs:%v", cf.Kafka.Addrs)
	kafkaBroker := kafka.NewBroker(
		broker.WithOptionContext(ctx),
		broker.WithAddress(cf.Kafka.Addrs...),
		broker.WithCodec("json"),
		kafka.WithBatchSize(1),
		kafka.WithAsync(true),
	)
	err := kafkaBroker.Init()
	if err != nil {
		helper.Fatalf("init kafka broker err:%v", err)
	}
	return &KafkaBroker{helper: helper, cf: cf.Kafka, Broker: kafkaBroker}
}
