package service

import "github.com/gin-gonic/gin"

func RegisterConnectorService(g *gin.Engine, s *ConnectorServiceService) {
	g.GET("/connect", s.Connect)
}
