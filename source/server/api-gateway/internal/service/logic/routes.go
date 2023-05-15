package service_logic

import (
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterLogicService(g *gin.RouterGroup, cs *LogicService) {
	g.GET("/connectorUrl", middleware.Auth(cs.Jwt), cs.GetConnectorUrl)
}
