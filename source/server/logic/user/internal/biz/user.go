package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"io"
	"math/rand"
	"strconv"
	"time"
	"user/api/v1/relationship"
	"user/api/v1/user"
	"user/internal/common"
	"user/internal/common/constant"
	oss2 "user/internal/components/oss"
	"user/internal/components/redis"
	"user/internal/components/sms"
	"user/internal/data/orm/model"
	"user/pkg"
)

type UserBiz struct {
	repo               UserRepo
	helper             *log.Helper
	redisCli           *redis.Redis
	smsCli             *sms.SmsClient
	ossClient          *oss2.OSSClient
	relationshipClient relationship.RelationShipClient
	uidGen             *snowflake.Node
}

func NewUserBiz(repo UserRepo, helper *log.Helper, redisCli *redis.Redis, smsCli *sms.SmsClient, ossClient *oss2.OSSClient, relationshipClient relationship.RelationShipClient, uidGen *snowflake.Node) *UserBiz {
	return &UserBiz{repo: repo, helper: helper, redisCli: redisCli, smsCli: smsCli, relationshipClient: relationshipClient, uidGen: uidGen, ossClient: ossClient}
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindUserByPhone(ctx context.Context, phone string) (*model.User, error)
	FindUserByUID(ctx context.Context, uid int64) (*model.User, error)
	FindUserByAccountID(ctx context.Context, accountID string) (*model.User, error)
	ModifyPassword(ctx context.Context, uid int64, password string) error
	GetAddress(ctx context.Context, cityId string) (*user.UserAddress, error)
	UpdateUser(ctx context.Context, user *model.User) error
	UpdateAccountID(ctx context.Context, uid int64, accountID, expire string) error
	UpdatePhone(ctx context.Context, uid int64, phone string) error
	UpdateAvatar(ctx context.Context, accountId, avatar string) error
	GetAddressList(ctx context.Context) (*user.AddressList, error)
	FindProfiles(ctx context.Context, uids []int64) ([]*model.User, error)
}

// Register 注册
func (u *UserBiz) Register(ctx context.Context, t int64, username, password, nickName, avatarUrl string) error {
	var us *model.User
	var err error
	u.helper.Infof("注册用户: %v, %v", t, username)
	//检查注册类型
	if t == constant.Phone { //手机号注册
		us, err = u.repo.FindUserByPhone(ctx, username)
		if err != nil {
			return err
		}
		if us != nil {
			u.helper.Errorf("该手机号已注册: %s", username)
			return pkg.InvalidArgumentError("该手机号已注册")
		}
		us = &model.User{}
		us.Phone = username
		randString, err := pkg.RandStringWithoutSpecial(8)
		if err != nil {
			u.helper.Errorf("生成随机字符串失败: %s", err.Error())
			return pkg.InternalError("生成随机字符串失败")
		}
		//生成账号 以pangchat_开头
		us.AccountID = fmt.Sprintf("%s%s", constant.AppPrefix, randString)
	} else if t == constant.Account { //账号注册
		us, err = u.repo.FindUserByAccountID(ctx, username)
		if err != nil {
			return err
		}
		if us != nil {
			u.helper.Errorf("该账号已注册: %s", username)
			return pkg.InvalidArgumentError("该账号已注册")
		}
		us = &model.User{}
		us.AccountID = username
	}
	//密码加密
	cryptedPasswd, err := bcrypt.GenerateFromPassword([]byte(password), constant.PasswordCost)
	if err != nil {
		u.helper.Errorf("密码加密失败: %s", err)
		return pkg.InternalError("密码加密失败")
	}
	us.Password = string(cryptedPasswd)
	//生成uid
	us.UID = u.uidGen.Generate().Int64()
	us.NickName = nickName
	us.AvatarURL = avatarUrl
	//创建默认好友分组
	_, err = u.relationshipClient.CreateFriendGroup(ctx, &relationship.CreateFriendGroupRequest{
		UserId:    us.UID,
		GroupName: "我的好友",
	})
	if err != nil {
		return pkg.InternalError("%s:创建默认好友分组失败", us.AccountID)
	}
	//创建用户
	err = u.repo.CreateUser(ctx, us)
	if err != nil {
		return pkg.InternalError("创建用户%s 失败", username)
	}
	return nil
}

