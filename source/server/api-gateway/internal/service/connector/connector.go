package service_connector

import (
	"api-gateway/api/v1/connector"
	"api-gateway/internal/components/auth"
	"api-gateway/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type ConnectorService struct {
	client connector.ConnectorServiceClient
	helper *log.Helper
	Jwt    *auth.JwtManager
}

func NewConnectorService(client connector.ConnectorServiceClient, helper *log.Helper, jwt *auth.JwtManager) *ConnectorService {
	return &ConnectorService{client: client, helper: helper, Jwt: jwt}
}
func (c *ConnectorService) Login(ctx *gin.Context) {
	var req LoginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		c.helper.Errorf("bind err: %v", err)
		pkg.Validator(ctx, err)
		return
	}
	loginReq := connector.LoginRequest{
		Type:     int64(req.T),
		Username: req.Username,
		Password: req.Password,
	}
	context, cancel := pkg.NewContext(ctx)
	defer cancel()
	resp, err := c.client.Login(context, &loginReq)
	if err != nil {
		c.helper.Errorf("Login error: %v", err)
		pkg.Fail(ctx, err)
		return
	}
	token := &auth.UserToken{
		Uid:          pkg.FormatInt(resp.Uid),
		NickName:     resp.NickName,
		AccountId:    resp.AccountId,
		PersonalDesc: resp.PersonalDesc,
		Avatar:       resp.AvatarUrl,
		City:         resp.Address.City,
		Province:     resp.Address.Province,
	}
	userToken, err := c.Jwt.SignUser(token)
	if err != nil {
		c.helper.Errorf("SignUser error: %v", err)
		pkg.Fail(ctx, err)
		return
	}
	loginRes := LoginResponse{
		Token: userToken,
	}
	pkg.Ok(ctx, loginRes)
}
func (c *ConnectorService) Logout(ctx *gin.Context) {
	token, exists := ctx.Get("token")
	if !exists {
		pkg.Forbidden(ctx)
		c.helper.Errorf("user token not found")
		return
	}
	userToken, ok := token.(*auth.UserToken)
	if !ok {
		pkg.Forbidden(ctx)
		c.helper.Errorf("user token not found")
		return
	}
	logoutReq := connector.LogoutRequest{
		Uid: userToken.Uid,
	}
	context, cancel := pkg.NewContext(ctx)
	defer cancel()
	_, err := c.client.Logout(context, &logoutReq)
	if err != nil {
		c.helper.Errorf("Logout error: %v", err)
		return
	}
	pkg.Ok(ctx, nil)
}
