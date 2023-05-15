package service

import (
	"connector/api/v1/connector"
	"connector/internal/biz"
	"connector/pkg"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type ConnectorServiceService struct {
	connector.UnimplementedConnectorServiceServer
	bz     *biz.ConnectorServiceBiz
	helper *log.Helper
}

func NewConnectorServiceService(bz *biz.ConnectorServiceBiz, helper *log.Helper) *ConnectorServiceService {
	return &ConnectorServiceService{bz: bz, helper: helper}
}

func (s *ConnectorServiceService) Login(ctx context.Context, req *connector.LoginRequest) (*connector.LoginResponse, error) {
	s.helper.Infof("Login request: %v", req)
	user, address, err := s.bz.Login(ctx, req.Type, req.Username, req.Password)
	if err != nil {
		s.helper.Errorf("Login error: %v", err)
		return nil, err
	}
	s.helper.Infof("Login success: %v", user)
	return &connector.LoginResponse{
		Uid:          user.UID,
		AccountId:    user.AccountID,
		NickName:     user.NickName,
		PersonalDesc: user.PersonalDesc,
		AvatarUrl:    user.AvatarURL,
		Address:      address,
	}, nil
}
func (s *ConnectorServiceService) Logout(ctx context.Context, req *connector.LogoutRequest) (*connector.LogoutResponse, error) {
	err := s.bz.Logout(ctx, req.Uid)
	if err != nil {
		s.helper.Errorf("Logout error: %v", err)
		return nil, err
	}
	s.helper.Infof("Logout success: %v", req.Uid)
	return &connector.LogoutResponse{}, nil
}
func (s *ConnectorServiceService) Connect(ctx *gin.Context) {
	uid := ctx.Query("uid")
	err := s.bz.Connect(ctx, uid)
	if err != nil {
		s.helper.Errorf("Connect error: %v", err)
		return
	}
	return
}
func (s *ConnectorServiceService) PushFriendRequests(ctx context.Context, req *connector.PushFriendRequestsRequest) (*connector.PushFriendRequestsResponse, error) {
	s.helper.Infof("PushFriendRequests request: %v", req)
	err := s.bz.PushFriendRequests(ctx, pkg.FormatInt(req.Uid), req.Requests)
	if err != nil {
		s.helper.Errorf("PushFriendRequests error: %v", err)
		return nil, err
	}
	s.helper.Infof("PushFriendRequests success: %v", req)
	return &connector.PushFriendRequestsResponse{}, nil
}
