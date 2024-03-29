package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "relationship/api/v1/relationship"
	"relationship/api/v1/universal"
	"relationship/internal/biz"
	"relationship/internal/common/constant"
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

// SendFriendRequest 发送好友请求
func (s *RelationShipService) SendFriendRequest(ctx context.Context, req *pb.SendFriendRequestRequest) (*pb.SendFriendRequestResponse, error) {
	res, err := s.biz.SendFriendRequest(ctx, req.RequesterId, req.ReceiverId, req.NoteName, req.GroupName, req.Desc)
	if err != nil {
		return nil, err
	}
	return &pb.SendFriendRequestResponse{
		FriendRequest: res,
	}, nil
}

// GetFriendRequestList 获取好友请求列表，分页
func (s *RelationShipService) GetFriendRequestList(ctx context.Context, req *pb.GetFriendRequestListRequest) (*pb.GetFriendRequestListResponse, error) {
	if req.PageNumber <= 0 {
		req.PageNumber = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	res, total, err := s.biz.GetFriendRequestList(ctx, req.UserId, int(req.PageNumber), int(req.PageSize))
	if err != nil {
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

// GetFriendRequest 获取好友请求
func (s *RelationShipService) GetFriendRequest(ctx context.Context, req *pb.GetFriendRequestRequest) (*pb.GetFriendRequestResponse, error) {
	friendRequest, err := s.biz.GetFriendRequest(ctx, req.RequestId)
	if err != nil {
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

// DealFriendRequest 处理好友请求
func (s *RelationShipService) DealFriendRequest(ctx context.Context, req *pb.DealFriendRequestRequest) (*pb.DealFriendRequestResponse, error) {
	switch req.Status {
	case constant.Pending, constant.Agreed, constant.Refused:
	default:
		s.helper.Errorf("无效的状态: %d", req.Status)
		return nil, pkg.InvalidArgumentError("\"无效的状态: %d", req.Status)
	}
	err := s.biz.DealFriendRequest(ctx, req.RequestId, req.Status, req.NoteName, req.GroupName)
	if err != nil {
		return nil, err
	}
	return &pb.DealFriendRequestResponse{}, nil
}

// GetFriendList 获取好友列表
func (s *RelationShipService) GetFriendList(ctx context.Context, req *pb.GetFriendListRequest) (*pb.GetFriendListResponse, error) {
	list, err := s.biz.GetFriendList(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetFriendListResponse{
		Friends: list,
	}, nil
}

// DeleteFriend 删除好友
func (s *RelationShipService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.DeleteFriendResponse, error) {
	err := s.biz.DeleteFriend(ctx, req.UserId, req.FriendId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteFriendResponse{}, nil
}

// GetFriendInfo 获取好友信息
func (s *RelationShipService) GetFriendInfo(ctx context.Context, req *pb.GetFriendInfoRequest) (*pb.GetFriendInfoResponse, error) {
	reply, err := s.biz.GetFriendInfo(ctx, req.FriendId)
	if err != nil {
		return nil, err
	}
	return &pb.GetFriendInfoResponse{
		CityName:     reply.CityName,
		ProvinceName: reply.ProvinceName,
		Desc:         reply.Desc,
		AccountId:    reply.AccountId,
	}, nil
}

// UpdateFriendInfo 更新好友信息
func (s *RelationShipService) UpdateFriendInfo(ctx context.Context, req *pb.UpdateFriendInfoRequest) (*pb.UpdateFriendInfoResponse, error) {
	if req.NoteName == "" && req.GroupName == "" {
		s.helper.Errorf("noteName和groupName不能同时为空")
		return nil, pkg.InvalidArgumentError("noteName和groupName不能同时为空")
	}
	err := s.biz.UpdateFriendInfo(ctx, req.UserId, req.FriendId, req.NoteName, req.GroupName)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateFriendInfoResponse{}, nil
}

// CreateFriendGroup 创建好友分组
func (s *RelationShipService) CreateFriendGroup(ctx context.Context, req *pb.CreateFriendGroupRequest) (*pb.CreateFriendGroupResponse, error) {
	err := s.biz.CreateFriendGroup(ctx, req.UserId, req.GroupName)
	if err != nil {
		return nil, err
	}
	return &pb.CreateFriendGroupResponse{}, nil
}

// UpdateFriendGroup  更新好友分组信息
func (s *RelationShipService) UpdateFriendGroup(ctx context.Context, req *pb.UpdateFriendGroupRequest) (*pb.UpdateFriendGroupResponse, error) {
	err := s.biz.UpdateFriendGroup(ctx, req.UserId, req.GroupName, req.NewGroupName)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateFriendGroupResponse{}, nil
}

// DeleteFriendGroup 删除好友分组
func (s *RelationShipService) DeleteFriendGroup(ctx context.Context, req *pb.DeleteFriendGroupRequest) (*pb.DeleteFriendGroupResponse, error) {
	err := s.biz.DeleteFriendGroup(ctx, req.UserId, req.GroupName)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteFriendGroupResponse{}, nil
}

// GetFriendGroupList 获取好友分组列表
func (s *RelationShipService) GetFriendGroupList(ctx context.Context, req *pb.GetFriendGroupListRequest) (*pb.GetFriendGroupListResponse, error) {
	list, err := s.biz.GetFriendGroupList(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetFriendGroupListResponse{
		GroupNames: list,
	}, nil
}

// GetFriendRequests 获取好友请求
func (s *RelationShipService) GetFriendRequests(ctx context.Context, req *pb.GetFriendRequestsRequest) (*pb.GetFriendRequestsResponse, error) {
	requests, err := s.biz.GetFriendRequests(ctx, req.RequestIds)
	if err != nil {
		return nil, err
	}
	return &pb.GetFriendRequestsResponse{
		FriendRequests: requests,
	}, nil
}

// GetOneFriend 获取单个好友
func (s *RelationShipService) GetOneFriend(ctx context.Context, req *pb.GetOneFriendRequest) (*pb.GetONeFriendResponse, error) {
	friend, err := s.biz.GetOneFriend(ctx, req.UserId, req.FriendId)
	if err != nil {
		return nil, err
	}
	return &pb.GetONeFriendResponse{
		Friend: friend,
	}, nil
}

// CreateGroup 创建群组
func (s *RelationShipService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	group, err := s.biz.CreateGroup(ctx, req.GroupLeaderId, req.GroupName, req.GroupAvatar, req.GroupDesc)
	if err != nil {
		return nil, err
	}
	return &pb.CreateGroupResponse{
		Group: group,
	}, nil
}

// GetGroupList 获取群组列表
func (s *RelationShipService) GetGroupList(ctx context.Context, req *pb.GetGroupListRequest) (*pb.GetGroupListResponse, error) {
	list, err := s.biz.GetGroupList(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetGroupListResponse{
		Groups: list,
	}, nil
}

// GetGroupInfo 获取群组信息
func (s *RelationShipService) GetGroupInfo(ctx context.Context, req *pb.GetGroupInfoRequest) (*pb.GetGroupInfoResponse, error) {
	group, err := s.biz.GetGroupInfo(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}
	return &pb.GetGroupInfoResponse{
		Group: group,
	}, nil
}

// UpdateGroupInfo 更新群组信息
func (s *RelationShipService) UpdateGroupInfo(ctx context.Context, req *pb.UpdateGroupInfoRequest) (*pb.UpdateGroupInfoResponse, error) {
	err := s.biz.UpdateGroupInfo(ctx, req.GroupId, req.GroupName, req.GroupAvatar, req.GroupDesc)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateGroupInfoResponse{}, nil
}

// DeleteGroup 删除群组
func (s *RelationShipService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*pb.DeleteGroupResponse, error) {
	err := s.biz.DeleteGroup(ctx, req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteGroupResponse{}, nil
}

// GetGroupMemberList 获取群组成员列表
func (s *RelationShipService) GetGroupMemberList(ctx context.Context, req *pb.GetGroupMemberListRequest) (*pb.GetGroupMemberListResponse, error) {
	list, err := s.biz.GetGroupMemberList(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}
	return &pb.GetGroupMemberListResponse{
		GroupMembers: list,
	}, nil
}

// GetGroupMemberInfo 获取群组成员信息
func (s *RelationShipService) GetGroupMemberInfo(ctx context.Context, req *pb.GetGroupMemberInfoRequest) (*pb.GetGroupMemberInfoResponse, error) {
	member, err := s.biz.GetGroupMemberInfo(ctx, req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetGroupMemberInfoResponse{
		GroupMember: member,
	}, nil
}

// UpdateGroupMemberInfo 更新群组成员信息
func (s *RelationShipService) UpdateGroupMemberInfo(ctx context.Context, req *pb.UpdateGroupMemberInfoRequest) (*pb.UpdateGroupMemberInfoResponse, error) {
	err := s.biz.UpdateGroupMemberInfo(ctx, req.GroupId, req.UserId, req.GroupNoteName, req.MemberNoteName)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateGroupMemberInfoResponse{}, nil
}

// DeleteGroupMember 删除群组成员
func (s *RelationShipService) DeleteGroupMember(ctx context.Context, req *pb.DeleteGroupMemberRequest) (*pb.DeleteGroupMemberResponse, error) {
	err := s.biz.DeleteGroupMember(ctx, req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteGroupMemberResponse{}, nil
}

// SendGroupRequest 发送群组请求
func (s *RelationShipService) SendGroupRequest(ctx context.Context, req *pb.SendGroupRequestRequest) (*pb.SendGroupRequestResponse, error) {
	request, err := s.biz.SendGroupRequest(ctx, req.RequesterId, req.GroupId, req.Desc)
	if err != nil {
		return nil, err
	}
	return &pb.SendGroupRequestResponse{
		GroupRequest: request,
	}, nil
}

// GetGroupRequestList 获取群组请求列表
func (s *RelationShipService) GetGroupRequestList(ctx context.Context, req *pb.GetGroupRequestListRequest) (*pb.GetGroupRequestListResponse, error) {
	list, err := s.biz.GetGroupRequestList(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}
	return &pb.GetGroupRequestListResponse{
		GroupRequests: list,
	}, nil
}

// GetGroupRequest 获取群组请求
func (s *RelationShipService) GetGroupRequest(ctx context.Context, req *pb.GetGroupRequestRequest) (*pb.GetGroupRequestResponse, error) {
	request, err := s.biz.GetGroupRequest(ctx, req.RequestId)
	if err != nil {
		return nil, err
	}
	return &pb.GetGroupRequestResponse{
		GroupRequest: request,
	}, nil
}

// GetGroupRequests 获取群组请求
func (s *RelationShipService) GetGroupRequests(ctx context.Context, req *pb.GetGroupRequestsRequest) (*pb.GetGroupRequestsResponse, error) {
	requests, err := s.biz.GetGroupRequests(ctx, req.RequestIds)
	if err != nil {
		return nil, err
	}
	return &pb.GetGroupRequestsResponse{
		GroupRequests: requests,
	}, nil
}

// DealGroupRequest 处理群组请求
func (s *RelationShipService) DealGroupRequest(ctx context.Context, req *pb.DealGroupRequestRequest) (*pb.DealGroupRequestResponse, error) {
	err := s.biz.DealGroupRequest(ctx, req.RequestId, req.Status)
	if err != nil {
		return nil, err
	}
	return &pb.DealGroupRequestResponse{}, nil
}

// CreateGroupAdmin 创建群组管理员
func (s *RelationShipService) CreateGroupAdmin(ctx context.Context, req *pb.CreateGroupAdminRequest) (*pb.CreateGroupAdminResponse, error) {
	err := s.biz.CreateGroupAdmin(ctx, req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateGroupAdminResponse{}, nil
}

// DeleteGroupAdmin 删除群组管理员
func (s *RelationShipService) DeleteGroupAdmin(ctx context.Context, req *pb.DeleteGroupAdminRequest) (*pb.DeleteGroupAdminResponse, error) {
	err := s.biz.DeleteGroupAdmin(ctx, req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteGroupAdminResponse{}, nil
}

// GetGroupAdminList 获取群组管理员列表
func (s *RelationShipService) GetGroupAdminList(ctx context.Context, req *pb.GetGroupAdminListRequest) (*pb.GetGroupAdminListResponse, error) {
	list, err := s.biz.GetGroupAdminList(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}
	return &pb.GetGroupAdminListResponse{
		GroupAdmins: list,
	}, nil
}

// GetGroupAdminInfo 获取群组管理员信息
func (s *RelationShipService) GetGroupAdminInfo(ctx context.Context, req *pb.GetGroupAdminInfoRequest) (*pb.GetGroupAdminInfoResponse, error) {
	admin, err := s.biz.GetGroupAdminInfo(ctx, req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetGroupAdminInfoResponse{
		GroupAdmin: admin,
	}, nil
}

// CheckAdmin 检查是否是群组管理员
func (s *RelationShipService) CheckAdmin(ctx context.Context, req *pb.CheckAdminRequest) (*pb.CheckAdminResponse, error) {
	isAdmin, err := s.biz.CheckAdmin(ctx, req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.CheckAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}

// CheckLeader 检查是否是群组群主
func (s *RelationShipService) CheckLeader(ctx context.Context, req *pb.CheckLeaderRequest) (*pb.CheckLeaderResponse, error) {
	isLeader, err := s.biz.CheckLeader(ctx, req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.CheckLeaderResponse{
		IsLeader: isLeader,
	}, nil
}

// CheckMember 检查是否是群组成员
func (s *RelationShipService) CheckMember(ctx context.Context, req *pb.CheckMemberRequest) (*pb.CheckMemberResponse, error) {
	isMember, err := s.biz.CheckMember(ctx, req.GroupId, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.CheckMemberResponse{
		IsMember: isMember,
	}, nil
}
