package data

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"relationship/internal/biz"
	"relationship/internal/data/orm/dal"
	"relationship/internal/data/orm/model"
	"relationship/pkg"
	"time"
)

type RelationshipRepoImpl struct {
	helper *log.Helper
	uidGen *snowflake.Node
}

func (r *RelationshipRepoImpl) FindFriendRequests(ctx context.Context, requestId []int64) ([]*model.FriendRequest, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	requests, err := dal.FriendRequest.WithContext(ctx).Where(dal.FriendRequest.RequestID.In(requestId...)).Find()
	if err != nil {
		return nil, pkg.InternalError("find friend requests error: %v", err)
	}
	return requests, nil
}

func (r *RelationshipRepoImpl) DeleteFriendGroup(ctx context.Context, uid int64, groupName string) error {

	err := dal.Q.Transaction(func(tx *dal.Query) error {
		if err := pkg.ContextErr(ctx); err != nil {
			return err
		}
		_, err := tx.FriendGroup.WithContext(ctx).Where(tx.FriendGroup.UserID.Eq(uid), tx.FriendGroup.GroupName.Eq(groupName)).Delete()
		if err != nil {
			return pkg.InternalError("delete friend group error: %v", err)
		}
		//将该分组下的好友移动到默认分组
		_, err = tx.Friend.WithContext(ctx).Where(tx.Friend.UserID.Eq(uid), tx.Friend.GroupName.Eq(groupName)).Updates(&model.Friend{
			GroupName: pkg.DefaultFriendGroup,
		})
		if err != nil {
			return pkg.InternalError("move friend error: %v", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RelationshipRepoImpl) UpdateFriendGroup(ctx context.Context, uid int64, groupName, newGroupName string) error {
	err := dal.Q.Transaction(func(tx *dal.Query) error {
		if err := pkg.ContextErr(ctx); err != nil {
			return err
		}
		_, err := tx.FriendGroup.WithContext(ctx).Where(tx.FriendGroup.UserID.Eq(uid), tx.FriendGroup.GroupName.Eq(groupName)).Updates(&model.FriendGroup{
			GroupName: newGroupName,
		})
		if err != nil {
			return pkg.InternalError("update friend group error: %v", err)
		}
		//更新好友表中的分组名
		_, err = tx.Friend.WithContext(ctx).Where(tx.Friend.UserID.Eq(uid), tx.Friend.GroupName.Eq(groupName)).Updates(&model.Friend{
			GroupName: newGroupName,
		})
		if err != nil {
			return pkg.InternalError("update friend group error: %v", err)
		}
		return nil
	})
	if err != nil {
		return err
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
		return pkg.InternalError("create friend group error: %v", err)
	}
	return nil
}

func (r *RelationshipRepoImpl) UpdateFriendInfo(ctx context.Context, uid, friendId int64, noteName, groupName string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	_, err := dal.Friend.WithContext(ctx).Where(dal.Friend.UserID.Eq(uid), dal.Friend.FriendID.Eq(friendId)).Updates(&model.Friend{
		NoteName:  noteName,
		GroupName: groupName,
	})
	if err != nil {
		return pkg.InternalError("update friend info error: %v", err)
	}
	return nil
}

func (r *RelationshipRepoImpl) DelFriend(ctx context.Context, uid, friendId int64) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	_, err := dal.Friend.WithContext(ctx).Where(dal.Friend.UserID.Eq(uid), dal.Friend.FriendID.Eq(friendId)).Delete()
	if err != nil {
		return pkg.InternalError("delete friend error: %v", err)
	}
	return nil
}

func (r *RelationshipRepoImpl) GetFriends(ctx context.Context, uid int64) ([]*model.Friend, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	friends, err := dal.Friend.WithContext(ctx).Where(dal.Friend.UserID.Eq(uid)).Find()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		return nil, err
	}
	return friends, nil
}

func (r *RelationshipRepoImpl) GetFriendGroups(ctx context.Context, uid int64) ([]*model.FriendGroup, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}

	friendGroups, err := dal.FriendGroup.WithContext(ctx).Where(dal.FriendGroup.UserID.Eq(uid)).Find()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		return nil, err
	}
	return friendGroups, nil
}

func (r *RelationshipRepoImpl) GetFriendsByGroup(ctx context.Context, uid int64, groupName string) ([]*model.Friend, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	friends, err := dal.Friend.WithContext(ctx).Where(dal.Friend.UserID.Eq(uid), dal.Friend.GroupName.Eq(groupName)).Find()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		return nil, err
	}
	return friends, nil
}

func (r *RelationshipRepoImpl) CreateFriend(ctx context.Context, uid, friendId int64, noteName, groupName string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	m := model.Friend{
		UserID:       uid,
		FriendID:     friendId,
		NoteName:     noteName,
		GroupName:    groupName,
		LastAckMsgID: "",
		BecomeAt:     time.Now(),
	}
	err := dal.Friend.WithContext(ctx).Create(&m)
	if err != nil {
		return pkg.InternalError("create friend error: %v", err)
	}
	return nil
}

