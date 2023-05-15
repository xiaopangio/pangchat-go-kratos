package service_user

import (
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserService(g *gin.RouterGroup, us *UserService) {
	{
		g.GET("/ping", us.Ping)
		g.POST("/register", us.Register)
		g.POST("/reset", us.ResetPassword)
		g.POST("/sms", us.SendSmsCode)
		g.POST("/verify", us.VerifyCode)
		g.POST("/avatar", us.UploadAvatar)
	}
	user := g.Group("/user")
	{
		user.Use(middleware.Auth(us.Jwt))
		user.GET("/profile", us.Profile)
		user.POST("/avatar", us.UploadAvatar)
		user.GET("/avatar", us.GetAvatar)
		user.POST("/modify", us.ModifyProfile)
		user.POST("/modifyId", us.ModifyAccountID)
		user.POST("/modifyPassword", us.ModifyPasswd)
		user.POST("/bindphone", us.BindPhone)
	}
}
