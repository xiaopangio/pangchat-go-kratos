package service_logic

import (
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterLogicService(g *gin.RouterGroup, cs *LogicService) {
	g.Use(middleware.Auth(cs.Jwt))
	g.GET("/connectorUrl", cs.GetConnectorUrl)
	g.GET("/toolOptions", cs.GetToolOptions)
	g.GET("/preEmojis", cs.GetPreEmojis)
	g.POST("/uploadFile", cs.UploadFile)
	g.GET("/downloadFile", cs.DownloadFile)
}
