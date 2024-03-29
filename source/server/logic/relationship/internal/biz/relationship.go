package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"relationship/api/v1/message"
	"relationship/api/v1/universal"
	"relationship/api/v1/user"
	"relationship/internal/common"
	"relationship/internal/common/constant"
	"relationship/internal/components/broker"
	"relationship/internal/conf"
	"relationship/internal/data/orm/model"
	"relationship/pkg"
	"sort"
	"strconv"
	"time"
)

type RelationshipRepo interface {
	CreateFriendRequest(ctx context.Context, uid, friendId int64, noteName, groupName, desc string) (*model.FriendRequest, error)
	GetFriendRequestByPage(ctx context.Context, uid int64, number, size int) ([]*model.FriendRequest, int, error)
	GetFriendRequestByRequestId(ctx context.Context, requestId int64) (*model.FriendRequest, error)
	GetFriendRequestsByRequestIds(ctx context.Context, requestId []int64) ([]*model.FriendRequest, error)
	UpdateFriendRequestStatus(ctx context.Context, requestId int64, status string) error
	CreateFriend(ctx context.Context, uid, friendId int64, noteName, groupName string) error
	DealFriendRequest(ctx context.Context, requestId int64, status, noteName, groupName string) (*model.FriendRequest, error)
	GetFriendGroups(ctx context.Context, uid int64) ([]*model.FriendGroup, error)
	GetFriendsByGroup(ctx context.Context, uid int64, groupName string) ([]*model.Friend, error)
	GetFriends(ctx context.Context, uid int64) ([]*model.Friend, error)
	DelFriend(ctx context.Context, uid, friendId int64) error
	UpdateFriendInfo(ctx context.Context, uid, friendId int64, noteName, groupName string) error
	CreateFriendGroup(ctx context.Context, uid int64, groupName string) error
	UpdateFriendGroup(ctx context.Context, uid int64, groupName, newGroupName string) error
	DeleteFriendGroup(ctx context.Context, uid int64, groupName string) error
	GetFriend(ctx context.Context, uid, friendId int64) (*model.Friend, error)
	CreateGroup(ctx context.Context, leaderId int64, groupName string, groupAvatar string, groupDesc string) (*universal.Group, error)
	GetGroupIds(ctx context.Context, userId int64) ([]string, error)
	GetGroups(ctx context.Context, groupIds []string) ([]*universal.Group, error)
	UpdateGroupInfo(ctx context.Context, groupId string, groupName string, groupAvatar string, groupDesc string) error
	DeleteGroup(ctx context.Context, groupId string) error
	GetGroupMembers(ctx context.Context, groupId string) ([]*model.GroupMember, error)
	GetGroupMember(ctx context.Context, groupId string, userId int64) (*model.GroupMember, error)
	UpdateGroupMemberInfo(ctx context.Context, groupId string, userId int64, groupNoteName string, memberNoteName string) error
	DeleteGroupMember(ctx context.Context, groupId string, userId int64) error
	CreateGroupRequest(ctx context.Context, requesterId int64, groupId string, desc string) (*model.GroupRequest, error)
	GetGroupRequests(ctx context.Context, groupId string) ([]*model.GroupRequest, error)
	GetGroupRequest(ctx context.Context, requestId int64) (*model.GroupRequest, error)
	GetGroupRequestsByIds(ctx context.Context, requestIds []int64) ([]*model.GroupRequest, error)
	CreateGroupMember(ctx context.Context, request *model.GroupRequest) error
	UpdateGroupRequestStatus(ctx context.Context, requestId int64, status string) error
	CreateAdmin(ctx context.Context, groupId string, userId int64) error
	UpdateMemberRole(ctx context.Context, groupId string, userId int64, role int) error
	DeleteAdmin(ctx context.Context, groupId string, userId int64) error
	GetGroupAdminIds(ctx context.Context, groupId string) ([]int64, error)
	GetMembersIn(ctx context.Context, id string, ids []int64) ([]*model.GroupMember, error)
	ExistAdmin(ctx context.Context, groupId string, userId int64) (bool, error)
	GetGroupLeader(ctx context.Context, groupId string) (int64, error)
	ExistMember(ctx context.Context, groupId string, userId int64) (bool, error)
	DealGroupRequest(ctx context.Context, requestId int64, status string) (*model.GroupRequest, error)
}
type RelationshipBiz struct {
	helper        *log.Helper
	repo          RelationshipRepo
	userClient    user.UserClient
	messageClient message.MessageServiceClient
	broker        *broker.KafkaBroker
	mqConfig      *conf.MessageQueue
	mysql         *gorm.DB
}

