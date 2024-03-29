package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"relationship/api/v1/universal"
	"relationship/internal/biz"
	"relationship/internal/common/constant"
	"relationship/internal/components/uid"
	"relationship/internal/data/orm/dal"
	"relationship/internal/data/orm/model"
	"relationship/pkg"
	"time"
)

type RelationshipRepoImpl struct {
	helper             *log.Helper
	friendUidGenerator *uid.FriendRequestUidGenerator
	groupUidGenerator  *uid.GroupRequestUidGenerator
}

func (r *RelationshipRepoImpl) DealGroupRequest(ctx context.Context, requestId int64, status string) (*model.GroupRequest, error) {
	var groupRequest *model.GroupRequest
	var err error
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		if err := pkg.ContextErr(ctx); err != nil {
			return err
		}
		// 更新请求状态
		_, err = tx.WithContext(ctx).GroupRequest.
			Where(tx.FriendRequest.RequestID.Eq(requestId)).
			Update(tx.GroupRequest.Status, status)
		if err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		// 拒绝则直接返回
		if status == constant.Refused {
			return nil
		}
		// 同意则添加群成员,先获取请求信息
		groupRequest, err = tx.GroupRequest.WithContext(ctx).
			Where(tx.GroupRequest.RequestID.Eq(requestId)).
			First()
		if err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		member := &model.GroupMember{
			MemberID:       groupRequest.RequesterID,
			GroupID:        groupRequest.GroupID,
			GroupNoteName:  "",
			MemberNoteName: "",
			BecomeAt:       time.Now(),
			Role:           0,
		}
		// 添加群成员
		if err = tx.GroupMember.WithContext(ctx).
			Create(member); err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		return nil
	})
	if err != nil {
		r.helper.Errorf(constant.TransactionErrorFormat, err)
		return nil, pkg.InternalError(constant.TransactionErrorFormat, err)
	}
	return groupRequest, err
}

func (r *RelationshipRepoImpl) ExistMember(ctx context.Context, groupId string, userId int64) (bool, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return false, err
	}
	if count, err := dal.GroupMember.WithContext(ctx).
		Where(
			dal.GroupMember.GroupID.Eq(groupId),
			dal.GroupMember.MemberID.Eq(userId),
			dal.GroupMember.IsDeleted.Eq(constant.NotDeleted),
		).
		Count(); err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return false, pkg.InternalError(constant.SqlErrorFormat, err)
	} else {
		return count > 0, nil
	}
}

// ExistAdmin 判断是否为群管理员
func (r *RelationshipRepoImpl) ExistAdmin(ctx context.Context, groupId string, userId int64) (bool, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return false, err
	}
	if count, err := dal.GroupAdmin.WithContext(ctx).
		Where(
			dal.GroupAdmin.GroupID.Eq(groupId),
			dal.GroupAdmin.AdministratorID.Eq(userId),
			dal.GroupAdmin.IsDeleted.Eq(constant.NotDeleted),
		).
		Count(); err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return false, pkg.InternalError(constant.SqlErrorFormat, err)
	} else {
		return count > 0, nil
	}
}

// GetGroupLeader 获取群主
func (r *RelationshipRepoImpl) GetGroupLeader(ctx context.Context, groupId string) (int64, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return 0, err
	}
	var group model.Group
	if _, err := dal.Group.WithContext(ctx).
		Where(
			dal.Group.GroupID.Eq(groupId),
			dal.Group.IsDeleted.Eq(constant.NotDeleted),
		).
		First(); pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return 0, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return group.GroupLeaderID, nil
}

func (r *RelationshipRepoImpl) GetMembersIn(ctx context.Context, groupId string, adminIds []int64) ([]*model.GroupMember, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	members, err := dal.GroupMember.WithContext(ctx).
		Where(
			dal.GroupMember.GroupID.Eq(groupId),
			dal.GroupMember.MemberID.In(adminIds...),
			dal.GroupMember.IsDeleted.Eq(constant.NotDeleted),
		).
		Find()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)

		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return members, nil
}

