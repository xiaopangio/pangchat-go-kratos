package mq_kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tx7do/kratos-transport/broker"
)

func ConnectDataCreator() broker.Any {
	return &ConnectInitMessage{}
}

type ConnectHandler func(ctx context.Context, topic string, event broker.Event, headers broker.Headers, msg *ConnectInitMessage) error

func RegisterConnectHandler(fn ConnectHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		var msg *ConnectInitMessage = nil

		eventMsgBody := event.Message().Body
		switch t := eventMsgBody.(type) {
		case []byte:
			msg = &ConnectInitMessage{}
			if err := json.Unmarshal(t, msg); err != nil {
				return err
			}
		case string:
			msg = &ConnectInitMessage{}
			if err := json.Unmarshal([]byte(t), msg); err != nil {
				return err
			}
		case *ConnectInitMessage:
			msg = t
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}

		// Notice 使用service中的handler函数做业务处理，把msg传进去～
		if err := fn(ctx, event.Topic(), event, event.Message().Headers, msg); err != nil {
			return err
		}
		return nil
	}
}
