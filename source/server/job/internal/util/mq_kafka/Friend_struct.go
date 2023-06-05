package mq_kafka

type FriendMessage struct {
	UserId   int64 `json:"user_id"`
	FriendId int64 `json:"friend_id"`
}
