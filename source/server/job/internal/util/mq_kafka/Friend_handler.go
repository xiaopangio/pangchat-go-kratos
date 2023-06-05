package mq_kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tx7do/kratos-transport/broker"
)

func FriendDataCreator() broker.Any {
	return &FriendMessage{}
}

type FriendHandler func(ctx context.Context, topic string, event broker.Event, headers broker.Headers, msg *FriendMessage) error

func RegisterFriendHandler(fn FriendHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		var msg *FriendMessage = nil
		eventMsgBody := event.Message().Body
		switch t := eventMsgBody.(type) {
		case []byte:
			msg = &FriendMessage{}
			if err := json.Unmarshal(t, msg); err != nil {
				return err
			}
		case string:
			msg = &FriendMessage{}
			if err := json.Unmarshal([]byte(t), msg); err != nil {
				return err
			}
		case *FriendMessage:
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
