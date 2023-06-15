package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "message/api/v1/message"
	"message/internal/biz"
)

type MessageServiceService struct {
	pb.UnimplementedMessageServiceServer
	biz    *biz.MessageBiz
	helper *log.Helper
}

func NewMessageServiceService(biz *biz.MessageBiz, helper *log.Helper) *MessageServiceService {
	return &MessageServiceService{biz: biz, helper: helper}
}

func (s *MessageServiceService) DealSingleMessage(ctx context.Context, req *pb.DealSingleMessageRequest) (*pb.DealSingleMessageResponse, error) {
	if err := s.biz.DealSingleMessage(ctx, req.Message); err != nil {
		return nil, err
	}
	return &pb.DealSingleMessageResponse{}, nil
}
func (s *MessageServiceService) DealGroupMessage(ctx context.Context, req *pb.DealGroupMessageRequest) (*pb.DealGroupMessageResponse, error) {
	s.helper.Infof("DealGroupMessage")
	return &pb.DealGroupMessageResponse{}, nil
}
func (s *MessageServiceService) UpdateAckMessage(ctx context.Context, req *pb.UpdateAckMessageRequest) (*pb.UpdateAckMessageResponse, error) {
	s.helper.Infof("UpdateAckMessage")
	if err := s.biz.UpdateAckMessage(ctx, req.SenderId, req.ReceiverId, req.MessageId); err != nil {
		return nil, err
	}
	return &pb.UpdateAckMessageResponse{}, nil
}
func (s *MessageServiceService) GetLatestUnreadMessageList(ctx context.Context, req *pb.GetLatestUnreadMessageListRequest) (*pb.GetLatestUnreadMessageListResponse, error) {
	s.helper.Infof("GetLatestUnreadMessageList")
	list, err := s.biz.GetLatestUnreadMessageList(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &pb.GetLatestUnreadMessageListResponse{
		List: list,
	}, nil
}
func (s *MessageServiceService) GetUnloadMessages(ctx context.Context, req *pb.GetUnloadMessagesRequest) (*pb.GetUnloadMessagesResponse, error) {
	s.helper.Infof("GetUnloadMessages")
	messages, err := s.biz.GetUnloadMessages(ctx, req.SenderId, req.ReceiverId, req.MessageId, req.Num)
	if err != nil {
		return nil, err
	}
	return &pb.GetUnloadMessagesResponse{
		Messages: messages,
	}, nil
}
func (s *MessageServiceService) GetAllMessages(ctx context.Context, req *pb.GetAllMessageRequest) (*pb.GetAllMessageResponse, error) {
	s.helper.Infof("GetAllMessages")
	messages, err := s.biz.GetAllMessages(ctx, req.SenderId, req.ReceiverId)
	if err != nil {
		return nil, err
	}
	return &pb.GetAllMessageResponse{
		Messages: messages,
	}, nil
}
func (s *MessageServiceService) InitUnreadMessage(ctx context.Context, req *pb.InitUnreadMessageRequest) (*pb.InitUnreadMessageResponse, error) {
	s.helper.Infof("InitUnreadMessage")
	err := s.biz.InitUnreadMessage(ctx, req.Uid, req.FriendId)
	if err != nil {
		return nil, err
	}
	return &pb.InitUnreadMessageResponse{}, nil
}
func (s *MessageServiceService) UpdateAckMessages(ctx context.Context, req *pb.UpdateAckMessagesRequest) (*pb.UpdateAckMessagesResponse, error) {
	s.helper.Infof("UpdateAckMessages")
	err := s.biz.UpdateAckMessages(ctx, req.ReceiverId, req.List)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAckMessagesResponse{}, nil
}