func (r *RelationshipRepoImpl) GetGroupAdminIds(ctx context.Context, groupId string) ([]int64, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	var adminIds []int64
	if admins, err := dal.GroupAdmin.WithContext(ctx).
		Select(dal.GroupAdmin.AdministratorID).
		Where(
			dal.GroupAdmin.GroupID.Eq(groupId),
			dal.GroupAdmin.IsDeleted.Eq(constant.NotDeleted),
		).
		Find(); pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)

		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	} else {
		for _, admin := range admins {
			adminIds = append(adminIds, admin.AdministratorID)
		}
	}
	return adminIds, nil
}

func (r *RelationshipRepoImpl) DeleteAdmin(ctx context.Context, groupId string, userId int64) error {
	if err := dal.Q.Transaction(func(tx *dal.Query) error {
		//逻辑删除管理员
		if _, err := tx.GroupAdmin.WithContext(ctx).
			Where(
				dal.GroupAdmin.GroupID.Eq(groupId),
				dal.GroupAdmin.AdministratorID.Eq(userId),
			).
			Update(dal.GroupAdmin.IsDeleted, constant.Deleted); err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		//更新群成员权限
		if _, err := tx.GroupMember.WithContext(ctx).
			Where(
				dal.GroupMember.GroupID.Eq(groupId),
				dal.GroupMember.MemberID.Eq(userId),
			).
			Update(dal.GroupMember.Role, constant.Normal); err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		return nil
	}); err != nil {
		r.helper.Errorf(constant.TransactionErrorFormat, err)
		return pkg.InternalError(constant.TransactionErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) UpdateMemberRole(ctx context.Context, groupId string, userId int64, role int) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	if _, err := dal.GroupMember.WithContext(ctx).
		Where(
			dal.GroupMember.GroupID.Eq(groupId),
			dal.GroupMember.MemberID.Eq(userId),
		).
		Update(dal.GroupMember.Role, role); err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) CreateAdmin(ctx context.Context, groupId string, userId int64) error {
	//	开启事务
	if err := dal.Q.Transaction(func(tx *dal.Query) error {
		if err := pkg.ContextErr(ctx); err != nil {
			return err
		}
		//创建管理员
		if err := tx.GroupAdmin.WithContext(ctx).Create(&model.GroupAdmin{
			GroupID:         groupId,
			AdministratorID: userId,
		}); err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		if err := pkg.ContextErr(ctx); err != nil {
			return err
		}
		//更新群成员权限
		if _, err := tx.GroupMember.WithContext(ctx).
			Where(
				dal.GroupMember.GroupID.Eq(groupId),
				dal.GroupMember.MemberID.Eq(userId),
			).
			Update(dal.GroupMember.Role, constant.Admin); err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		return nil
	}); err != nil {
		r.helper.Errorf(constant.TransactionErrorFormat, err)
		return pkg.InternalError(constant.TransactionErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) UpdateGroupRequestStatus(ctx context.Context, requestId int64, status string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	if _, err := dal.GroupRequest.WithContext(ctx).
		Where(dal.GroupRequest.RequestID.Eq(requestId)).
		Update(dal.GroupRequest.Status, status); err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return err
	}
	return nil
}

func (r *RelationshipRepoImpl) CreateGroupMember(ctx context.Context, request *model.GroupRequest) error {
	m := &model.GroupMember{
		MemberID:  request.RequesterID,
		GroupID:   request.GroupID,
		BecomeAt:  time.Now(),
		IsDeleted: 0,
		Role:      constant.Normal,
	}
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	if err := dal.GroupMember.WithContext(ctx).Create(m); err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) GetGroupRequestsByIds(ctx context.Context, requestIds []int64) ([]*model.GroupRequest, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	requests, err := dal.GroupRequest.WithContext(ctx).
		Where(dal.GroupRequest.RequestID.In(requestIds...)).
		Find()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return requests, nil
}

func (r *RelationshipRepoImpl) GetGroupRequest(ctx context.Context, requestId int64) (*model.GroupRequest, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	request, err := dal.GroupRequest.WithContext(ctx).
		Where(dal.GroupRequest.RequestID.Eq(requestId)).
		First()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return request, nil
}

func (r *RelationshipRepoImpl) GetGroupRequests(ctx context.Context, groupId string) ([]*model.GroupRequest, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	requests, err := dal.GroupRequest.WithContext(ctx).
		Where(
			dal.GroupRequest.GroupID.Eq(groupId),
			dal.GroupRequest.Status.Eq(constant.Pending),
		).
		Find()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return requests, nil
}

func (r *RelationshipRepoImpl) CreateGroupRequest(ctx context.Context, requesterId int64, groupId string, desc string) (*model.GroupRequest, error) {
	m := &model.GroupRequest{
		RequestID:   r.friendUidGenerator.Generate().Int64(),
		RequesterID: requesterId,
		GroupID:     groupId,
		Desc:        desc,
		Status:      constant.Pending,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	if err := dal.GroupRequest.WithContext(ctx).
		Create(m); err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return m, nil
}

func (r *RelationshipRepoImpl) DeleteGroupMember(ctx context.Context, groupId string, userId int64) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	if _, err := dal.GroupMember.WithContext(ctx).
		Where(
			dal.GroupMember.GroupID.Eq(groupId),
			dal.GroupMember.MemberID.Eq(userId),
			dal.GroupMember.IsDeleted.Eq(constant.NotDeleted),
		).
		Update(dal.GroupMember.IsDeleted, constant.Deleted); err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return err
	}
	return nil
}

func (r *RelationshipRepoImpl) UpdateGroupMemberInfo(ctx context.Context, groupId string, userId int64, groupNoteName string, memberNoteName string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	m := &model.GroupMember{
		GroupNoteName:  groupNoteName,
		MemberNoteName: memberNoteName,
	}
	if _, err := dal.GroupMember.WithContext(ctx).
		Where(
			dal.GroupMember.GroupID.Eq(groupId),
			dal.GroupMember.MemberID.Eq(userId),
			dal.GroupMember.IsDeleted.Eq(constant.NotDeleted),
		).
		Updates(m); err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) GetGroupMember(ctx context.Context, groupId string, userId int64) (*model.GroupMember, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	member, err := dal.GroupMember.WithContext(ctx).
		Where(
			dal.GroupMember.GroupID.Eq(groupId),
			dal.GroupMember.MemberID.Eq(userId),
			dal.GroupMember.IsDeleted.Eq(constant.NotDeleted),
		).First()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return member, nil
}

func (r *RelationshipRepoImpl) GetGroupMembers(ctx context.Context, groupId string) ([]*model.GroupMember, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	members, err := dal.GroupMember.WithContext(ctx).
		Where(
			dal.GroupMember.GroupID.Eq(groupId),
			dal.GroupMember.IsDeleted.Eq(constant.NotDeleted),
		).
		Select(
			dal.GroupMember.MemberID,
			dal.GroupMember.MemberNoteName,
			dal.GroupMember.BecomeAt,
			dal.GroupMember.Role,
		).
		Find()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return members, nil
}

// DeleteGroup 删除群组
func (r *RelationshipRepoImpl) DeleteGroup(ctx context.Context, groupId string) error {
	//开启事务
	if err := dal.Q.Transaction(func(tx *dal.Query) error {
		if _, err := tx.GroupMember.WithContext(ctx).
			Where(tx.GroupMember.GroupID.Eq(groupId)).
			Update(dal.GroupMember.IsDeleted, constant.Deleted); err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		if _, err := tx.Group.WithContext(ctx).
			Where(tx.Group.GroupID.Eq(groupId)).
			Update(dal.Group.IsDeleted, constant.Deleted); err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		return nil
	}); err != nil {
		r.helper.Errorf(constant.TransactionErrorFormat, err)
		return pkg.InternalError(constant.TransactionErrorFormat, err)
	}
	return nil
}

// UpdateGroupInfo 更新群组信息
func (r *RelationshipRepoImpl) UpdateGroupInfo(ctx context.Context, groupId string, groupName string, groupAvatar string, groupDesc string) error {
	m := &model.Group{
		GroupName:   groupName,
		GroupAvatar: groupAvatar,
		GourpDesc:   groupDesc,
	}
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	if _, err := dal.Group.WithContext(ctx).
		Where(
			dal.Group.GroupID.Eq(groupId),
			dal.Group.IsDeleted.Eq(constant.NotDeleted),
		).
		Updates(m); err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return nil
}

// GetGroupIds 根据用户id获取到自己加入的群组ids
func (r *RelationshipRepoImpl) GetGroupIds(ctx context.Context, userId int64) ([]string, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	groups, err := dal.GroupMember.WithContext(ctx).
		Where(
			dal.GroupMember.MemberID.Eq(userId),
			dal.GroupMember.IsDeleted.Eq(constant.NotDeleted),
		).
		Select(dal.GroupMember.GroupID).
		Find()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	var ids []string
	for _, group := range groups {
		ids = append(ids, group.GroupID)
	}
	return ids, nil
}

// GetGroups 根据群组ids获取群组信息
func (r *RelationshipRepoImpl) GetGroups(ctx context.Context, groupIds []string) ([]*universal.Group, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	groups, err := dal.Group.WithContext(ctx).
		Where(
			dal.Group.GroupID.In(groupIds...),
			dal.Group.IsDeleted.Eq(constant.NotDeleted)).
		Find()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	var res []*universal.Group
	for _, group := range groups {
		res = append(res, &universal.Group{
			GroupId:     group.GroupID,
			GroupName:   group.GroupName,
			GroupAvatar: group.GroupAvatar,
			GroupDesc:   group.GourpDesc,
			CreateAt:    group.CreateAt.String(),
		})
	}
	return res, nil
}

// CreateGroup 创建群聊
func (r *RelationshipRepoImpl) CreateGroup(ctx context.Context, leaderId int64, groupName string, groupAvatar string, groupDesc string) (*universal.Group, error) {
	m := &model.Group{
		GroupID:       r.groupUidGenerator.Generate(),
		GroupName:     groupName,
		GroupAvatar:   groupAvatar,
		GourpDesc:     groupDesc,
		CreateAt:      time.Now(),
		GroupLeaderID: leaderId,
	}
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	if err := dal.Group.WithContext(ctx).
		Create(m); err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	member := &model.GroupMember{
		MemberID: leaderId,
		GroupID:  m.GroupID,
		BecomeAt: time.Now(),
		Role:     constant.Leader,
	}
	if err := dal.GroupMember.WithContext(ctx).
		Create(member); err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return &universal.Group{
		GroupId:     m.GroupID,
		GroupName:   m.GroupName,
		GroupAvatar: m.GroupAvatar,
		GroupDesc:   m.GourpDesc,
		CreateAt:    m.CreateAt.String(),
	}, nil
}

func (r *RelationshipRepoImpl) GetFriend(ctx context.Context, uid, friendId int64) (*model.Friend, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	friend, err := dal.Friend.WithContext(ctx).
		Where(
			dal.Friend.UserID.Eq(uid),
			dal.Friend.FriendID.Eq(friendId),
			dal.Friend.IsDeleted.Eq(constant.NotDeleted),
		).First()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return friend, nil
}

func (r *RelationshipRepoImpl) GetFriendRequestsByRequestIds(ctx context.Context, requestId []int64) ([]*model.FriendRequest, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	requests, err := dal.FriendRequest.WithContext(ctx).
		Where(dal.FriendRequest.RequestID.In(requestId...)).
		Find()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return requests, nil
}

func (r *RelationshipRepoImpl) DeleteFriendGroup(ctx context.Context, uid int64, groupName string) error {
	err := dal.Q.Transaction(func(tx *dal.Query) error {
		if err := pkg.ContextErr(ctx); err != nil {
			return err
		}
		_, err := tx.FriendGroup.WithContext(ctx).
			Where(
				tx.FriendGroup.UserID.Eq(uid),
				tx.FriendGroup.GroupName.Eq(groupName),
			).
			Delete()
		if err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		//将该分组下的好友移动到默认分组
		_, err = tx.Friend.WithContext(ctx).
			Where(
				tx.Friend.UserID.Eq(uid),
				tx.Friend.GroupName.Eq(groupName),
			).
			Updates(&model.Friend{
				GroupName: constant.DefaultFriendGroup,
			})
		if err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		return nil
	})
	if err != nil {
		r.helper.Errorf(constant.TransactionErrorFormat, err)
		return pkg.InternalError(constant.TransactionErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) UpdateFriendGroup(ctx context.Context, uid int64, groupName, newGroupName string) error {
	err := dal.Q.Transaction(func(tx *dal.Query) error {
		if err := pkg.ContextErr(ctx); err != nil {
			return err
		}
		_, err := tx.FriendGroup.WithContext(ctx).
			Where(
				tx.FriendGroup.UserID.Eq(uid),
				tx.FriendGroup.GroupName.Eq(groupName),
			).
			Update(tx.FriendGroup.GroupName, newGroupName)
		if err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		r.helper.Infof("更新好友表中的分组名,uid:%d,groupName:%s,newGroupName:%s", uid, groupName, newGroupName)
		//更新好友表中的分组名
		_, err = tx.Friend.WithContext(ctx).
			Where(
				tx.Friend.UserID.Eq(uid),
				tx.Friend.GroupName.Eq(groupName),
			).
			Update(tx.FriendGroup.GroupName, newGroupName)
		if err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		r.helper.Info("更新好友表中的分组名成功")
		return nil
	})
	if err != nil {
		r.helper.Errorf(constant.TransactionErrorFormat, err)
		return pkg.InternalError(constant.TransactionErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) CreateFriendGroup(ctx context.Context, uid int64, groupName string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	err := dal.FriendGroup.WithContext(ctx).Create(&model.FriendGroup{
		UserID:    uid,
		GroupName: groupName,
	})
	if err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) UpdateFriendInfo(ctx context.Context, uid, friendId int64, noteName, groupName string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	_, err := dal.Friend.WithContext(ctx).
		Where(
			dal.Friend.UserID.Eq(uid),
			dal.Friend.FriendID.Eq(friendId),
		).
		Updates(&model.Friend{
			NoteName:  noteName,
			GroupName: groupName,
		})
	if err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) DelFriend(ctx context.Context, uid, friendId int64) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	_, err := dal.Friend.WithContext(ctx).
		Where(
			dal.Friend.UserID.Eq(uid),
			dal.Friend.FriendID.Eq(friendId),
		).
		Update(dal.Friend.IsDeleted, constant.Deleted)
	if err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) GetFriends(ctx context.Context, uid int64) ([]*model.Friend, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	friends, err := dal.Friend.WithContext(ctx).
		Where(
			dal.Friend.UserID.Eq(uid),
			dal.Friend.IsDeleted.Eq(constant.NotDeleted),
		).
		Find()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return friends, nil
}

func (r *RelationshipRepoImpl) GetFriendGroups(ctx context.Context, uid int64) ([]*model.FriendGroup, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}

	friendGroups, err := dal.FriendGroup.WithContext(ctx).
		Where(dal.FriendGroup.UserID.Eq(uid)).
		Find()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return friendGroups, nil
}

func (r *RelationshipRepoImpl) GetFriendsByGroup(ctx context.Context, uid int64, groupName string) ([]*model.Friend, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	friends, err := dal.Friend.WithContext(ctx).
		Where(
			dal.Friend.UserID.Eq(uid),
			dal.Friend.GroupName.Eq(groupName),
		).Find()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return friends, nil
}

func (r *RelationshipRepoImpl) CreateFriend(ctx context.Context, uid, friendId int64, noteName, groupName string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	m := model.Friend{
		UserID:    uid,
		FriendID:  friendId,
		NoteName:  noteName,
		GroupName: groupName,
		BecomeAt:  time.Now(),
	}
	err := dal.Friend.WithContext(ctx).
		Create(&m)
	if err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) DealFriendRequest(ctx context.Context, requestId int64, status, noteName, groupName string) (*model.FriendRequest, error) {
	var request *model.FriendRequest
	var err error
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		if err = pkg.ContextErr(ctx); err != nil {
			return err
		}
		// update request status
		_, err = tx.WithContext(ctx).FriendRequest.
			Where(tx.FriendRequest.RequestID.Eq(requestId)).
			Update(tx.FriendRequest.Status, status)
		if err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		// if status is refused, return
		if status == constant.Refused {
			return nil
		}
		// get request
		request, err = tx.WithContext(ctx).FriendRequest.
			Where(tx.FriendRequest.RequestID.Eq(requestId)).
			First()
		if err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		var friends []*model.Friend
		// create friend
		friends = append(friends, &model.Friend{
			UserID:    request.RequesterID,
			FriendID:  request.ReceiverID,
			NoteName:  request.NoteName,
			GroupName: request.GroupName,
			BecomeAt:  time.Now(),
		})
		friends = append(friends, &model.Friend{
			UserID:    request.ReceiverID,
			FriendID:  request.RequesterID,
			NoteName:  noteName,
			GroupName: groupName,
			BecomeAt:  time.Now(),
		})

		if err = pkg.ContextErr(ctx); err != nil {
			return err
		}
		err = tx.WithContext(ctx).Friend.
			CreateInBatches(friends, 10)
		if err != nil {
			r.helper.Errorf(constant.SqlErrorFormat, err)
			return pkg.InternalError(constant.SqlErrorFormat, err)
		}
		return nil
	})
	if err != nil {
		err = pkg.InternalError(constant.TransactionErrorFormat, err)
		r.helper.Errorf(err.Error())
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return request, nil
}

func (r *RelationshipRepoImpl) UpdateFriendRequestStatus(ctx context.Context, requestId int64, status string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	_, err := dal.FriendRequest.WithContext(ctx).
		Where(dal.FriendRequest.RequestID.Eq(requestId)).
		Update(dal.FriendRequest.Status, status)
	if err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return nil
}

func (r *RelationshipRepoImpl) GetFriendRequestByRequestId(ctx context.Context, requestId int64) (*model.FriendRequest, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	result, err := dal.FriendRequest.WithContext(ctx).
		Where(dal.FriendRequest.RequestID.Eq(requestId)).
		First()
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return result, nil
}

func (r *RelationshipRepoImpl) GetFriendRequestByPage(ctx context.Context, uid int64, offset, size int) ([]*model.FriendRequest, int, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, 0, err
	}
	result, count, err := dal.FriendRequest.WithContext(ctx).
		Where(dal.FriendRequest.RequesterID.Eq(uid)).
		Or(dal.FriendRequest.ReceiverID.Eq(uid)).
		FindByPage(offset, size)
	if pkg.IsNotRecordNotFoundError(err) != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, 0, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	if count == 0 {
		return nil, 0, nil
	}
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, 0, err
	}
	count, err = dal.FriendRequest.WithContext(ctx).
		Where(dal.FriendRequest.RequesterID.Eq(uid)).
		Or(dal.FriendRequest.ReceiverID.Eq(uid)).
		Count()
	r.helper.Infof("count: %d", count)
	if err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, 0, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return result, int(count), nil
}

