package biz

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"message/api/v1/message"
	"message/api/v1/universal"
	"message/internal/components/broker"
	"message/internal/components/redis"
	"message/internal/conf"
	"message/pkg"
	"strings"
)

type MessageBiz struct {
	helper        *log.Helper
	broker        *broker.KafkaBroker
	repo          MessageRepo
	redisCli      *redis.Redis
	snowSlakeNode *snowflake.Node
	mqConfig      *conf.MessageQueue
	mysql         *gorm.DB
}

func NewMessageBiz(helper *log.Helper, broker *broker.KafkaBroker, mysql *gorm.DB, repo MessageRepo, cf *conf.Bootstrap, redisCli *redis.Redis, snowSlakeNode *snowflake.Node) *MessageBiz {
	return &MessageBiz{helper: helper, broker: broker, repo: repo, mysql: mysql, redisCli: redisCli, snowSlakeNode: snowSlakeNode, mqConfig: cf.MessageQueue}
}

type MessageRepo interface {
	StoreMessage(ctx context.Context, message *universal.Message) error
	GetLatestMessage(ctx context.Context, uid, friendId int64) (*universal.Message, error)
	GetUnreadMessageCount(ctx context.Context, uid, friendId, messageId int64) (int, error)
	GetMessagesBefore(ctx context.Context, senderId, receiverId, messageId int64, limit int) ([]*universal.Message, error)
	GetMessages(ctx context.Context, uid, friendId int64) ([]*universal.Message, error)
}

func BuildMessage(senderId int64, messageId int64) string {
	return fmt.Sprintf("%d.%d", senderId, messageId)
}
func BuildMessageKey(receiverId, senderId int64) string {
	return fmt.Sprintf("%s.%d.%d", redis.SingleMessageAckPrefix, receiverId, senderId)
}
func BuildMessageKeyPrefix(uid int64) string {
	return fmt.Sprintf("%s.%d", redis.SingleMessageAckPrefix, uid)
}
func (b *MessageBiz) DealSingleMessage(ctx context.Context, msg *universal.Message) error {
	messageId := b.snowSlakeNode.Generate()
	msg.MessageId = messageId.Int64()
	key := BuildMessageKey(msg.ReceiverId, msg.SenderId)
	if _, err := b.redisCli.Get(key); err != nil {
		if err == redis.Nil {
			//之前的消息确认由于异常原因丢失，构建redis key
			s := BuildMessage(msg.SenderId, msg.MessageId-1)
			if err = b.redisCli.Set(key, s, 0); err != nil {
				b.helper.Errorf("redis set 失败: %v", err)
				return pkg.InternalError("redis set 失败: %v", err)
			}
		} else {
			b.helper.Errorf("redis get 失败: %v", err)
			return pkg.InternalError("redis get 失败: %v", err)
		}
	}
	if err := b.broker.Publish(b.mqConfig.MessageTopic, msg); err != nil {
		b.helper.Errorf("publish message error: %v", err)
		return pkg.InternalError("publish message error")
	}
	if err := b.repo.StoreMessage(ctx, msg); err != nil {
		b.helper.Errorf("store message error: %v", err)
		return pkg.InternalError("store message error")
	}
	return nil
}
func (b *MessageBiz) UpdateAckMessage(ctx context.Context, senderId, receiverId, messageId int64) error {
	s := BuildMessage(senderId, messageId) // 格式为 senderId.messageId
	key := BuildMessageKey(receiverId, senderId)
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	if err := b.redisCli.Set(key, s, 0); err != nil {
		b.helper.Errorf("redis set 失败: %v", err)
		return pkg.InternalError("redis set 失败: %v", err)
	}
	return nil
}

func (b *MessageBiz) GetLatestUnreadMessageList(ctx context.Context, uid int64) ([]*universal.UnreadMessageInfo, error) {
	prefix := BuildMessageKeyPrefix(uid)
	keys, err := b.redisCli.GetPrefix(prefix)
	if err != nil {
		b.helper.Errorf("get redis message error: %v", err)
		return nil, pkg.InternalError("get redis message error: %v", err)
	}
	var msgsInRedis []string
	for _, key := range keys {
		msg, err := b.redisCli.Get(key)
		if err != nil {
			b.helper.Errorf("get redis message error: %v", err)
			continue
		}
		msgsInRedis = append(msgsInRedis, msg)
	}
	var unreadMessageInfos []*universal.UnreadMessageInfo
	for _, msgInRedis := range msgsInRedis {
		msg := strings.Split(msgInRedis, ".")
		if len(msg) != 2 {
			continue
		}
		senderId := pkg.ParseInt64(msg[0])
		messageId := pkg.ParseInt64(msg[1])
		unreadMessageCount, err := b.repo.GetUnreadMessageCount(ctx, uid, senderId, messageId)
		if err != nil {
			continue
		}
		if unreadMessageCount == 0 {
			continue
		}
		latestMessage, err := b.repo.GetLatestMessage(ctx, uid, senderId)
		if err != nil {
			continue
		}
		unreadMessageInfos = append(unreadMessageInfos, &universal.UnreadMessageInfo{
			LatestMessage: latestMessage,
			UnreadCount:   int64(unreadMessageCount),
		})
	}
	return unreadMessageInfos, nil
}

func (b *MessageBiz) GetUnloadMessages(ctx context.Context, senderId, receiverId, messageId, num int64) ([]*universal.Message, error) {
	messages, err := b.repo.GetMessagesBefore(ctx, senderId, receiverId, messageId, int(num))
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (b *MessageBiz) GetAllMessages(ctx context.Context, senderId, receiverId int64) ([]*universal.Message, error) {
	messages, err := b.repo.GetMessages(ctx, senderId, receiverId)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (b *MessageBiz) InitUnreadMessage(ctx context.Context, uid, friendId int64) error {
	m1 := BuildMessage(uid, 0)
	m2 := BuildMessage(friendId, 0)
	key1 := BuildMessageKey(friendId, uid)
	key2 := BuildMessageKey(uid, friendId)
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	if err := b.redisCli.Set(key1, m1, 0); err != nil {
		b.helper.Errorf("set redis message error: %v", err)
		return pkg.InternalError("set redis message error: %v", err)
	}
	if err := b.redisCli.Set(key2, m2, 0); err != nil {
		b.helper.Errorf("set redis message error: %v", err)
		return pkg.InternalError("set redis message error: %v", err)
	}
	return nil
}

func (b *MessageBiz) UpdateAckMessages(ctx context.Context, receiverId int64, ackMessageInfos []*message.AckMessageInfo) error {
	for _, info := range ackMessageInfos {
		if err := b.UpdateAckMessage(ctx, info.SenderId, receiverId, info.MessageId); err != nil {
			return err
		}
	}
	return nil
}