func (r *RelationshipRepoImpl) DealFriendRequest(ctx context.Context, requestId int64, status string) (int64, error) {
	var userId int64
	err := dal.Q.Transaction(func(tx *dal.Query) error {
		// check status is pending or not
		if status == pkg.Pending {
			return pkg.InvalidArgumentError("pending status is not allowed")
		}
		if err := pkg.ContextErr(ctx); err != nil {
			return err
		}
		//check request status is pending or not
		res, err := tx.WithContext(ctx).Friend.Where(tx.FriendRequest.RequestID.Eq(requestId)).First()
		if err = pkg.IsNotRecordNotFoundError(err); err != nil {
			return err
		}
		if status == pkg.Agreed || status == pkg.Refused {
			return pkg.InvalidArgumentError("request is already dealt")
		}
		if err := pkg.ContextErr(ctx); err != nil {
			return err
		}
		// update request status
		_, err = tx.WithContext(ctx).Friend.Where(tx.FriendRequest.RequestID.Eq(requestId)).Update(tx.FriendRequest.Status, status)
		if err != nil {
			return pkg.InternalError("update friend request status error: %v", err)
		}
		// if status is refused, return
		if status == pkg.Refused {
			return nil
		}
		// create friend
		m := model.Friend{
			UserID:       res.UserID,
			FriendID:     res.FriendID,
			NoteName:     res.NoteName,
			GroupName:    res.GroupName,
			LastAckMsgID: "",
			BecomeAt:     time.Now(),
		}
		if err := pkg.ContextErr(ctx); err != nil {
			return err
		}
		err = tx.WithContext(ctx).Friend.Create(&m)
		if err != nil {
			return pkg.InternalError("create friend error: %v", err)
		}
		userId = m.UserID
		return nil
	})
	return userId, err
}

func (r *RelationshipRepoImpl) UpdateFriendRequestStatus(ctx context.Context, requestId int64, status string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	_, err := dal.FriendRequest.WithContext(ctx).Where(dal.FriendRequest.RequestID.Eq(requestId)).Update(dal.FriendRequest.Status, status)
	if err != nil {
		return pkg.InternalError("update friend request status error: %v", err)
	}
	return nil
}

func (r *RelationshipRepoImpl) FindFriendRequest(ctx context.Context, requestId int64) (*model.FriendRequest, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	result, err := dal.FriendRequest.WithContext(ctx).Where(dal.FriendRequest.RequestID.Eq(requestId)).First()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *RelationshipRepoImpl) FindAllFriendRequestByPage(ctx context.Context, uid int64, offset, size int) ([]*model.FriendRequest, int, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, 0, err
	}
	result, count, err := dal.FriendRequest.WithContext(ctx).Where(dal.FriendRequest.RequesterID.Eq(uid)).Or(dal.FriendRequest.ReceiverID.Eq(uid)).FindByPage(offset, size)
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return nil, 0, pkg.NotFoundError("not found friend request")
	}
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, 0, err
	}
	count, err = dal.FriendRequest.WithContext(ctx).Where(dal.FriendRequest.RequesterID.Eq(uid)).Or(dal.FriendRequest.ReceiverID.Eq(uid)).Count()
	r.helper.Infof("count: %d", count)
	if err != nil {
		return nil, 0, pkg.InternalError("count friend request error: %v", err)
	}
	return result, int(count), nil
}

func (r *RelationshipRepoImpl) CreateFriendRequest(ctx context.Context, uid, friendId int64, noteName, groupName, desc string) (*model.FriendRequest, error) {
	//check friend
	count, err := dal.Friend.WithContext(ctx).Where(dal.Friend.UserID.Eq(uid), dal.Friend.FriendID.Eq(friendId)).Count()
	if err != nil {
		return nil, pkg.InternalError("count friend error: %v", err)
	}
	if count > 0 {
		return nil, pkg.InvalidArgumentError("already friend")
	}
	//check request
	count, err = dal.FriendRequest.WithContext(ctx).
		Where(dal.FriendRequest.RequesterID.Eq(uid), dal.FriendRequest.ReceiverID.Eq(friendId), dal.FriendRequest.Status.Eq("0")).
		Or(dal.FriendRequest.RequesterID.Eq(friendId), dal.FriendRequest.ReceiverID.Eq(uid), dal.FriendRequest.Status.Eq("0")).
		Count()
	if err != nil {
		return nil, pkg.InternalError("count friend request error: %v", err)
	}
	if count > 0 {
		return nil, pkg.InvalidArgumentError("already request")
	}
	m := model.FriendRequest{
		RequestID:   r.uidGen.Generate().Int64(),
		RequesterID: uid,
		ReceiverID:  friendId,
		NoteName:    noteName,
		GroupName:   groupName,
		Desc:        desc,
		Status:      "0", // 0:未处理 1:已同意 2:已拒绝
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	err = dal.FriendRequest.WithContext(ctx).Create(&m)
	if err != nil {
		return nil, pkg.InvalidArgumentError("create friend request error: %v", err)
	}
	return &m, nil
}

func NewRelationshipRepoImpl(helper *log.Helper, uidGen *snowflake.Node) biz.RelationshipRepo {
	return &RelationshipRepoImpl{helper: helper, uidGen: uidGen}
}
