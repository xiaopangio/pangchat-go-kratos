package service_message

import "api-gateway/api/v1/universal"

type GetUnloadMessagesRequest struct {
	SenderId   string `form:"sender_id" json:"sender_id" binding:"required"`
	ReceiverId string `form:"receiver_id" json:"receiver_id" binding:"required"`
	MessageId  string `form:"message_id" json:"message_id" binding:"required"`
	Num        int    `form:"num" json:"num" binding:"required"`
}
type GetUnloadMessagesResponse struct {
	Messages []*universal.Message `json:"messages"`
}
type GetAllMessagesRequest struct {
	SenderId   string `form:"sender_id" json:"sender_id" binding:"required"`
	ReceiverId string `form:"receiver_id" json:"receiver_id" binding:"required"`
}
type GetAllMessagesResponse struct {
	Messages []*universal.Message `json:"messages"`
}
