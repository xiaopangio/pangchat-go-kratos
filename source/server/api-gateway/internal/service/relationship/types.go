package service_relationship

import (
	"api-gateway/api/v1/logic/relationship"
	"api-gateway/api/v1/universal"
)

// SendFriendRequestRequest 发送好友请求请求
type SendFriendRequestRequest struct {
	RequesterId string `json:"requester_id"`
	ReceiverId  string `json:"receiver_id"`
	NoteName    string `json:"note_name"`
	Desc        string `json:"desc"`
	GroupName   string `json:"group_name"`
}

// SendFriendRequestResponse 发送好友请求响应
type SendFriendRequestResponse struct {
	FriendRequest *FriendRequest `json:"request"`
}
type FriendRequest struct {
	RequestId   string `json:"request_id"`
	RequesterId string `json:"requester_id"`
	ReceiverId  string `json:"receiver_id"`
	Desc        string `json:"desc"`
	Status      string `json:"status"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	NickName    string `json:"nick_name"`
	Avatar      string `json:"avatar"`
}

// GetFriendRequestListRequest 获取好友请求列表请求
type GetFriendRequestListRequest struct {
	UserId     string `json:"user_id"`
	PageNumber int    `json:"page_number"`
	PageSize   int    `json:"page_size"`
}

// GetFriendRequestListResponse 获取好友请求列表响应
type GetFriendRequestListResponse struct {
	Total int64                      `json:"total"`
	List  []*universal.FriendRequest `json:"list"`
}

// GetFriendRequestRequest 获取好友请求请求
type GetFriendRequestRequest struct {
	RequestId string `json:"request_id"`
}

// GetFriendRequestResponse 获取好友请求响应
type GetFriendRequestResponse struct {
	FriendRequest *universal.FriendRequest `json:"friend_request"`
}

// DealFriendRequestRequest 处理好友请求请求
type DealFriendRequestRequest struct {
	RequestId string `json:"request_id"`
	Status    int    `json:"status"`
}

// DealFriendRequestResponse 处理好友请求响应
type DealFriendRequestResponse struct {
}

// GetFriendListRequest 获取好友列表请求
type GetFriendListRequest struct {
	UserId string `json:"user_id"`
}

// GetFriendListResponse 获取好友列表响应
type GetFriendListResponse struct {
	*relationship.GetFriendListResponse `json:"friend_list"`
}

// DeleteFriendRequest 删除好友请求
type DeleteFriendRequest struct {
	UserId   string `json:"user_id"`
	FriendId string `json:"friend_id"`
}

// DeleteFriendResponse 删除好友响应
type DeleteFriendResponse struct {
}

// GetFriendInfoRequest 获取好友信息请求
type GetFriendInfoRequest struct {
	FriendId string `json:"friend_id"`
}

// GetFriendInfoResponse 获取好友信息响应
type GetFriendInfoResponse struct {
	CityName     string `json:"city_name"`
	ProvinceName string `json:"province_name"`
	Desc         string `json:"desc"`
}

// UpdateFriendInfoRequest 更新好友信息请求
type UpdateFriendInfoRequest struct {
	UserId    string `json:"user_id"`
	FriendId  string `json:"friend_id"`
	NoteName  string `json:"note_name"`
	GroupName string `json:"group_name"`
}

// UpdateFriendInfoResponse 更新好友信息响应
type UpdateFriendInfoResponse struct {
}

// CreateFriendGroupRequest 创建好友分组请求
type CreateFriendGroupRequest struct {
	UserId    string `json:"user_id"`
	GroupName string `json:"group_name"`
}

// CreateFriendGroupResponse 创建好友分组响应
type CreateFriendGroupResponse struct {
}

// UpdateFriendGroupRequest 更新好友分组请求
type UpdateFriendGroupRequest struct {
	UserId       string `json:"user_id"`
	GroupName    string `json:"group_name"`
	NewGroupName string `json:"new_group_name"`
}

// UpdateFriendGroupResponse 更新好友分组响应
type UpdateFriendGroupResponse struct {
}

// DeleteFriendGroupRequest 删除好友分组请求
type DeleteFriendGroupRequest struct {
	UserId    string `json:"user_id"`
	GroupName string `json:"group_name"`
}

// DeleteFriendGroupResponse 删除好友分组响应
type DeleteFriendGroupResponse struct {
}

// GetFriendGroupListRequest 获取好友分组列表请求
type GetFriendGroupListRequest struct {
	UserId string `json:"user_id" form:"user_id" binding:"required"`
}

// GetFriendGroupListResponse 获取好友分组列表响应
type GetFriendGroupListResponse struct {
	GroupNames []string `json:"group_names"`
}
