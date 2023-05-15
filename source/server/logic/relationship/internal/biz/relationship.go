package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"relationship/api/v1/relationship"
	"relationship/api/v1/universal"
	"relationship/api/v1/user"
	"relationship/internal/components/broker"
	"relationship/internal/conf"
	"relationship/internal/data/orm/model"
	"relationship/pkg"
	"sort"
	"time"
)

type RelationshipRepo interface {
	CreateFriendRequest(ctx context.Context, uid, friendId int64, noteName, groupName, desc string) (*model.FriendRequest, error)
	FindAllFriendRequestByPage(ctx context.Context, uid int64, number, size int) ([]*model.FriendRequest, int, error)
	FindFriendRequest(ctx context.Context, requestId int64) (*model.FriendRequest, error)
	FindFriendRequests(ctx context.Context, requestId []int64) ([]*model.FriendRequest, error)
	UpdateFriendRequestStatus(ctx context.Context, requestId int64, status string) error
	CreateFriend(ctx context.Context, uid, friendId int64, noteName, groupName string) error
	DealFriendRequest(ctx context.Context, requestId int64, status string) (int64, error)
	GetFriendGroups(ctx context.Context, uid int64) ([]*model.FriendGroup, error)
	GetFriendsByGroup(ctx context.Context, uid int64, groupName string) ([]*model.Friend, error)
	GetFriends(ctx context.Context, uid int64) ([]*model.Friend, error)
	DelFriend(ctx context.Context, uid, friendId int64) error
	UpdateFriendInfo(ctx context.Context, uid, friendId int64, noteName, groupName string) error
	CreateFriendGroup(ctx context.Context, uid int64, groupName string) error
	UpdateFriendGroup(ctx context.Context, uid int64, groupName, newGroupName string) error
	DeleteFriendGroup(ctx context.Context, uid int64, groupName string) error
}
type RelationshipBiz struct {
	helper     *log.Helper
	repo       RelationshipRepo
	userClient user.UserClient
	broker     *broker.KafkaBroker
	mqConfig   *conf.MessageQueue
	mysql      *gorm.DB
}

func NewRelationshipBiz(helper *log.Helper, userClient user.UserClient, broker *broker.KafkaBroker, cf *conf.Bootstrap, repo RelationshipRepo, mysql *gorm.DB) *RelationshipBiz {
	return &RelationshipBiz{helper: helper, userClient: userClient, broker: broker, mqConfig: cf.MessageQueue, repo: repo, mysql: mysql}
}