func NewRelationshipBiz(helper *log.Helper, userClient user.UserClient, messageClient message.MessageServiceClient, broker *broker.KafkaBroker, cf *conf.Bootstrap, repo RelationshipRepo, mysql *gorm.DB) *RelationshipBiz {
	return &RelationshipBiz{helper: helper, userClient: userClient, messageClient: messageClient, broker: broker, mqConfig: cf.MessageQueue, repo: repo, mysql: mysql}
}

// SendFriendRequest 发送好友请求
func (b *RelationshipBiz) SendFriendRequest(ctx context.Context, requesterId, receiverId int64, noteName, groupName, desc string) (*universal.FriendRequest, error) {
	friendReq, err := b.repo.CreateFriendRequest(ctx, requesterId, receiverId, noteName, groupName, desc)
	if err != nil {
		return nil, err
	}
	// 发送消息到mq，异步通知接收者
	m := model.FriendRequestMessage{
		RequestId: friendReq.RequestID,
		UserId:    receiverId,
		PublishAt: time.Now(),
	}
	err = b.broker.Publish(b.mqConfig.FriendRequestTopic, m)
	if err != nil {
		b.helper.Errorf("发送好友请求消息到mq失败: %v", err)
		return nil, pkg.InternalError("发送好友请求消息到mq失败: %v", err)
	}
	reply, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{Uids: []int64{receiverId}})
	if err != nil {
		b.helper.Errorf("获取用户信息失败: %v", err)
		return nil, pkg.InternalError("获取用户信息失败: %v", err)
	}
	var NickName, Avatar string
	if len(reply.Profiles) == 1 {
		NickName = reply.Profiles[0].NickName
		Avatar = reply.Profiles[0].Avatar
	}
	res := &universal.FriendRequest{
		RequestId:   friendReq.RequestID,
		RequesterId: friendReq.RequesterID,
		ReceiverId:  friendReq.ReceiverID,
		Desc:        friendReq.Desc,
		Status:      friendReq.Status,
		CreateTime:  pkg.FormatTime(friendReq.CreateAt),
		UpdateTime:  pkg.FormatTime(friendReq.UpdateAt),
		NickName:    NickName,
		Avatar:      Avatar,
	}
	return res, nil
}

// GetFriendRequestList 获取好友请求列表
func (b *RelationshipBiz) GetFriendRequestList(ctx context.Context, id int64, number int, size int) ([]*model.FriendRequest, int, error) {
	offset := (number - 1) * size
	res, count, err := b.repo.GetFriendRequestByPage(ctx, id, offset, size)
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil

}

// GetFriendRequest 获取好友请求
func (b *RelationshipBiz) GetFriendRequest(ctx context.Context, id int64) (*model.FriendRequest, error) {
	friendRequest, err := b.repo.GetFriendRequestByRequestId(ctx, id)
	if err != nil {
		return nil, err
	}
	return friendRequest, nil
}

