package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "relationship/api/v1/relationship"
	"relationship/api/v1/universal"
	"relationship/internal/biz"
	"relationship/pkg"
)

type RelationShipService struct {
	pb.UnimplementedRelationShipServer
	biz    *biz.RelationshipBiz
	helper *log.Helper
}

func NewRelationShipService(biz *biz.RelationshipBiz, helper *log.Helper) *RelationShipService {
	return &RelationShipService{biz: biz, helper: helper}
}

func (s *RelationShipService) SendFriendRequest(ctx context.Context, req *pb.SendFriendRequestRequest) (*pb.SendFriendRequestResponse, error) {
	res, err := s.biz.SendFriendRequest(ctx, req.RequesterId, req.ReceiverId, req.NoteName, req.GroupName, req.Desc)
	if err != nil {
		s.helper.Errorf("send friend request error: %v", err)
		return nil, err
	}
	return &pb.SendFriendRequestResponse{
		FriendRequest: res,
	}, nil
}
func (s *RelationShipService) GetFriendRequestList(ctx context.Context, req *pb.GetFriendRequestListRequest) (*pb.GetFriendRequestListResponse, error) {
	if req.PageNumber <= 0 {
		req.PageNumber = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	res, total, err := s.biz.GetFriendRequestList(ctx, req.UserId, int(req.PageNumber), int(req.PageSize))
	if err != nil {
		s.helper.Errorf("get friend request list error: %v", err)
		return nil, err
	}
	var list []*universal.FriendRequest
	for _, v := range res {
		list = append(list, &universal.FriendRequest{
			RequestId:   v.RequestID,
			RequesterId: v.RequesterID,
			ReceiverId:  v.ReceiverID,
			Desc:        v.Desc,
			Status:      v.Status,
			CreateTime:  pkg.FormatTime(v.CreateAt),
			UpdateTime:  pkg.FormatTime(v.UpdateAt),
		})
	}
	return &pb.GetFriendRequestListResponse{
		FriendRequests: list,
		Total:          int64(total),
	}, nil
}
func (s *RelationShipService) GetFriendRequest(ctx context.Context, req *pb.GetFriendRequestRequest) (*pb.GetFriendRequestResponse, error) {
	friendRequest, err := s.biz.GetFriendRequest(ctx, req.RequestId)
	if err != nil {
		s.helper.Errorf("get friend request error: %v", err)
		return nil, err
	}
	s.helper.Infof("friend request: %v", friendRequest)
	return &pb.GetFriendRequestResponse{
		FriendRequest: &universal.FriendRequest{
			RequestId:   friendRequest.RequestID,
			RequesterId: friendRequest.RequesterID,
			ReceiverId:  friendRequest.ReceiverID,
			Desc:        friendRequest.Desc,
			Status:      friendRequest.Status,
			CreateTime:  pkg.FormatTime(friendRequest.CreateAt),
			UpdateTime:  pkg.FormatTime(friendRequest.UpdateAt),
		},
	}, nil
}
func (s *RelationShipService) DealFriendRequest(ctx context.Context, req *pb.DealFriendRequestRequest) (*pb.DealFriendRequestResponse, error) {
	switch req.Status {
	case pkg.Pending, pkg.Agreed, pkg.Refused:
	default:
		s.helper.Errorf("invalid status: %d", req.Status)
		return nil, pkg.InvalidArgumentError("invalid status: %d", req.Status)
	}
	err := s.biz.DealFriendRequest(ctx, req.RequestId, req.Status)
	if err != nil {
		s.helper.Errorf("deal friend request error: %v", err)
		return nil, err
	}
	return &pb.DealFriendRequestResponse{}, nil
}
func (s *RelationShipService) GetFriendList(ctx context.Context, req *pb.GetFriendListRequest) (*pb.GetFriendListResponse, error) {
	list, err := s.biz.GetFriendList(ctx, req.UserId)
	if err != nil {
		s.helper.Errorf("get friend list error: %v", err)
		return nil, err
	}
	return &pb.GetFriendListResponse{
		FriendGroups: list,
	}, nil
}
func (s *RelationShipService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.DeleteFriendResponse, error) {
	err := s.biz.DeleteFriend(ctx, req.UserId, req.FriendId)
	if err != nil {
		s.helper.Errorf("delete friend error: %v", err)
		return nil, err
	}
	return &pb.DeleteFriendResponse{}, nil
}
func (s *RelationShipService) GetFriendInfo(ctx context.Context, req *pb.GetFriendInfoRequest) (*pb.GetFriendInfoResponse, error) {
	reply, err := s.biz.GetFriendInfo(ctx, req.FriendId)
	if err != nil {
		s.helper.Errorf("get friend info error: %v", err)
		return nil, err
	}
	return &pb.GetFriendInfoResponse{
		CityName:     reply.CityName,
		ProvinceName: reply.ProvinceName,
		Desc:         reply.Desc,
	}, nil
}
func (s *RelationShipService) UpdateFriendInfo(ctx context.Context, req *pb.UpdateFriendInfoRequest) (*pb.UpdateFriendInfoResponse, error) {
	if req.NoteName == "" && req.GroupName == "" {
		s.helper.Errorf("note name and group name can not be empty at the same time")
		return nil, pkg.InvalidArgumentError("note name and group name can not be empty at the same time")
	}
	err := s.biz.UpdateFriendInfo(ctx, req.UserId, req.FriendId, req.NoteName, req.GroupName)
	if err != nil {
		s.helper.Errorf("update friend info error: %v", err)
		return nil, err
	}
	return &pb.UpdateFriendInfoResponse{}, nil
}
func (s *RelationShipService) CreateFriendGroup(ctx context.Context, req *pb.CreateFriendGroupRequest) (*pb.CreateFriendGroupResponse, error) {
	err := s.biz.CreateFriendGroup(ctx, req.UserId, req.GroupName)
	if err != nil {
		s.helper.Errorf("create friend group error: %v", err)
		return nil, err
	}
	return &pb.CreateFriendGroupResponse{}, nil
}
func (s *RelationShipService) UpdateFriendGroup(ctx context.Context, req *pb.UpdateFriendGroupRequest) (*pb.UpdateFriendGroupResponse, error) {
	err := s.biz.UpdateFriendGroup(ctx, req.UserId, req.GroupName, req.NewGroupName)
	if err != nil {
		s.helper.Errorf("update friend group error: %v", err)
		return nil, err
	}
	return &pb.UpdateFriendGroupResponse{}, nil
}
func (s *RelationShipService) DeleteFriendGroup(ctx context.Context, req *pb.DeleteFriendGroupRequest) (*pb.DeleteFriendGroupResponse, error) {
	err := s.biz.DeleteFriendGroup(ctx, req.UserId, req.GroupName)
	if err != nil {
		s.helper.Errorf("delete friend group error: %v", err)
		return nil, err
	}
	return &pb.DeleteFriendGroupResponse{}, nil
}
func (s *RelationShipService) GetFriendGroupList(ctx context.Context, req *pb.GetFriendGroupListRequest) (*pb.GetFriendGroupListResponse, error) {
	s.helper.Info("get friend group list")
	list, err := s.biz.GetFriendGroupList(ctx, req.UserId)
	if err != nil {
		s.helper.Errorf("get friend group list error: %v", err)
		return nil, err
	}
	return &pb.GetFriendGroupListResponse{
		GroupNames: list,
	}, nil
}
func (s *RelationShipService) GetFriendRequests(ctx context.Context, req *pb.GetFriendRequestsRequest) (*pb.GetFriendRequestsResponse, error) {
	s.helper.Info("get friend requests")
	requests, err := s.biz.GetFriendRequests(ctx, req.RequestIds)
	if err != nil {
		s.helper.Errorf("get friend requests error: %v", err)
		return nil, err
	}
	return &pb.GetFriendRequestsResponse{
		FriendRequests: requests,
	}, nil
}
