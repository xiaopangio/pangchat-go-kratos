package service_message

import (
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterMessageService(g *gin.RouterGroup, message *Message) {
	messageRouter := g.Group("/message")
	messageRouter.Use(middleware.Auth(message.Jwt))
	{
		messageRouter.GET("/unload", message.GetUnloadMessages)
		messageRouter.GET("/all", message.GetAllMessages)
	}
}