// DealFriendRequest 处理好友请求
func (b *RelationshipBiz) DealFriendRequest(ctx context.Context, requestId int64, status, noteName, groupName string) error {
	if !common.CheckRequestStatus(status) {
		b.helper.Errorf("status 取值错误 :%v", status)
		return pkg.InvalidArgumentError("status 取值错误 :%v", status)
	}
	request, err := b.repo.GetFriendRequestByRequestId(ctx, requestId)
	if err != nil {
		return err
	}
	if request.Status != constant.Pending {
		b.helper.Errorf("该请求已处理: %v", request.Status)
		return pkg.InvalidArgumentError("该请求已处理: %v", request.Status)
	}
	req, err := b.repo.DealFriendRequest(ctx, requestId, status, noteName, groupName)
	if err != nil {
		return err
	}
	m := &model.FriendRequestMessage{
		RequestId: requestId,
		UserId:    req.RequesterID,
		PublishAt: time.Now(),
	}
	err = b.broker.Publish(b.mqConfig.FriendRequestTopic, m)
	if err != nil {
		b.helper.Errorf("发送好友请求消息到mq失败: %v", err)
		return pkg.InternalError("发送好友请求消息到mq失败: %v", err)
	}
	if status == constant.Refused {
		return nil
	}
	friendMsg := &model.FriendMessage{
		UserId:   req.RequesterID,
		FriendId: req.ReceiverID,
	}
	err = b.broker.Publish(b.mqConfig.FriendTopic, friendMsg)
	if err != nil {
		b.helper.Errorf("发送好友消息到mq失败: %v", err)
		return err
	}
	if _, err = b.messageClient.InitUnreadMessage(ctx, &message.InitUnreadMessageRequest{
		Uid:      req.RequesterID,
		FriendId: req.ReceiverID,
	}); err != nil {
		b.helper.Errorf("初始化未读消息失败: %v", err)
		return pkg.InternalError("初始化未读消息失败: %v", err)
	}
	return nil
}

// GetFriendList 获取好友列表
func (b *RelationshipBiz) GetFriendList(ctx context.Context, userId int64) ([]*universal.Friend, error) {
	friends, err := b.repo.GetFriends(ctx, userId)
	var uids []int64
	for _, friend := range friends {
		uids = append(uids, friend.FriendID)
	}
	users, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{Uids: uids})
	if err != nil {
		b.helper.Errorf("获取用户信息失败: %v", err)
		return nil, pkg.InternalError("获取用户信息失败: %v", err)
	}
	m := make(map[int64]*user.ShortProfile)
	for _, profile := range users.Profiles {
		m[profile.Uid] = profile
	}
	var res []*universal.Friend
	for _, friend := range friends {
		profile := m[friend.FriendID]
		res = append(res, &universal.Friend{
			FriendId:  friend.FriendID,
			NoteName:  friend.NoteName,
			GroupName: friend.GroupName,
			NickName:  profile.NickName,
			Avatar:    profile.Avatar,
		})
	}
	return res, nil
}

// DeleteFriend 删除好友
func (b *RelationshipBiz) DeleteFriend(ctx context.Context, id, id2 int64) error {
	err := b.repo.DelFriend(ctx, id, id2)
	if err != nil {
		return err
	}
	return nil
}

// GetFriendInfo 获取好友信息
func (b *RelationshipBiz) GetFriendInfo(ctx context.Context, id int64) (*user.GetAddressAndDescReply, error) {
	res, err := b.userClient.GetAddressAndDesc(ctx, &user.GetAddressAndDescRequest{Uid: id})
	if err != nil {
		b.helper.Errorf("获取好友信息: %v", err)
		return nil, err
	}
	return res, nil
}

// UpdateFriendInfo 更新好友信息
func (b *RelationshipBiz) UpdateFriendInfo(ctx context.Context, userId, friendId int64, noteName string, groupName string) error {
	err := b.repo.UpdateFriendInfo(ctx, userId, friendId, noteName, groupName)
	if err != nil {
		return err
	}
	return nil
}

// CreateFriendGroup 创建好友分组
func (b *RelationshipBiz) CreateFriendGroup(ctx context.Context, id int64, name string) error {
	err := b.repo.CreateFriendGroup(ctx, id, name)
	if err != nil {
		return err
	}
	return nil
}

