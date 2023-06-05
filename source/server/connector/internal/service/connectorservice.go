package service

import (
	"connector/api/v1/connector"
	"connector/internal/biz"
	"connector/pkg"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

// ConnectorServiceService 连接器服务
type ConnectorServiceService struct {
	connector.UnimplementedConnectorServiceServer
	bz     *biz.ConnectorServiceBiz
	helper *log.Helper
}

// NewConnectorServiceService 新建连接器服务实例
func NewConnectorServiceService(bz *biz.ConnectorServiceBiz, helper *log.Helper) *ConnectorServiceService {
	return &ConnectorServiceService{bz: bz, helper: helper}
}

// Login 登录
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

// Logout 登出
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

// PushFriendRequests 推送好友请求到前端
func (s *ConnectorServiceService) PushFriendRequests(ctx context.Context, req *connector.PushFriendRequestsRequest) (*connector.PushFriendRequestsResponse, error) {
	s.helper.Infof("PushFriendRequests request: %v", req)
	err := s.bz.PushFriendRequests(ctx, pkg.FormatInt(req.Uid), req.Requests)
	if err != nil {
		s.helper.Errorf("PushFriendRequests error: %v", err)
		return nil, err
	}
	return &connector.PushFriendRequestsResponse{}, nil
}

// PushFriend 推送好友到前端
func (s *ConnectorServiceService) PushFriend(ctx context.Context, req *connector.PushFriendRequest) (*connector.PushFriendResponse, error) {
	s.helper.Infof("PushFriend request: %v", req)
	err := s.bz.PushFriend(ctx, pkg.FormatInt(req.Uid), req.Friends)
	if err != nil {
		s.helper.Errorf("PushFriend error: %v", err)
		return nil, err
	}
	s.helper.Infof("PushFriend success: %v", req)
	return &connector.PushFriendResponse{}, nil
}

// PushMessage 推送消息到前端
func (s *ConnectorServiceService) PushMessage(ctx context.Context, req *connector.PushMessageRequest) (*connector.PushMessageResponse, error) {
	s.helper.Infof("PushMessage request: %v", req)
	err := s.bz.PushMessage(ctx, pkg.FormatInt(req.Uid), req.Message)
	if err != nil {
		s.helper.Errorf("PushMessage error: %v", err)
		return nil, err
	}
	s.helper.Infof("PushMessage success: %v", req)
	return &connector.PushMessageResponse{}, nil
}

// ReplyMessage 用户发送消息后的回执，比如生成的消息ID
func (s *ConnectorServiceService) ReplyMessage(ctx context.Context, req *connector.ReplyMessageRequest) (*connector.ReplyMessageResponse, error) {
	s.helper.Infof("ReplyMessage request: %v", req)
	err := s.bz.ReplyMessage(ctx, pkg.FormatInt(req.Uid), req.Message)
	if err != nil {
		s.helper.Errorf("ReplyMessage error: %v", err)
		return nil, err
	}
	s.helper.Infof("ReplyMessage success: %v", req)
	return &connector.ReplyMessageResponse{}, nil
}

// PushUnreadMessageList 推送未读消息列表到前端
func (s *ConnectorServiceService) PushUnreadMessageList(ctx context.Context, req *connector.PushUnreadMessageListRequest) (*connector.PushUnreadMessageListResponse, error) {
	s.helper.Infof("PushUnreadMessageList request: %v", req)
	err := s.bz.PushUnreadMessageList(ctx, pkg.FormatInt(req.Uid), req.List)
	if err != nil {
		s.helper.Errorf("PushUnreadMessageList error: %v", err)
		return nil, err
	}
	s.helper.Infof("PushUnreadMessageList success: %v", req)
	return &connector.PushUnreadMessageListResponse{}, nil
}
