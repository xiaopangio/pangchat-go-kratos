package model

type MessageInRedis struct {
	ReceiverId string `json:"receiver_id"`
	MessageId  string `json:"message_id"`
}