func (r *RelationshipRepoImpl) CreateFriendRequest(ctx context.Context, uid, friendId int64, noteName, groupName, desc string) (*model.FriendRequest, error) {
	//check friend
	count, err := dal.Friend.WithContext(ctx).
		Where(
			dal.Friend.UserID.Eq(uid),
			dal.Friend.FriendID.Eq(friendId),
		).
		Count()
	if err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	if count > 0 {
		return nil, pkg.InvalidArgumentError("already friend")
	}
	//check request
	count, err = dal.FriendRequest.WithContext(ctx).
		Where(
			dal.FriendRequest.RequesterID.Eq(uid),
			dal.FriendRequest.ReceiverID.Eq(friendId),
			dal.FriendRequest.Status.Eq(constant.Pending)).
		Or(
			dal.FriendRequest.RequesterID.Eq(friendId),
			dal.FriendRequest.ReceiverID.Eq(uid),
			dal.FriendRequest.Status.Eq(constant.Pending)).
		Count()
	if err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	if count > 0 {
		return nil, pkg.InvalidArgumentError("already request")
	}
	m := model.FriendRequest{
		RequestID:   r.friendUidGenerator.Generate().Int64(),
		RequesterID: uid,
		ReceiverID:  friendId,
		NoteName:    noteName,
		GroupName:   groupName,
		Desc:        desc,
		Status:      constant.Pending, // 0:未处理 1:已同意 2:已拒绝
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	err = dal.FriendRequest.WithContext(ctx).
		Create(&m)
	if err != nil {
		r.helper.Errorf(constant.SqlErrorFormat, err)
		return nil, pkg.InternalError(constant.SqlErrorFormat, err)
	}
	return &m, nil
}

func NewRelationshipRepoImpl(helper *log.Helper, friendUidGenerator *uid.FriendRequestUidGenerator, groupUidGenerator *uid.GroupRequestUidGenerator) biz.RelationshipRepo {
	return &RelationshipRepoImpl{
		helper:             helper,
		friendUidGenerator: friendUidGenerator,
		groupUidGenerator:  groupUidGenerator,
	}
}
