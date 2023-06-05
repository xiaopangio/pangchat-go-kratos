package model

import "time"

type FriendRequestMessage struct {
	RequestId int64     `json:"request_id"`
	UserId    int64     `json:"user_id"`
	PublishAt time.Time `json:"publish_at"`
}
type FriendMessage struct {
	UserId   int64 `json:"user_id"`
	FriendId int64 `json:"friend_id"`
}