// SendSmsCode 发送短信验证码
func (u *UserBiz) SendSmsCode(ctx context.Context, phone string) error {
	//判断手机号是否已经发送过验证码
	_, err := u.redisCli.Get(sms.SmsCodePrefix + phone)
	if err == nil {
		u.helper.Errorf("您已经发送过验证码了")
		return pkg.InvalidArgumentError("您已经发送过验证码了")
	}
	//生成验证码
	code, err := pkg.Code()
	if err != nil {
		u.helper.Errorf("%s:生成验证码失败: %s", phone, err.Error())
		return pkg.InternalError("%s:生成验证码失败", phone)
	}
	// 生成短信请求
	smsRequest, err := u.smsCli.CreateSmsRequest(phone, code)
	if err != nil {
		u.helper.Errorf("%s:生成短信请求失败: %s", phone, err.Error())
		return pkg.InternalError("%s:生成短信请求失败", phone)
	}
	if err = pkg.ContextErr(ctx); err != nil {
		return err
	}
	// 发送短信
	result, err := u.smsCli.SendSms(smsRequest)
	if err != nil {
		u.helper.Errorf("%s:发送短信失败: %s", phone, err.Error())
		return pkg.InternalError("%s:发送短信失败", phone)
	}
	if *result.Body.Code != "OK" {
		u.helper.Errorf("%s:发送短信失败: %s", phone, *result.Body.Message)
		return pkg.InternalError("%s:发送短信失败", phone)
	}
	smsCode := &model.SmsCode{
		ID:         pkg.UuidShort(),
		Phone:      phone,
		BizID:      *result.Body.BizId,
		Code:       code,
		CreateTime: time.Now().Unix(),
	}
	//保存短信验证码
	marshal, err := json.Marshal(smsCode)
	if err != nil {
		u.helper.Errorf("%s:序列化短信验证码失败: %s", phone, err.Error())
		return pkg.InternalError("%s:序列化短信验证码失败", phone)
	}
	err = u.redisCli.Set(sms.SmsCodePrefix+phone, string(marshal), time.Minute*5)
	if err != nil {
		u.helper.Errorf("%s:保存短信验证码失败: %s", phone, err.Error())
		return pkg.InternalError("%s:保存短信验证码失败", phone)
	}
	return nil
}

// VerifySmsCode 验证短信验证码
func (u *UserBiz) VerifySmsCode(ctx context.Context, phone, code string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	err := u.smsCli.VerifySmsCode(phone, code)
	if err != nil {
		return err
	}
	err = u.redisCli.Set(sms.SmsPassPrefix+phone, "", 3*time.Minute)
	if err != nil {
		u.helper.Errorf("%s:保存验证结果失败: %s", phone, err.Error())
		return pkg.InternalError("%s:保存验证结果失败", phone)
	}
	return nil
}

// ModifyPasswd 修改密码
func (u *UserBiz) ModifyPasswd(ctx context.Context, uid int64, oldPasswd, newPasswd string) error {
	us, err := u.repo.FindUserByUID(ctx, uid)
	if err != nil {
		return err
	}
	if us == nil {
		u.helper.Errorf("%s:用户不存在", uid)
		return pkg.InvalidArgumentError("%s:用户不存在", uid)
	}
	err = bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(oldPasswd))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			u.helper.Errorf("%s:密码错误", uid)
			return pkg.InvalidArgumentError("密码错误")
		} else {
			u.helper.Errorf("%s:密码校验失败: %s", uid, err.Error())
			return pkg.InternalError("密码校验失败")
		}
	}
	err = u.repo.ModifyPassword(ctx, uid, newPasswd)
	if err != nil {
		return err
	}
	return nil
}

// ResetPassword 重置密码
func (u *UserBiz) ResetPassword(ctx context.Context, uid int64, phone, password string) error {
	us, err := u.repo.FindUserByUID(ctx, uid)
	if err != nil {
		return err
	}
	if us == nil {
		u.helper.Errorf("%s:用户不存在", uid)
		return pkg.InvalidArgumentError("%s:用户不存在", uid)
	}
	_, err = u.redisCli.Get(sms.SmsPassPrefix + phone)
	if err != nil {
		u.helper.Errorf("%s:请先通过手机验证", uid)
		return pkg.UnauthenticatedError("%s:请先通过手机验证", uid)
	}
	err = u.repo.ModifyPassword(ctx, uid, password)
	if err != nil {
		return err
	}
	err = u.redisCli.Del(sms.SmsPassPrefix + phone)
	if err != nil {
		u.helper.Errorf("%s:删除验证结果失败: %s", uid, err.Error())
		return pkg.InternalError("删除验证结果失败: %s", err)
	}
	return nil
}

// completeProfile 完善用户信息
func (u *UserBiz) completeProfile(ctx context.Context, profile *model.User) (*user.UserProfile, error) {
	userAddress, err := u.repo.GetAddress(ctx, profile.CityID)
	if err != nil {
		return nil, err
	}
	if userAddress == nil {
		userAddress = &user.UserAddress{}
	}
	return &user.UserProfile{
		UserId:       profile.UID,
		AccountId:    profile.AccountID,
		NickName:     profile.NickName,
		UserAddress:  userAddress,
		PersonalDesc: profile.PersonalDesc,
		Avatar:       profile.AvatarURL,
	}, nil
}

