package service_user

import (
	"api-gateway/api/v1/logic/user"
	"api-gateway/internal/components/auth"
	"api-gateway/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"mime/multipart"
	"strings"
)

type UserService struct {
	client user.UserClient
	helper *log.Helper
	Jwt    *auth.JwtManager
}

func NewUserService(client user.UserClient, helper *log.Helper, jwt *auth.JwtManager) *UserService {
	return &UserService{client: client, helper: helper, Jwt: jwt}
}

func (u *UserService) GetUserFromToken(c *gin.Context) (*auth.UserToken, bool) {
	token, exists := c.Get("token")
	if !exists {
		pkg.Forbidden(c)
		u.helper.Errorf("user token not found")
		return &auth.UserToken{}, false
	}
	userToken := token.(*auth.UserToken)
	return userToken, true
}

// Ping ping
func (u *UserService) Ping(c *gin.Context) {
	_, err := u.client.Ping(c.Request.Context(), &user.PingRequest{})
	if err != nil {
		u.helper.Errorf("ping err: %v", err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Register 注册
func (u *UserService) Register(c *gin.Context) {
	var req RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		u.helper.Errorf("bind err: %v", err)
		pkg.Validator(c, err)
		return
	}
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	userReq := &user.RegisterUserRequest{
		Type:            int64(req.Type),
		Username:        req.Username,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		NickName:        req.NickName,
		AvatarUrl:       req.AvatarUrl,
	}
	reply, err := u.client.RegisterUser(ctx, userReq)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	pkg.Ok(c, reply)
	return
}

// SendSmsCode  发送短信验证码
func (u *UserService) SendSmsCode(c *gin.Context) {
	var req SmsCodeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		u.helper.Errorf("bind err: %v", err)

		pkg.Validator(c, err)
		return
	}
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	smsReq := &user.SendSmsCodeRequest{
		Phone: req.Phone,
	}
	reply, err := u.client.SendSmsCode(ctx, smsReq)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	pkg.Ok(c, reply)
	return
}

// VerifyCode 验证短信验证码
func (u *UserService) VerifyCode(c *gin.Context) {
	var req VerifyCodeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		u.helper.Errorf("bind err: %v", err)
		pkg.Validator(c, err)
		return
	}
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	verifyReq := &user.VerifyCodeRequest{
		Phone:   req.Phone,
		SmsCode: req.SmsCode,
	}
	reply, err := u.client.VerifyCode(ctx, verifyReq)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	pkg.Ok(c, reply)
	return
}

// ModifyPasswd 修改密码
func (u *UserService) ModifyPasswd(c *gin.Context) {
	var req ModifyPasswdRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		u.helper.Errorf("bind err: %v", err)
		pkg.Validator(c, err)
		return
	}
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	modifyReq := &user.ModifyPasswdRequest{
		Uid:                pkg.ParseInt64(req.Uid),
		OldPassword:        req.OldPassword,
		NewPassword:        req.NewPassword,
		NewPasswordConfirm: req.NewPasswordConfirm,
	}
	reply, err := u.client.ModifyPasswd(ctx, modifyReq)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	pkg.Ok(c, reply)
	return
}

// ResetPassword 重置密码
func (u *UserService) ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		u.helper.Errorf("bind err: %v", err)
		pkg.Validator(c, err)
		return
	}
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	resetReq := &user.ResetPasswordRequest{
		Uid:             pkg.ParseInt64(req.Uid),
		Phone:           req.Phone,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	}
	reply, err := u.client.ResetPassword(ctx, resetReq)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	pkg.Ok(c, reply)
	return
}

// ModifyAccountID 修改账户ID
func (u *UserService) ModifyAccountID(c *gin.Context) {
	var req ModifyAccountIDRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		u.helper.Errorf("bind err: %v", err)
		pkg.Validator(c, err)
		return
	}
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	modifyReq := &user.ModifyAccountIDRequest{
		Uid:       pkg.ParseInt64(req.Uid),
		AccountId: req.AccountId,
	}
	reply, err := u.client.ModifyAccountID(ctx, modifyReq)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	pkg.Ok(c, reply)
	return
}

// ModifyProfile 修改个人资料
func (u *UserService) ModifyProfile(c *gin.Context) {
	var req ModifyProfileRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		u.helper.Errorf("bind err: %v", err)
		pkg.Validator(c, err)
		return
	}
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	modifyReq := &user.ModifyProfileRequest{
		User: &user.UserProfile{
			UserId:       pkg.ParseInt64(req.UserId),
			NickName:     req.NickName,
			PersonalDesc: req.PersonalDesc,
			UserAddress: &user.UserAddress{
				CityId: req.CityId,
			},
		},
	}
	reply, err := u.client.ModifyProfile(ctx, modifyReq)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	token := &auth.UserToken{
		Uid:          pkg.FormatInt(reply.User.UserId),
		NickName:     reply.User.NickName,
		AccountId:    reply.User.AccountId,
		PersonalDesc: reply.User.PersonalDesc,
		Avatar:       reply.User.Avatar,
	}
	tokenString, err := u.Jwt.SignUser(token)
	if err != nil {
		u.helper.Errorf("sign user token err: %v", err)
		pkg.Fail(c, err)
		return
	}
	res := &ModifyProfileResponse{
		Token: tokenString,
	}
	pkg.Ok(c, res)
	return
}

