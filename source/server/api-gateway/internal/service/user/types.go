package service_user

import "api-gateway/api/v1/logic/user"

// RegisterRequest 注册请求
type RegisterRequest struct {
	Type            int    `json:"type" binding:"required" label:"账户类型"`
	Username        string `json:"username" binding:"required,min=8,max=25" label:"账户"`
	Password        string `json:"password" binding:"required,len=40" label:"密码"`
	PasswordConfirm string `json:"password_confirm" binding:"required,len=40" label:"密码确认"`
	NickName        string `json:"nick_name" binding:"required,max=20" label:"昵称"`
	AvatarUrl       string `json:"avatar_url" binding:"required,max=255" label:"头像"`
}

// RegisterResponse 注册响应
type RegisterResponse struct {
}

// SmsCodeRequest 短信验证码请求
type SmsCodeRequest struct {
	Phone string `json:"phone" binding:"required,len=11" label:"手机号"`
}

// SmsCodeResponse 短信验证码响应
type SmsCodeResponse struct {
}

// ModifyPasswdRequest 修改密码请求
type ModifyPasswdRequest struct {
	Uid                string `json:"uid" binding:"required" label:"UID"`
	OldPassword        string `json:"old_password" binding:"required,len=40" label:"原密码"`
	NewPassword        string `json:"new_password" binding:"required,len=40" label:"新密码"`
	NewPasswordConfirm string `json:"new_password_confirm" binding:"required,len=40" label:"确认新密码"`
}

// ModifyPasswdResponse 修改密码响应
type ModifyPasswdResponse struct {
}

// ResetPasswordRequest 重置密码请求
type ResetPasswordRequest struct {
	Uid             string `json:"uid" binding:"required" label:"UID"`
	Phone           string `json:"phone" binding:"required,len=11" label:"手机号"`
	Password        string `json:"password" binding:"required,len=40" label:"新密码"`
	PasswordConfirm string `json:"password_confirm" binding:"required,len=40" label:"确认新密码"`
}

// ResetPasswordResponse 重置密码响应
type ResetPasswordResponse struct {
}

// ModifyAccountIDRequest 修改账户ID请求
type ModifyAccountIDRequest struct {
	Uid       string `json:"uid" binding:"required" label:"UID"`
	AccountId string `json:"account_id" binding:"required" label:"accountId"`
}

// ModifyAccountIDResponse 修改账户ID响应
type ModifyAccountIDResponse struct {
	AccountId string `json:"account_id"`
}

// ModifyProfileRequest 修改个人资料请求
type ModifyProfileRequest struct {
	UserId       string `json:"user_id" binding:"required" label:"UID"`
	NickName     string `json:"nick_name" binding:"max=10" label:"昵称"`
	PersonalDesc string `json:"personal_desc" label:"个性签名"`
	CityId       string `json:"city_id"  label:"城市编号"`
}

// ModifyProfileResponse 修改个人资料响应
type ModifyProfileResponse struct {
	Token string `json:"token"`
}

// ProfileRequest 个人资料请求
type ProfileRequest struct {
}

// ProfileResponse 个人资料响应
type ProfileResponse struct {
	UserId    string `json:"user_id"`
	AccountId string `json:"account_id"`
	NickName  string `json:"nick_name"`
	Desc      string `json:"desc"`
	City      string `json:"city"`
	Province  string `json:"province"`
	AvatarUrl string `json:"avatar_url"`
}

// AddressListRequest 地址列表请求
type AddressListRequest struct {
}

// AddressListResponse 地址列表响应
type AddressListResponse struct {
	Provinces []*user.Province `json:"provinces"`
}

// BindPhoneRequest 绑定手机号请求
type BindPhoneRequest struct {
	Phone string `json:"phone" binding:"required" label:"手机号"`
}

// VerifyCodeRequest 验证验证码请求
type VerifyCodeRequest struct {
	Phone   string `json:"phone" binding:"required,len=11" label:"手机号"`
	SmsCode string `json:"code" binding:"required,len=4" label:"验证码"`
}

// VerifyCodeResponse 验证验证码响应
type VerifyCodeResponse struct {
}

// GetAvatarRequest 获取头像请求
type GetAvatarRequest struct {
	AvatarUrl string `json:"avatar_url" binding:"required" label:"头像地址"`
}

// GetAvatarResponse 获取头像响应
type GetAvatarResponse struct {
}

// UploadAvatarRequest 上传头像请求
type UploadAvatarRequest struct {
}

// UploadAvatarResponse 上传头像响应
type UploadAvatarResponse struct {
	Token string `json:"token"`
}
