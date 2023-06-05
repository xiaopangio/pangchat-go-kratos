package service_online

import (
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterOnlineService(g *gin.RouterGroup, online *Online) {
	onlineRouter := g.Group("/online")
	onlineRouter.Use(middleware.Auth(online.Jwt))
	{
		onlineRouter.DELETE("", online.UnRegisterDevice)
	}
}