func (b *RelationshipBiz) SendFriendRequest(ctx context.Context, requesterId, receiverId int64, noteName, groupName, desc string) (*universal.FriendRequest, error) {
	friendReq, err := b.repo.CreateFriendRequest(ctx, requesterId, receiverId, noteName, groupName, desc)
	if err != nil {
		b.helper.Errorf("send friend request error: %v", err)
		return nil, err
	}
	// 发送消息到mq，异步通知接收者
	message := model.FriendRequestMessage{
		RequestId: friendReq.RequestID,
		UserId:    receiverId,
		PublishAt: time.Now(),
	}
	err = b.broker.Publish(b.mqConfig.FriendRequestTopic, message)
	if err != nil {
		b.helper.Errorf("send friend request to mq error: %v", err)
		return nil, err
	}
	reply, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{Uids: []int64{receiverId}})
	if err != nil {
		b.helper.Errorf("get profiles error: %v", err)
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

func (b *RelationshipBiz) GetFriendRequestList(ctx context.Context, id int64, number int, size int) ([]*model.FriendRequest, int, error) {
	offset := (number - 1) * size
	res, count, err := b.repo.FindAllFriendRequestByPage(ctx, id, offset, size)
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil

}

func (b *RelationshipBiz) GetFriendRequest(ctx context.Context, id int64) (*model.FriendRequest, error) {
	friendRequest, err := b.repo.FindFriendRequest(ctx, id)
	if err != nil {
		b.helper.Errorf("get friend request error: %v", err)
		return nil, err
	}
	return friendRequest, nil
}

func (b *RelationshipBiz) DealFriendRequest(ctx context.Context, id int64, status string) error {
	userId, err := b.repo.DealFriendRequest(ctx, id, status)
	if err != nil {
		b.helper.Errorf("deal friend request error: %v", err)
		return err
	}
	message := &model.FriendRequestMessage{
		RequestId: id,
		UserId:    userId,
		PublishAt: time.Now(),
	}
	err = b.broker.Publish(b.mqConfig.FriendRequestTopic, message)
	if err != nil {
		b.helper.Errorf("send deal friend request to mq error: %v", err)
		return err
	}
	return nil
}

func (b *RelationshipBiz) GetFriendList(ctx context.Context, userId int64) ([]*relationship.FriendGroup, error) {
	friendGroups, err := b.repo.GetFriendGroups(ctx, userId)
	if err != nil {
		b.helper.Errorf("get friend groups error: %v", err)
		return nil, err
	}
	friends, err := b.repo.GetFriends(ctx, userId)
	var uids []int64
	for _, friend := range friends {
		uids = append(uids, friend.FriendID)
	}
	users, err := b.userClient.GetProfiles(ctx, &user.GetProfilesRequest{Uids: uids})
	if err != nil {
		b.helper.Errorf("get profiles error: %v", err)
		return nil, err
	}
	m := make(map[int64]*user.ShortProfile)
	for _, profile := range users.Profiles {
		m[profile.Uid] = profile
	}
	var res []*relationship.FriendGroup
	for _, group := range friendGroups {
		var fs []*relationship.Friend
		for _, friend := range friends {
			fs = append(fs, &relationship.Friend{
				FriendId: friend.FriendID,
				NickName: m[friend.FriendID].NickName,
				NoteName: friend.NoteName,
				Avatar:   m[friend.FriendID].Avatar,
			})
		}
		res = append(res, &relationship.FriendGroup{
			GroupName: group.GroupName,
			Friends:   fs,
		})
	}
	return res, nil
}

func (b *RelationshipBiz) DeleteFriend(ctx context.Context, id, id2 int64) error {
	err := b.repo.DelFriend(ctx, id, id2)
	if err != nil {
		b.helper.Errorf("delete friend error: %v", err)
		return err
	}
	return nil
}

func (b *RelationshipBiz) GetFriendInfo(ctx context.Context, id int64) (*user.GetAddressAndDescReply, error) {
	res, err := b.userClient.GetAddressAndDesc(ctx, &user.GetAddressAndDescRequest{Uid: id})
	if err != nil {
		b.helper.Errorf("get address and desc error: %v", err)
		return nil, err
	}
	return res, nil
}

func (b *RelationshipBiz) UpdateFriendInfo(ctx context.Context, userId, friendId int64, noteName string, groupName string) error {
	err := b.repo.UpdateFriendInfo(ctx, userId, friendId, noteName, groupName)
	if err != nil {
		b.helper.Errorf("update friend info error: %v", err)
		return err
	}
	return nil
}

func (b *RelationshipBiz) CreateFriendGroup(ctx context.Context, id int64, name string) error {
	err := b.repo.CreateFriendGroup(ctx, id, name)
	if err != nil {
		b.helper.Errorf("create friend group error: %v", err)
		return err
	}
	return nil
}

func (b *RelationshipBiz) UpdateFriendGroup(ctx context.Context, id int64, name, newName string) error {
	err := b.repo.UpdateFriendGroup(ctx, id, name, newName)
	if err != nil {
		b.helper.Errorf("update friend group error: %v", err)
		return err
	}
	return nil
}

func (b *RelationshipBiz) DeleteFriendGroup(ctx context.Context, userId int64, groupName string) error {
	err := b.repo.DeleteFriendGroup(ctx, userId, groupName)
	if err != nil {
		b.helper.Errorf("delete friend group error: %v", err)
		return err
	}
	return nil
}

func (b *RelationshipBiz) GetFriendGroupList(ctx context.Context, id int64) ([]string, error) {
	res, err := b.repo.GetFriendGroups(ctx, id)
	if err != nil {
		b.helper.Errorf("get friend group list error: %v", err)
		return nil, err
	}
	var groupNames []string
	for _, group := range res {
		groupNames = append(groupNames, group.GroupName)
	}
	return groupNames, nil
}

func (b *RelationshipBiz) GetFriendRequests(ctx context.Context, requestIds []int64) ([]*universal.FriendRequest, error) {
	requests, err := b.repo.FindFriendRequests(ctx, requestIds)
	if err != nil {
		b.helper.Errorf("get friend requests error: %v", err)
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
		b.helper.Errorf("get profiles error: %v", err)
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