// UpdateFriendGroup 更新好友分组
func (b *RelationshipBiz) UpdateFriendGroup(ctx context.Context, id int64, name, newName string) error {
	if name == newName {
		b.helper.Errorf("新旧分组名不能相同, name: %s, newName: %s", name, newName)
		return pkg.InvalidArgumentError("新旧分组名不能相同, name: %s, newName: %s", name, newName)
	}
	if name == constant.DefaultFriendGroup {
		b.helper.Errorf("默认分组不可修改")
		return pkg.InvalidArgumentError("默认分组不可修改")
	}
	err := b.repo.UpdateFriendGroup(ctx, id, name, newName)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFriendGroup 删除好友分组
func (b *RelationshipBiz) DeleteFriendGroup(ctx context.Context, userId int64, groupName string) error {
	if groupName == constant.DefaultFriendGroup {
		b.helper.Errorf("默认分组不可删除")
		return pkg.InvalidArgumentError("默认分组不可删除")
	}
	err := b.repo.DeleteFriendGroup(ctx, userId, groupName)
	if err != nil {
		return err
	}
	return nil
}

// GetFriendGroupList 获取好友分组列表
func (b *RelationshipBiz) GetFriendGroupList(ctx context.Context, id int64) ([]string, error) {
	res, err := b.repo.GetFriendGroups(ctx, id)
	if err != nil {
		return nil, err
	}
	var groupNames []string
	for _, group := range res {
		groupNames = append(groupNames, group.GroupName)
	}
	return groupNames, nil
}

// GetFriendRequests 获取好友请求列表
func (b *RelationshipBiz) GetFriendRequests(ctx context.Context, requestIds []int64) ([]*universal.FriendRequest, error) {
	requests, err := b.repo.GetFriendRequestsByRequestIds(ctx, requestIds)
	if err != nil {
		return nil, err
	}
	var uids []int64
	sort.Slice(requests, func(i, j int) bool {
		return requests[i].RequesterID < requests[j].RequesterID
	})
	for _, request := range requests {
		uids = append(uids, request.RequesterID)
	}
	profiles, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{Uids: uids})
	if err != nil {
		b.helper.Errorf("获取用户信息失败: %v", err)
		return nil, err
	}
	shortProfiles := profiles.Profiles
	sort.Slice(shortProfiles, func(i, j int) bool {
		return shortProfiles[i].Uid < shortProfiles[j].Uid
	})

	var res []*universal.FriendRequest
	for index, request := range requests {
		res = append(res, &universal.FriendRequest{
			RequestId:   request.RequestID,
			RequesterId: request.RequesterID,
			ReceiverId:  request.ReceiverID,
			Desc:        request.Desc,
			Status:      request.Status,
			CreateTime:  pkg.FormatTime(request.CreateAt),
			UpdateTime:  pkg.FormatTime(request.UpdateAt),
			NickName:    shortProfiles[index].NickName,
			Avatar:      shortProfiles[index].Avatar,
		})
	}
	return res, nil
}

// GetOneFriend 获取单个好友
func (b *RelationshipBiz) GetOneFriend(ctx context.Context, userId int64, friendId int64) (*universal.Friend, error) {
	friend, err := b.repo.GetFriend(ctx, userId, friendId)
	if err != nil {
		return nil, err
	}
	profile, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{
		Uids: []int64{friend.FriendID},
	})
	if err != nil {
		b.helper.Errorf("获取用户信息失败: %v", err)
		return nil, err
	}
	res := &universal.Friend{
		FriendId:  friend.FriendID,
		NickName:  profile.Profiles[0].NickName,
		NoteName:  friend.NoteName,
		Avatar:    profile.Profiles[0].Avatar,
		GroupName: friend.GroupName,
	}
	return res, nil
}

// CreateGroup 创建群组
func (b *RelationshipBiz) CreateGroup(ctx context.Context, leaderId int64, groupName string, groupAvatar string, groupDesc string) (*universal.Group, error) {
	group, err := b.repo.CreateGroup(ctx, leaderId, groupName, groupAvatar, groupDesc)
	if err != nil {
		return nil, err
	}
	return group, nil
}