// Profile 个人资料
func (u *UserService) Profile(c *gin.Context) {
	accountId := c.Query("account_id")
	if accountId == "" {
		pkg.FailMessage(c, "account_id is empty")
		return
	}
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	profileReq := &user.ProfileRequest{
		AccountId: accountId,
	}
	u.helper.Info(u.client)
	reply, err := u.client.Profile(ctx, profileReq)
	u.helper.Infof("reply: %v", reply)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	res := &ProfileResponse{
		UserId:    pkg.FormatInt(reply.User.UserId),
		AccountId: reply.User.AccountId,
		NickName:  reply.User.NickName,
		Desc:      reply.User.PersonalDesc,
		City:      reply.User.UserAddress.City,
		Province:  reply.User.UserAddress.Province,
		AvatarUrl: reply.User.Avatar,
	}
	pkg.Ok(c, res)
	return
}

// AddressList 地址列表
func (u *UserService) AddressList(c *gin.Context) {
	var req AddressListRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		u.helper.Errorf("bind err: %v", err)
		pkg.Validator(c, err)
		return
	}
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	addressReq := &user.AddressListRequest{}
	reply, err := u.client.AddressList(ctx, addressReq)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	res := &AddressListResponse{
		Provinces: reply.AddressList.Provinces,
	}
	pkg.Ok(c, res)
	return
}

// BindPhone 绑定手机号
func (u *UserService) BindPhone(c *gin.Context) {
	var req BindPhoneRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		u.helper.Errorf("bind err: %v", err)
		pkg.Validator(c, err)
		return
	}
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	userToken, ok := u.GetUserFromToken(c)
	if !ok {
		return
	}
	bindReq := &user.BindPhoneRequest{
		Uid:   pkg.ParseInt64(userToken.Uid),
		Phone: req.Phone,
	}
	reply, err := u.client.BindPhone(ctx, bindReq)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	pkg.Ok(c, reply)
	return
}

// GetAvatar 获取头像
func (u *UserService) GetAvatar(c *gin.Context) {
	avatarUrl := c.Query("avatar_url")
	if avatarUrl == "" {
		pkg.FailMessage(c, "avatar_url not exists")
		return
	}
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	avatarReq := &user.GetAvatarRequest{
		AvatarUrl: avatarUrl,
	}
	getAvatarStream, err := u.client.GetAvatar(ctx, avatarReq)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	//设置响应头，返回的是一个流
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+avatarUrl)
	imageType := strings.Split(avatarUrl, ".")[1]
	c.Header("Content-Transfer-Encoding", "image/"+imageType)
	for {
		userResp, err := getAvatarStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			pkg.Fail(c, err)
			return
		}
		_, err = c.Writer.Write(userResp.GetData())
		if err != nil {
			pkg.Fail(c, err)
			return
		}
	}
}

// UploadAvatar 上传头像
func (u *UserService) UploadAvatar(c *gin.Context) {
	ctx, cancel := pkg.NewContext(c)
	defer cancel()
	uploadAvatarStream, err := u.client.UploadAvatar(ctx)
	err = pkg.HandlerError(c, err)
	if err != nil {
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		u.helper.Errorf("get file err: %v", err)
		pkg.FailMessage(c, "get file err")
		return
	}
	f, err := file.Open()
	if err != nil {
		u.helper.Errorf("open file err: %v", err)
		pkg.FailMessage(c, "open file err")
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			u.helper.Errorf("close file err: %v", err)
			pkg.FailMessage(c, "close file err")
			return
		}
	}(f)
	ok := true
	userToken := &auth.UserToken{
		Uid: "",
	}
	if c.Request.RequestURI != "/api/v1/avatar" {
		userToken, ok = u.GetUserFromToken(c)
		if !ok {
			return
		}
	}
	imageName := strings.Split(file.Filename, ".")[0]
	imageType := strings.Split(file.Filename, ".")[1]
	imageInfo := &user.ImageInfo{
		ImageName: imageName,
		ImageType: imageType,
		AccountId: userToken.AccountId,
	}
	userReq := &user.UploadAvatarRequest{
		Data: &user.UploadAvatarRequest_Info{
			Info: imageInfo,
		},
	}
	err = uploadAvatarStream.Send(userReq)
	if err != nil {
		u.helper.Errorf("send err: %v", err)
		pkg.Fail(c, err)
		return
	}
	buf := make([]byte, 1<<20)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			u.helper.Errorf("read err: %v", err)
			pkg.Fail(c, err)
			return
		}
		userReq := &user.UploadAvatarRequest{
			Data: &user.UploadAvatarRequest_ChunkData{
				ChunkData: buf[:n],
			},
		}
		err = uploadAvatarStream.Send(userReq)
		if err != nil {
			u.helper.Errorf("send err: %v", err)
			pkg.Fail(c, err)
			return
		}
	}
	reply, err := uploadAvatarStream.CloseAndRecv()
	if err != nil && err != io.EOF {
		pkg.Fail(c, err)
		return
	}
	token := &auth.UserToken{
		Uid:          pkg.FormatInt(reply.User.UserId),
		NickName:     reply.User.NickName,
		AccountId:    reply.User.AccountId,
		PersonalDesc: reply.User.PersonalDesc,
		Avatar:       reply.User.Avatar,
	}
	tokenString, err := u.Jwt.SignUser(token)
	resp := &UploadAvatarResponse{
		Token: tokenString,
	}
	pkg.Ok(c, resp)
}
