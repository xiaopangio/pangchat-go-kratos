package online

import (
	"api-gateway/api/v1/online"
	"api-gateway/internal/components/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type Online struct {
	helper *log.Helper
	Jwt    *auth.JwtManager
	client *online.OnlineClient
}

func NewOnline(helper *log.Helper, jwt *auth.JwtManager, client *online.OnlineClient) *Online {
	return &Online{helper: helper, Jwt: jwt, client: client}
}
func (o *Online) RegisterDevice(ctx *gin.Context) {

}