// GetGroupList 获取群组列表
func (b *RelationshipBiz) GetGroupList(ctx context.Context, userId int64) ([]*universal.Group, error) {
	groupIds, err := b.repo.GetGroupIds(ctx, userId)
	if err != nil {
		return nil, err
	}
	groups, err := b.repo.GetGroups(ctx, groupIds)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// GetGroupInfo 获取群组信息
func (b *RelationshipBiz) GetGroupInfo(ctx context.Context, groupId string) (*universal.Group, error) {
	groups, err := b.repo.GetGroups(ctx, []string{groupId})
	if err != nil {
		return nil, err
	}
	if len(groups) == 0 {
		b.helper.Errorf("该群组不存在: %s", groupId)
		return nil, pkg.InternalError("该群组不存在: %s", groupId)
	}
	return groups[0], nil
}

// UpdateGroupInfo 更新群组信息
func (b *RelationshipBiz) UpdateGroupInfo(ctx context.Context, groupId string, groupName string, groupAvatar string, groupDesc string) error {
	err := b.repo.UpdateGroupInfo(ctx, groupId, groupName, groupAvatar, groupDesc)
	if err != nil {
		return err
	}
	return nil
}

// DeleteGroup 删除群组, 只有群主才能删除群组,user_id必须是群主
func (b *RelationshipBiz) DeleteGroup(ctx context.Context, groupId string, userId int64) error {
	groups, err := b.repo.GetGroups(ctx, []string{groupId})
	if len(groups) == 0 {
		b.helper.Errorf("该群组不存在: %s", groupId)
		return pkg.InternalError("该群组不存在: %s", groupId)
	}
	group := groups[0]
	if group.GroupLeaderId != userId {
		b.helper.Errorf("非群主不能删除群组: %s", groupId)
		return pkg.InternalError("非群主不能删除群组: %s", groupId)
	}
	err = b.repo.DeleteGroup(ctx, groupId)
	if err != nil {
		return err
	}
	return nil
}

// GetGroupMemberList 获取群组成员列表
func (b *RelationshipBiz) GetGroupMemberList(ctx context.Context, groupId string) ([]*universal.GroupMember, error) {
	members, err := b.repo.GetGroupMembers(ctx, groupId)
	if err != nil {
		b.helper.Errorf(err.Error())
		return nil, err
	}
	var uids []int64
	for _, member := range members {
		uids = append(uids, member.MemberID)
	}
	profiles, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{Uids: uids})
	if err != nil {
		b.helper.Errorf("获取用户信息失败: %v", err)
		return nil, pkg.InternalError("获取用户信息失败: %v", err)
	}
	m := make(map[int64]*user.ShortProfile)
	for _, profile := range profiles.Profiles {
		m[profile.Uid] = profile
	}
	var res []*universal.GroupMember
	for _, member := range members {
		profile := m[member.MemberID]
		res = append(res, &universal.GroupMember{
			GroupId:        groupId,
			MemberId:       member.MemberID,
			MemberNoteName: member.MemberNoteName,
			MemberAvatar:   profile.Avatar,
			JoinAt:         member.BecomeAt.String(),
			MemberRole:     strconv.FormatInt(int64(member.Role), 10),
		})
	}
	return res, nil
}

// GetGroupMemberInfo 获取群组成员信息
func (b *RelationshipBiz) GetGroupMemberInfo(ctx context.Context, groupId string, userId int64) (*universal.GroupMember, error) {
	member, err := b.repo.GetGroupMember(ctx, groupId, userId)
	if err != nil {
		return nil, err
	}
	profiles, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{Uids: []int64{userId}})
	if err != nil {
		b.helper.Errorf("获取用户信息失败: %v", err)
		return nil, pkg.InternalError("获取用户信息失败: %v", err)
	}
	if len(profiles.Profiles) == 0 {
		b.helper.Errorf("该用户不存在: %d", userId)
		return nil, pkg.InternalError("该用户不存在: %d", userId)
	}
	profile := profiles.Profiles[0]
	res := &universal.GroupMember{
		GroupId:        groupId,
		MemberId:       userId,
		MemberNoteName: member.MemberNoteName,
		MemberAvatar:   profile.Avatar,
		JoinAt:         member.BecomeAt.String(),
		MemberRole:     strconv.FormatInt(int64(member.Role), 10),
	}
	return res, nil
}

// UpdateGroupMemberInfo 更新群组成员信息
func (b *RelationshipBiz) UpdateGroupMemberInfo(ctx context.Context, groupId string, userId int64, groupNoteName string, memberNoteName string) error {
	err := b.repo.UpdateGroupMemberInfo(ctx, groupId, userId, groupNoteName, memberNoteName)
	if err != nil {
		return err
	}
	return nil
}

// DeleteGroupMember 删除群组成员
func (b *RelationshipBiz) DeleteGroupMember(ctx context.Context, groupId string, userId int64) error {
	isMember, err := b.CheckMember(ctx, groupId, userId)
	if err != nil {
		return err
	}
	if !isMember {
		b.helper.Warnf("用户：%d 不是群组：%s 成员", userId, groupId)
		return pkg.InternalError("用户：%d 不是群组：%s 成员", userId, groupId)
	}
	err = b.repo.DeleteGroupMember(ctx, groupId, userId)
	if err != nil {
		return err
	}
	return nil
}

