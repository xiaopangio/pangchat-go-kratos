package service_online

import (
	"api-gateway/api/v1/online"
	"api-gateway/internal/components/auth"
	"api-gateway/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type Online struct {
	helper *log.Helper
	Jwt    *auth.JwtManager
	client online.OnlineClient
}

func NewOnline(helper *log.Helper, jwt *auth.JwtManager, client online.OnlineClient) *Online {
	return &Online{helper: helper, Jwt: jwt, client: client}
}
func (o *Online) UnRegisterDevice(ctx *gin.Context) {
	var req UnregisterDeviceRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		o.helper.Info(req.Uid)
		pkg.Validator(ctx, err)
		return
	}
	o.helper.Info(req.Uid)
	_, err := o.client.UnregisterDevice(ctx, &online.UnregisterDeviceRequest{
		Uid: pkg.ParseInt64(req.Uid),
	})
	err = pkg.HandlerError(ctx, err)
	if err != nil {
		return
	}
	pkg.Ok(ctx, nil)
}