// GetProfileByPhone 根据手机号获取用户信息
func (u *UserBiz) GetProfileByPhone(ctx context.Context, phone string) (*user.UserProfile, error) {
	u.helper.Infof("获取用户信息: %s", phone)
	profile, err := u.repo.FindUserByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		u.helper.Errorf("%s:用户不存在", phone)
		return nil, pkg.InvalidArgumentError("%s:用户不存在", phone)
	}
	completeProfile, err := u.completeProfile(ctx, profile)
	if err != nil {
		return nil, err
	}
	return completeProfile, nil
}

// GetProfileByUid 根据uid获取用户信息
func (u *UserBiz) GetProfileByUid(ctx context.Context, uid int64) (*user.UserProfile, error) {
	u.helper.Infof("获取用户信息: %s", uid)
	profile, err := u.repo.FindUserByUID(ctx, uid)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		u.helper.Errorf("%s:用户不存在", uid)
		return nil, pkg.InvalidArgumentError("%s:用户不存在", uid)
	}
	completeProfile, err := u.completeProfile(ctx, profile)
	if err != nil {
		return nil, err
	}
	return completeProfile, nil
}

// GetProfileByAccountID 根据accountId获取用户信息
func (u *UserBiz) GetProfileByAccountID(ctx context.Context, accountId string) (*user.UserProfile, error) {
	u.helper.Infof("获取用户信息: %s", accountId)
	tu, err := u.repo.FindUserByAccountID(ctx, accountId)
	if err != nil {
		return nil, err
	}
	if tu == nil {
		u.helper.Errorf("%s:用户不存在", accountId)
		return nil, pkg.InvalidArgumentError("用户不存在")
	}
	completeProfile, err := u.completeProfile(ctx, tu)
	if err != nil {
		return nil, err
	}
	return completeProfile, nil
}

// ModifyProfile 修改用户信息, 仅限于修改昵称, 个人简介, 头像, 城市
func (u *UserBiz) ModifyProfile(ctx context.Context, user *user.UserProfile) error {
	ur, err := u.repo.FindUserByUID(ctx, user.UserId)
	if err != nil {
		return err
	}
	if ur == nil {
		u.helper.Errorf("%s:用户不存在", user.UserId)
		return pkg.InvalidArgumentError("%s:用户不存在", user.UserId)
	}
	updateUser := &model.User{
		UID:          user.UserId,
		NickName:     user.NickName,
		PersonalDesc: user.PersonalDesc,
		CityID:       user.UserAddress.CityId,
		AvatarURL:    user.Avatar,
	}

	err = u.repo.UpdateUser(ctx, updateUser)
	if err != nil {
		return err
	}
	return nil
}

// ModifyAccountID 修改用户账号
func (u *UserBiz) ModifyAccountID(ctx context.Context, uid int64, accountId string) error {
	us, err := u.repo.FindUserByUID(ctx, uid)
	if err != nil {
		return err
	}
	if us == nil {
		u.helper.Errorf("%s:用户不存在", uid)
		return pkg.InvalidArgumentError("%s:用户不存在", uid)
	}
	if !pkg.CompareTime(us.Expire) {
		u.helper.Errorf("%s:一年内只能修改一次，您已经修改过了", uid)
		return pkg.PermissionDeniedError("%s:一年内只能修改一次，您已经修改过了", uid)
	}
	expire := strconv.FormatInt(time.Now().Add(24*time.Hour*365).Unix(), 10)
	err = u.repo.UpdateAccountID(ctx, uid, accountId, expire)
	if err != nil {
		return err
	}
	return nil
}

// AddressList 获取地址列表
func (u *UserBiz) AddressList(ctx context.Context) (*user.AddressList, error) {
	//先从redis中获取
	result, err := u.redisCli.Get("addressList")
	if err != nil && err != redis.Nil {
		u.helper.Error("获取地址列表失败: %s", err)
		return nil, pkg.InternalError("获取地址列表失败: %s", err)
	}
	if result != "" {
		var addressList user.AddressList
		err = json.Unmarshal([]byte(result), &addressList)
		if err != nil {
			u.helper.Error("json解析地址列表失败: %s", err)
			return nil, pkg.InternalError("json解析地址列表失败: %s", err)
		}
		return &addressList, nil
	}
	//redis中没有，从数据库中获取
	addressList, err := u.repo.GetAddressList(ctx)
	if err != nil {
		return nil, err
	}
	//存入redis
	addressListBytes, err := json.Marshal(addressList)
	if err != nil {
		u.helper.Error("json编码地址列表失败: %s", err)
		return nil, pkg.InternalError("json编码地址列表失败: %s", err)
	}
	err = u.redisCli.Set("addressList", string(addressListBytes), 0)
	if err != nil {
		u.helper.Error("存入地址列表失败: %s", err)
		return nil, pkg.InternalError("存入地址列表失败: %s", err)
	}
	return addressList, nil
}