// SendGroupRequest 发送入群申请
func (b *RelationshipBiz) SendGroupRequest(ctx context.Context, requesterId int64, groupId string, desc string) (*universal.GroupRequest, error) {
	request, err := b.repo.CreateGroupRequest(ctx, requesterId, groupId, desc)
	if err != nil {
		return nil, err
	}
	// 发送通知
	if err = b.broker.Publish(b.mqConfig.GroupRequestTopic, &model.GroupRequestMessage{
		RequestId: request.RequestID,
	}); err != nil {
		b.helper.Errorf("发送入群申请通知失败: %v", err)
		return nil, pkg.InternalError("发送入群申请通知失败: %v", err)
	}
	groupInfo, err := b.GetGroupInfo(ctx, groupId)
	if err != nil {
		return nil, err
	}

	return &universal.GroupRequest{
		RequestId:   request.RequestID,
		RequesterId: requesterId,
		GroupId:     groupId,
		Desc:        desc,
		Status:      request.Status,
		CreateAt:    request.CreateAt.String(),
		UpdateAt:    request.UpdateAt.String(),
		NickName:    groupInfo.GroupName,   //这里指的是群昵称
		Avatar:      groupInfo.GroupAvatar, //这里指的是群头像
	}, nil
}

// GetGroupRequestList 获取入群申请列表,只获取未处理的申请
func (b *RelationshipBiz) GetGroupRequestList(ctx context.Context, groupId string) ([]*universal.GroupRequest, error) {
	requests, err := b.repo.GetGroupRequests(ctx, groupId)
	if err != nil {
		return nil, err
	}
	var uids []int64
	for _, request := range requests {
		uids = append(uids, request.RequesterID)
	}
	profiles, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{Uids: uids})
	if err != nil {
		b.helper.Errorf("获取用户信息失败: %v", err)
		return nil, pkg.InternalError("获取用户信息失败: %v", err)
	}
	m := make(map[int64]*user.ShortProfile)
	for _, profile := range profiles.Profiles {
		m[profile.Uid] = profile
	}
	var res []*universal.GroupRequest
	for _, request := range requests {
		profile := m[request.RequesterID]
		res = append(res, &universal.GroupRequest{
			RequestId:   request.RequestID,
			RequesterId: request.RequesterID,
			GroupId:     request.GroupID,
			Desc:        request.Desc,
			Status:      request.Status,
			CreateAt:    request.CreateAt.String(),
			UpdateAt:    request.UpdateAt.String(),
			NickName:    profile.NickName,
			Avatar:      profile.Avatar,
		})
	}
	return res, nil
}

// GetGroupRequest 获取入群申请详情
func (b *RelationshipBiz) GetGroupRequest(ctx context.Context, requestId int64) (*universal.GroupRequest, error) {
	request, err := b.repo.GetGroupRequest(ctx, requestId)
	if err != nil {
		return nil, err
	}
	profiles, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{Uids: []int64{request.RequesterID}})
	if err != nil {
		b.helper.Errorf("获取用户信息失败: %v", err)
		return nil, pkg.InternalError("获取用户信息失败: %v", err)
	}
	if len(profiles.Profiles) == 0 {
		b.helper.Errorf("该用户不存在：%d", request.RequesterID)
		return nil, pkg.InternalError("该用户不存在：%d", request.RequesterID)
	}
	profile := profiles.Profiles[0]
	return &universal.GroupRequest{
		RequestId:   request.RequestID,
		RequesterId: request.RequesterID,
		GroupId:     request.GroupID,
		Desc:        request.Desc,
		Status:      request.Status,
		CreateAt:    request.CreateAt.String(),
		UpdateAt:    request.UpdateAt.String(),
		NickName:    profile.NickName,
		Avatar:      profile.Avatar,
	}, nil
}

