package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"message/api/v1/universal"
	"message/internal/biz"
	"message/internal/data/orm/dal"
	"message/internal/data/orm/model"
	"message/pkg"
)

type MessageRepoImpl struct {
	helper *log.Helper
}

func (m *MessageRepoImpl) GetMessages(ctx context.Context, uid, friendId string) ([]*universal.Message, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	messages, err := dal.Message.WithContext(ctx).
		Where(dal.Message.SenderID.Eq(friendId), dal.Message.ReceiverID.Eq(uid)).
		Or(dal.Message.SenderID.Eq(uid), dal.Message.ReceiverID.Eq(friendId)).
		Order(dal.Message.MessageID.Desc()).
		Find()
	if err != nil {
		return nil, pkg.InternalError("get messages error")
	}
	var result []*universal.Message
	for _, message := range messages {
		result = append(result, &universal.Message{
			MessageId:  message.MessageID,
			Type:       int64(message.Type),
			Content:    message.Content,
			SenderId:   message.SenderID,
			ReceiverId: message.ReceiverID,
			SendAt:     message.SendAt,
		})
	}
	return result, nil
}

func (m *MessageRepoImpl) GetMessagesBefore(ctx context.Context, senderId, receiverId string, messageId string, limit int) ([]*universal.Message, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	messages, err := dal.Message.WithContext(ctx).
		Where(dal.Message.SenderID.Eq(senderId)).
		Where(dal.Message.ReceiverID.Eq(receiverId)).
		Where(dal.Message.MessageID.Lt(messageId)).
		Order(dal.Message.MessageID.Desc()).
		Limit(limit).
		Find()
	if err != nil {
		return nil, pkg.InternalError("get messages before error")
	}
	var result []*universal.Message
	for _, message := range messages {
		result = append(result, &universal.Message{
			MessageId:  message.MessageID,
			Type:       int64(message.Type),
			Content:    message.Content,
			SenderId:   message.SenderID,
			ReceiverId: message.ReceiverID,
			SendAt:     message.SendAt,
		})
	}
	return result, nil
}

func (m *MessageRepoImpl) GetLatestMessage(ctx context.Context, uid, friendId string) (*universal.Message, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	message, err := dal.Message.WithContext(ctx).
		Where(dal.Message.ReceiverID.
			Eq(uid)).
		Where(dal.Message.SenderID.Eq(friendId)).
		Order(dal.Message.MessageID.Desc()).
		First()
	if err != nil {
		return nil, pkg.InternalError("get latest message error")
	}
	return &universal.Message{
		MessageId:  message.MessageID,
		Type:       int64(message.Type),
		Content:    message.Content,
		SenderId:   message.SenderID,
		ReceiverId: message.ReceiverID,
		SendAt:     message.SendAt,
	}, nil
}

func (m *MessageRepoImpl) GetUnreadMessageCount(ctx context.Context, uid, friendId string, messageId string) (int, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return 0, err
	}
	count, err := dal.Message.WithContext(ctx).
		Where(dal.Message.SenderID.Eq(friendId)).
		Where(dal.Message.ReceiverID.Eq(uid)).
		Where(dal.Message.MessageID.Gt(messageId)).
		Count()
	if err != nil {
		return 0, pkg.InternalError("get unread message count error")
	}
	return int(count), nil
}

func NewMessageRepoImpl(helper *log.Helper) biz.MessageRepo {
	return &MessageRepoImpl{helper: helper}
}

func (m *MessageRepoImpl) StoreMessage(ctx context.Context, message *universal.Message) error {
	m.helper.Infof("StoreMessage")
	ms := &model.Message{
		MessageID:  message.MessageId,
		Type:       int32(message.Type),
		Content:    message.Content,
		SenderID:   message.SenderId,
		ReceiverID: message.ReceiverId,
		SendAt:     message.SendAt,
	}
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	if err := dal.Message.WithContext(ctx).Create(ms); err != nil {
		return pkg.InternalError("store message error")
	}
	return nil
}