// BindPhone 绑定手机号
func (u *UserBiz) BindPhone(ctx context.Context, uid int64, phone, code string) error {
	us, err := u.repo.FindUserByUID(ctx, uid)
	if err != nil {
		return err
	}
	if us == nil {
		u.helper.Errorf("%s:用户不存在", uid)
		return pkg.InvalidArgumentError("用户不存在")
	}
	res, err := u.redisCli.Get(sms.SmsCodePrefix + phone)
	if err != nil && err != redis.Nil {
		u.helper.Errorf("获取验证码失败: %s", err)
		return pkg.InternalError("获取验证码失败: %s", err)
	}
	if res == "" {
		u.helper.Errorf("验证码不存在")
		return pkg.InvalidArgumentError("验证码不存在")
	}
	smsCode := model.SmsCode{}
	err = json.Unmarshal([]byte(res), &smsCode)
	if err != nil {
		u.helper.Errorf("解析验证码失败: %s", err)
		return pkg.InternalError("解析验证码失败: %s", err)
	}
	if smsCode.Code != code {
		u.helper.Errorf("验证码错误")
		return pkg.InvalidArgumentError("验证码错误")
	}
	err = u.repo.UpdatePhone(ctx, uid, phone)
	if err != nil {
		return err
	}
	err = u.redisCli.Del(sms.SmsCodePrefix + phone)
	if err != nil {
		u.helper.Errorf("删除验证码失败: %s", err)
		return pkg.InternalError("删除验证结果失败: %s", err)
	}
	return nil
}

func (u *UserBiz) UploadAvatar(ctx context.Context, stream user.User_UploadAvatarServer) error {
	req, err := stream.Recv()
	if err != nil {
		return pkg.InternalError("获取图片信息失败: %s", err)
	}
	imageName := req.GetInfo().ImageName
	imageType := req.GetInfo().ImageType
	accountId := req.GetInfo().AccountId
	imageName = imageName + strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(rand.Int63(), 10)
	imagePath := "img/" + imageName + "." + imageType
	reader := &common.UploadAvatarReaderWriter{
		Stream: stream,
		Buffer: [1 << 20]byte{},
	}
	err = u.ossClient.Bucket.PutObject(imagePath, reader)
	if err != nil {
		return pkg.InternalError("上传图片失败: %s", err)
	}
	if accountId == "" {
		err := stream.SendAndClose(&user.UploadAvatarReply{
			User: &user.UserProfile{
				Avatar: imagePath,
			},
		})
		if err != nil {
			return pkg.InternalError("上传图片失败: %s", err)
		}
		return nil
	}
	err = u.repo.UpdateAvatar(ctx, accountId, imagePath)
	if err != nil {
		return err
	}
	userProfile, err := u.GetProfileByAccountID(ctx, accountId)
	if err != nil {
		return err
	}
	err = stream.SendAndClose(&user.UploadAvatarReply{
		User: userProfile,
	})
	if err != nil {
		return err
	}
	return nil
}
func (u *UserBiz) GetAvatar(ctx context.Context, avatarUrl string, stream user.User_GetAvatarServer) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	reader, err := u.ossClient.Bucket.GetObject(avatarUrl)
	if err != nil {
		return pkg.InternalError("获取图片失败: %s", err)
	}
	defer func(reader io.ReadCloser) {
		err := reader.Close()
		if err != nil {
			u.helper.Error("关闭图片流失败: %s", err)
		}
	}(reader)
	buffer := make([]byte, 1<<20)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return pkg.InternalError("读取图片失败: %s", err)
		}
		err = stream.Send(&user.GetAvatarReply{
			Data: buffer[:n],
		})
		if err != nil {
			return pkg.InternalError("发送图片失败: %s", err)
		}
	}
}

// GetProfiles 根据uids获取用户信息
func (u *UserBiz) GetProfiles(ctx context.Context, uids []int64) ([]*model.User, error) {
	profiles, err := u.repo.FindProfiles(ctx, uids)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

// GetAddressAndDesc 获取用户地址和描述
func (u *UserBiz) GetAddressAndDesc(ctx context.Context, uid int64) (*model.User, *user.UserAddress, error) {
	us, err := u.repo.FindUserByUID(ctx, uid)
	address, err := u.repo.GetAddress(ctx, us.CityID)
	if err != nil {
		return us, nil, err
	}
	return us, address, nil
}