// GetGroupRequests 获取入群申请详情
func (b *RelationshipBiz) GetGroupRequests(ctx context.Context, requestIds []int64) ([]*universal.GroupRequest, error) {
	requests, err := b.repo.GetGroupRequestsByIds(ctx, requestIds)
	if err != nil {

		return nil, err
	}
	var uids []int64
	for _, request := range requests {
		uids = append(uids, request.RequesterID)
	}
	profiles, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{Uids: uids})
	if err != nil {
		b.helper.Errorf("获取用户信息失败: %v", err)
		return nil, pkg.InternalError("获取用户信息失败: %v", err)
	}
	m := make(map[int64]*user.ShortProfile)
	for _, profile := range profiles.Profiles {
		m[profile.Uid] = profile
	}
	var res []*universal.GroupRequest
	for _, request := range requests {
		profile := m[request.RequesterID]
		res = append(res, &universal.GroupRequest{
			RequestId:   request.RequestID,
			RequesterId: request.RequesterID,
			GroupId:     request.GroupID,
			Desc:        request.Desc,
			Status:      request.Status,
			CreateAt:    request.CreateAt.String(),
			UpdateAt:    request.UpdateAt.String(),
			NickName:    profile.NickName,
			Avatar:      profile.Avatar,
		})
	}
	return res, nil
}

// DealGroupRequest 处理入群申请
func (b *RelationshipBiz) DealGroupRequest(ctx context.Context, requestId int64, status string) error {
	if !common.CheckRequestStatus(status) {
		b.helper.Errorf("status 取值错误 :%v", status)
		return pkg.InvalidArgumentError("status 取值错误 :%v", status)
	}
	request, err := b.repo.GetGroupRequest(ctx, requestId)
	if err != nil {
		return err
	}
	if request.Status != constant.Pending {
		b.helper.Errorf("该申请已处理")
		return pkg.InvalidArgumentError("该申请已处理")
	}
	req, err := b.repo.DealGroupRequest(ctx, requestId, status)
	// 发送通知
	if err = b.broker.Publish(b.mqConfig.GroupRequestTopic, &model.GroupRequestMessage{
		RequestId:   requestId,
		RequesterId: req.RequesterID,
	}); err != nil {
		b.helper.Errorf("发送通知失败: %v", err)
		return pkg.InternalError("发送通知失败: %v", err)
	}
	return nil
}

// CreateGroupAdmin 设置群管理员
func (b *RelationshipBiz) CreateGroupAdmin(ctx context.Context, groupId string, userId int64) error {
	member, err := b.repo.GetGroupMember(ctx, groupId, userId)
	if err != nil {
		return err
	}
	if member == nil {
		b.helper.Errorf("该成员不存在: %d", userId)
		return pkg.InternalError("该成员不存在: %d", userId)
	}
	if member.Role == constant.Admin {
		b.helper.Errorf("该成员已经是管理员: %d", userId)
		return pkg.InternalError("该成员已经是管理员: %d", userId)
	}
	if member.Role == constant.Leader {
		b.helper.Errorf("该成员是群主: %d", userId)
		return pkg.InternalError("该成员是群主: %d", userId)
	}
	// 更新成员角色
	if err = b.repo.CreateAdmin(ctx, groupId, userId); err != nil {
		return err
	}
	return nil
}

// DeleteGroupAdmin 取消群管理员
func (b *RelationshipBiz) DeleteGroupAdmin(ctx context.Context, groupId string, userId int64) error {
	member, err := b.repo.GetGroupMember(ctx, groupId, userId)
	if err != nil {
		return err
	}
	if member == nil {
		b.helper.Errorf("该成员不存在: %d", userId)
		return pkg.InternalError("该成员不存在: %d", userId)
	}
	if member.Role != constant.Admin {
		b.helper.Errorf("该成员不是管理员: %d", userId)
		return pkg.InternalError("该成员不是管理员: %d", userId)
	}
	// 更新成员角色
	if err = b.repo.DeleteAdmin(ctx, groupId, userId); err != nil {
		return err
	}
	return nil
}

