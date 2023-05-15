package service_logic

import (
	"api-gateway/api/v1/logic/logic"
	"api-gateway/internal/components/auth"
	"api-gateway/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type LogicService struct {
	client logic.LogicClient
	helper *log.Helper
	Jwt    *auth.JwtManager
}

func NewLogicService(client logic.LogicClient, helper *log.Helper, jwt *auth.JwtManager) *LogicService {
	return &LogicService{client: client, helper: helper, Jwt: jwt}
}
func (l *LogicService) GetConnectorUrl(ctx *gin.Context) {
	reply, err := l.client.GetConnectorUrl(ctx, &logic.GetConnectorUrlRequest{})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, reply)
}
