package service_relationship

import (
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRelationshipService(g *gin.RouterGroup, rs *Relationship) {
	relationship := g.Group("/relationship")
	relationship.Use(middleware.Auth(rs.Jwt))
	{
		relationship.POST("/friend/request", rs.SendFriendRequest)
		relationship.GET("/friend/request", rs.GetFriendRequest)
		relationship.GET("/friend/request/list", rs.GetFriendRequestList)
		relationship.PUT("/friend/request", rs.DealFriendRequest)
		relationship.GET("/friend/list", rs.GetFriendList)
		relationship.DELETE("/friend", rs.DeleteFriend)
		relationship.GET("/friend", rs.GetFriendInfo)
		relationship.PUT("/friend", rs.UpdateFriendInfo)
		relationship.POST("/friend/group", rs.CreateFriendGroup)
		relationship.PUT("/friend/group", rs.UpdateFriendGroup)
		relationship.DELETE("/friend/group", rs.DeleteFriendGroup)
		relationship.GET("/friend/group/list", rs.GetFriendGroupList)
	}
}