// GetGroupAdminList 获取群管理员列表
func (b *RelationshipBiz) GetGroupAdminList(ctx context.Context, groupId string) ([]*universal.GroupMember, error) {
	adminIds, err := b.repo.GetGroupAdminIds(ctx, groupId)
	if err != nil {
		return nil, err
	}
	members, err := b.repo.GetMembersIn(ctx, groupId, adminIds)
	if err != nil {
		return nil, err
	}
	var uids []int64
	for _, member := range members {
		uids = append(uids, member.MemberID)
	}
	profiles, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{Uids: uids})
	if err != nil {
		b.helper.Errorf("获取用户信息失败: %v", err)
		return nil, pkg.InternalError("获取用户信息失败: %v", err)
	}
	m := make(map[int64]*user.ShortProfile)
	for _, profile := range profiles.Profiles {
		m[profile.Uid] = profile
	}
	var res []*universal.GroupMember
	for _, member := range members {
		profile := m[member.MemberID]
		res = append(res, &universal.GroupMember{
			GroupId:        member.GroupID,
			MemberId:       member.MemberID,
			MemberNoteName: member.MemberNoteName,
			MemberAvatar:   profile.Avatar,
			JoinAt:         member.BecomeAt.String(),
			MemberRole:     pkg.TransferIntToString(member.Role),
		})
	}
	return res, nil
}

// CompleteMembers 补全群成员信息, 从用户服务获取
func (b *RelationshipBiz) CompleteMembers(members []*model.GroupMember) ([]*universal.GroupMember, error) {
	var uids []int64
	for _, member := range members {
		uids = append(uids, member.MemberID)
	}
	profiles, err := b.userClient.GetProfiles(context.Background(), &user.GetProfilesRequest{Uids: uids})
	if err != nil {
		b.helper.Errorf("获取用户信息失败: %v", err)
		return nil, pkg.InternalError("获取用户信息失败: %v", err)
	}
	m := make(map[int64]*user.ShortProfile)
	for _, profile := range profiles.Profiles {
		m[profile.Uid] = profile
	}
	var res []*universal.GroupMember
	for _, member := range members {
		profile := m[member.MemberID]
		res = append(res, &universal.GroupMember{
			GroupId:        member.GroupID,
			MemberId:       member.MemberID,
			MemberNoteName: member.MemberNoteName,
			MemberAvatar:   profile.Avatar,
			JoinAt:         member.BecomeAt.String(),
			MemberRole:     pkg.TransferIntToString(member.Role),
		})
	}
	return res, nil
}

// CompleteMember 补全群成员信息, 从用户服务获取
func (b *RelationshipBiz) CompleteMember(member *model.GroupMember) (*universal.GroupMember, error) {
	members, err := b.CompleteMembers([]*model.GroupMember{member})
	if err != nil {
		return nil, err
	}
	if len(members) == 0 {
		b.helper.Errorf("该成员不存在: %d", member.MemberID)
		return nil, pkg.InternalError("该成员不存在: %d", member.MemberID)
	}
	return members[0], nil
}

// GetGroupAdminInfo 获取群管理员信息
func (b *RelationshipBiz) GetGroupAdminInfo(ctx context.Context, groupId string, userId int64) (*universal.GroupMember, error) {
	member, err := b.repo.GetGroupMember(ctx, groupId, userId)
	if err != nil {
		return nil, err
	}
	if member == nil {
		b.helper.Errorf("该成员不存在: %d", userId)
		return nil, pkg.InternalError("该成员不存在: %d", userId)
	}
	if member.Role != constant.Admin {
		b.helper.Errorf("该成员不是管理员: %d", userId)
		return nil, pkg.InternalError("该成员不是管理员: %d", userId)
	}
	res, err := b.CompleteMember(member)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CheckAdmin 检查是否是群管理员
func (b *RelationshipBiz) CheckAdmin(ctx context.Context, groupId string, userId int64) (bool, error) {
	isAdmin, err := b.repo.ExistAdmin(ctx, groupId, userId)
	if err != nil {
		return false, err
	}
	return isAdmin, nil
}

// CheckLeader 检查是否是群主
func (b *RelationshipBiz) CheckLeader(ctx context.Context, groupId string, userId int64) (bool, error) {
	leaderId, err := b.repo.GetGroupLeader(ctx, groupId)
	if err != nil {
		return false, err
	}
	return leaderId == userId, nil
}

// CheckMember 检查是否是群成员
func (b *RelationshipBiz) CheckMember(ctx context.Context, groupId string, userId int64) (bool, error) {
	isMember, err := b.repo.ExistMember(ctx, groupId, userId)
	if err != nil {
		return false, err
	}
	return isMember, nil
}
