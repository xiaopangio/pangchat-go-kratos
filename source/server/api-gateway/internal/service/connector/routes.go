package service_connector

import (
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterConnectorService(g *gin.RouterGroup, cs *ConnectorService) {
	g.POST("/login", cs.Login)
	g.POST("/logout", middleware.Auth(cs.Jwt), cs.Logout)
}
