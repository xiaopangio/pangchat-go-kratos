package service_message

import (
	"api-gateway/api/v1/message"
	"api-gateway/internal/components/auth"
	"api-gateway/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type Message struct {
	helper *log.Helper
	Jwt    *auth.JwtManager
	client message.MessageServiceClient
}

func NewMessage(helper *log.Helper, jwt *auth.JwtManager, client message.MessageServiceClient) *Message {
	return &Message{helper: helper, Jwt: jwt, client: client}
}

func (m *Message) GetUnloadMessages(ctx *gin.Context) {
	var req GetUnloadMessagesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	reply, err := m.client.GetUnloadMessages(ctx, &message.GetUnloadMessagesRequest{
		SenderId:   req.SenderId,
		ReceiverId: req.ReceiverId,
		MessageId:  req.MessageId,
		Num:        int64(req.Num),
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, &GetUnloadMessagesResponse{
		Messages: reply.Messages,
	})

}
func (m *Message) GetAllMessages(ctx *gin.Context) {
	var req GetAllMessagesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	reply, err := m.client.GetAllMessages(ctx, &message.GetAllMessageRequest{
		SenderId:   req.SenderId,
		ReceiverId: req.ReceiverId,
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, &GetAllMessagesResponse{
		Messages: reply.Messages,
	})
}
