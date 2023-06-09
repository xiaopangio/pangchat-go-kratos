package pkg

import "errors"

// 定义错误
var (
	ErrNotGroupMember       = errors.New("不是群成员")
	ErrNotFriend            = errors.New("不是好友")
	ErrAlreadyFriend        = errors.New("已经是好友")
	ErrAlreadyGroup         = errors.New("已经是群成员")
	ErrAlreadyFriendRequest = errors.New("已经发送好友请求")
)
