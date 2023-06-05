package service_relationship

import (
	"api-gateway/api/v1/logic/relationship"
	"api-gateway/internal/components/auth"
	"api-gateway/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
)

type Relationship struct {
	helper *log.Helper
	client relationship.RelationShipClient
	Jwt    *auth.JwtManager
}

func NewRelationship(helper *log.Helper, client relationship.RelationShipClient, Jwt *auth.JwtManager) *Relationship {
	return &Relationship{helper: helper, client: client, Jwt: Jwt}
}
func (r *Relationship) SendFriendRequest(ctx *gin.Context) {
	var req SendFriendRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	reply, err := r.client.SendFriendRequest(ctx, &relationship.SendFriendRequestRequest{
		RequesterId: pkg.ParseInt64(req.RequesterId),
		ReceiverId:  pkg.ParseInt64(req.ReceiverId),
		NoteName:    req.NoteName,
		Desc:        req.Desc,
		GroupName:   req.GroupName,
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	res := &SendFriendRequestResponse{
		FriendRequest: &FriendRequest{
			RequestId:   pkg.FormatInt(reply.FriendRequest.RequestId),
			RequesterId: pkg.FormatInt(reply.FriendRequest.RequesterId),
			ReceiverId:  pkg.FormatInt(reply.FriendRequest.ReceiverId),
			Desc:        reply.FriendRequest.Desc,
			Status:      reply.FriendRequest.Status,
			CreateTime:  reply.FriendRequest.CreateTime,
			UpdateTime:  reply.FriendRequest.UpdateTime,
			NickName:    reply.FriendRequest.NickName,
			Avatar:      reply.FriendRequest.Avatar,
		},
	}
	pkg.Ok(ctx, res)
}
func (r *Relationship) GetFriendRequestList(ctx *gin.Context) {
	var req GetFriendRequestListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	reply, err := r.client.GetFriendRequestList(ctx, &relationship.GetFriendRequestListRequest{
		UserId:     pkg.ParseInt64(req.UserId),
		PageNumber: int64(req.PageNumber),
		PageSize:   int64(req.PageSize),
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	resp := &GetFriendRequestListResponse{
		Total: reply.Total,
		List:  reply.FriendRequests,
	}
	pkg.Ok(ctx, resp)
}
func (r *Relationship) GetFriendRequest(ctx *gin.Context) {
	var req GetFriendRequestRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	reply, err := r.client.GetFriendRequest(ctx, &relationship.GetFriendRequestRequest{
		RequestId: pkg.ParseInt64(req.RequestId),
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	resp := &GetFriendRequestResponse{
		FriendRequest: reply.FriendRequest,
	}
	pkg.Ok(ctx, resp)
}
func (r *Relationship) DealFriendRequest(ctx *gin.Context) {
	var req DealFriendRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	_, err := r.client.DealFriendRequest(ctx, &relationship.DealFriendRequestRequest{
		RequestId: pkg.ParseInt64(req.RequestId),
		Status: strconv.Itoa(
			req.Status),
		NoteName:  req.NoteName,
		GroupName: req.GroupName,
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, nil)
}
func (r *Relationship) GetFriendList(ctx *gin.Context) {
	var req GetFriendListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	reply, err := r.client.GetFriendList(ctx, &relationship.GetFriendListRequest{
		UserId: pkg.ParseInt64(req.UserId),
	})
	r.helper.Info("GetFriendList", "reply", reply)
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, reply)
}
func (r *Relationship) DeleteFriend(ctx *gin.Context) {
	var req DeleteFriendRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	_, err := r.client.DeleteFriend(ctx, &relationship.DeleteFriendRequest{
		UserId:   pkg.ParseInt64(req.UserId),
		FriendId: pkg.ParseInt64(req.FriendId),
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, nil)
}
func (r *Relationship) GetFriendInfo(ctx *gin.Context) {
	var req GetFriendInfoRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	reply, err := r.client.GetFriendInfo(ctx, &relationship.GetFriendInfoRequest{
		FriendId: pkg.ParseInt64(req.FriendId),
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	resp := &GetFriendInfoResponse{
		CityName:     reply.CityName,
		ProvinceName: reply.ProvinceName,
		Desc:         reply.Desc,
		AccountId:    reply.AccountId,
	}
	pkg.Ok(ctx, resp)
}
func (r *Relationship) UpdateFriendInfo(ctx *gin.Context) {
	var req UpdateFriendInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	_, err := r.client.UpdateFriendInfo(ctx, &relationship.UpdateFriendInfoRequest{
		UserId:    pkg.ParseInt64(req.UserId),
		FriendId:  pkg.ParseInt64(req.FriendId),
		NoteName:  req.NoteName,
		GroupName: req.GroupName,
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, nil)

}
func (r *Relationship) CreateFriendGroup(ctx *gin.Context) {
	var req CreateFriendGroupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	_, err := r.client.CreateFriendGroup(ctx, &relationship.CreateFriendGroupRequest{
		UserId:    pkg.ParseInt64(req.UserId),
		GroupName: req.GroupName,
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, nil)
}
func (r *Relationship) UpdateFriendGroup(ctx *gin.Context) {
	var req UpdateFriendGroupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	_, err := r.client.UpdateFriendGroup(ctx, &relationship.UpdateFriendGroupRequest{
		UserId:       pkg.ParseInt64(req.UserId),
		GroupName:    req.GroupName,
		NewGroupName: req.NewGroupName,
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, nil)
}
func (r *Relationship) DeleteFriendGroup(ctx *gin.Context) {
	var req DeleteFriendGroupRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	_, err := r.client.DeleteFriendGroup(ctx, &relationship.DeleteFriendGroupRequest{
		UserId:    pkg.ParseInt64(req.UserId),
		GroupName: req.GroupName,
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	pkg.Ok(ctx, nil)
}
func (r *Relationship) GetFriendGroupList(ctx *gin.Context) {
	var req GetFriendGroupListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		pkg.Validator(ctx, err)
		return
	}
	r.helper.Info(req.UserId)
	reply, err := r.client.GetFriendGroupList(ctx, &relationship.GetFriendGroupListRequest{
		UserId: pkg.ParseInt64(req.UserId),
	})
	if err = pkg.HandlerError(ctx, err); err != nil {
		return
	}
	resp := &GetFriendGroupListResponse{
		GroupNames: reply.GroupNames,
	}
	pkg.Ok(ctx, resp)
}
