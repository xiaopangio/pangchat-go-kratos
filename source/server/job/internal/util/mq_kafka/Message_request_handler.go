package mq_kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tx7do/kratos-transport/broker"
	"job/api/v1/universal"
)

func MessageDataCreator() broker.Any {
	return &universal.Message{}
}

type MessageHandler func(ctx context.Context, topic string, event broker.Event, headers broker.Headers, msg *universal.Message) error

func RegisterMessageHandler(fn MessageHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		var msg *universal.Message = nil
		eventMsgBody := event.Message().Body
		switch t := eventMsgBody.(type) {
		case []byte:
			msg = &universal.Message{}
			if err := json.Unmarshal(t, msg); err != nil {
				return err
			}
		case string:
			msg = &universal.Message{}
			if err := json.Unmarshal([]byte(t), msg); err != nil {
				return err
			}
		case *universal.Message:
			msg = t
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		if err := fn(ctx, event.Topic(), event, event.Message().Headers, msg); err != nil {
			return err
		}
		return nil
	}
}
