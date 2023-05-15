package mq_kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tx7do/kratos-transport/broker"
)

func FriendRequestDataCreator() broker.Any {
	return &FriendRequestMessage{}
}

type FriendRequestHandler func(ctx context.Context, topic string, event broker.Event, headers broker.Headers, msg *FriendRequestMessage) error

func RegisterFriendRequestHandler(fn FriendRequestHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		var msg *FriendRequestMessage = nil

		eventMsgBody := event.Message().Body
		switch t := eventMsgBody.(type) {
		case []byte:
			fmt.Println("byte............................")
			msg = &FriendRequestMessage{}
			if err := json.Unmarshal(t, msg); err != nil {
				return err
			}
		case string:
			fmt.Println("string............................")
			msg = &FriendRequestMessage{}
			if err := json.Unmarshal([]byte(t), msg); err != nil {
				return err
			}
		case *FriendRequestMessage:
			fmt.Println("相同类型............................", t.RequestId, t.UserId, t.PublishAt)
			msg = t
		default:
			fmt.Println("default............................")
			return fmt.Errorf("unsupported type: %T", t)
		}

		// Notice 使用service中的handler函数做业务处理，把msg传进去～
		fmt.Println("event.Message().Headers:>>>>>>>>>>> ", event.Message().Headers)
		if err := fn(ctx, event.Topic(), event, event.Message().Headers, msg); err != nil {
			return err
		}

		fmt.Println("msg:>>>>>>>>>>>> ", msg)

		return nil
	}
}
