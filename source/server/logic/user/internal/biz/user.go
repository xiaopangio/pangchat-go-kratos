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

func (u *UserBiz) Register(ctx context.Context, t int64, username, password, nickName, avatarUrl string) error {
	var us *model.User
	var err error
	u.helper.Infof("register us: %s,type: %d", username, t)
	//检查注册类型
	if t == pkg.Phone { //手机号注册
		us, err = u.repo.FindUserByPhone(ctx, username)
		if err != nil {
			return err
		}
		if us != nil {
			return pkg.AlreadyExistsError("该手机号已注册")
		}
		us = &model.User{}
		us.Phone = username
		randString, err := pkg.RandString(8)
		if err != nil {
			return pkg.InternalError("生成随机字符串失败")
		}
		us.AccountID = fmt.Sprintf("%s%s", randString, username)
	} else if t == pkg.Account { //账号注册
		us, err = u.repo.FindUserByAccountID(ctx, username)
		if err != nil {
			return err
		}
		if us != nil {
			return pkg.AlreadyExistsError("该用户名已注册")
		}
		us = &model.User{}
		us.AccountID = username
	} else {
		return pkg.InvalidArgumentError("非法类型: %v", t)
	}
	//密码加密
	cryptedPasswd, err := bcrypt.GenerateFromPassword([]byte(password), pkg.PasswordCost)
	if err != nil {
		return pkg.InternalError("密码加密错误: %s", err)
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
		return pkg.InternalError("创建默认好友分组失败: %s", err)
	}
	//创建用户
	err = u.repo.CreateUser(ctx, us)
	if err != nil {
		return pkg.InternalError("创建用户失败: %s", err)
	}
	return nil
}
func (u *UserBiz) SendSmsCode(ctx context.Context, phone string) error {
	_, err := u.redisCli.Get(sms.SmsCodePrefix + phone)
	if err == nil {
		u.helper.Error(err.Error())
		return pkg.AlreadyExistsError("您已经发送过验证码了")
	}
	code, err := pkg.Code()
	if err != nil {
		return pkg.InternalError("生成验证码失败: %s", err)
	}
	smsRequest, err := u.smsCli.CreateSmsRequest(phone, code)
	if err != nil {
		return pkg.InternalError("创建短信请求失败: %s", err)
	}
	result, err := u.smsCli.SendSms(smsRequest)
	if err != nil {
		return pkg.InternalError("发送短信失败: %s", err)
	}
	if *result.Body.Code != "OK" {
		return pkg.InternalError("发送短信失败: %s", *result.Body.Message)
	}
	smsCode := &model.SmsCode{
		ID:         pkg.UuidShort(),
		Phone:      phone,
		BizID:      *result.Body.BizId,
		Code:       code,
		CreateTime: time.Now().Unix(),
	}
	marshal, err := json.Marshal(smsCode)
	if err != nil {
		return pkg.InternalError("序列化短信验证码失败: %s", err)
	}
	err = u.redisCli.Set(sms.SmsCodePrefix+phone, string(marshal), time.Minute*5)
	if err != nil {
		return pkg.InternalError("保存短信验证码失败: %s", err)
	}
	return nil
}
func (u *UserBiz) VerifySmsCode(ctx context.Context, phone, code string) error {
	err := u.smsCli.VerifySmsCode(phone, code)
	if err != nil {
		return err
	}
	err = u.redisCli.Set(sms.SmsPassPrefix+phone, "", 3*time.Minute)
	if err != nil {
		return pkg.InternalError("保存验证结果失败: %s", err)
	}
	return nil
}
func (u *UserBiz) ModifyPasswd(ctx context.Context, uid int64, oldPasswd, newPasswd string) error {
	user, err := u.repo.FindUserByUID(ctx, uid)
	if err != nil {
		return err
	}
	if user == nil {
		return pkg.InvalidArgumentError("用户不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPasswd))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return pkg.InvalidArgumentError("密码错误")
		} else {
			return pkg.InternalError("密码验证错误")
		}
	}
	err = u.repo.ModifyPassword(ctx, uid, newPasswd)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserBiz) ResetPassword(ctx context.Context, uid int64, phone, password string) error {
	user, err := u.repo.FindUserByUID(ctx, uid)
	if err != nil {
		return err
	}
	if user == nil {
		return pkg.InvalidArgumentError("用户不存在")
	}
	_, err = u.redisCli.Get(sms.SmsPassPrefix + phone)
	if err != nil {
		return pkg.UnauthenticatedError("请先通过手机验证")
	}
	err = u.repo.ModifyPassword(ctx, uid, password)
	if err != nil {
		return err
	}
	err = u.redisCli.Del(sms.SmsPassPrefix + phone)
	if err != nil {
		return pkg.InternalError("删除验证结果失败: %s", err)
	}
	return nil
}
func (u *UserBiz) Profile(ctx context.Context, accountId string) (*user.UserProfile, error) {
	u.helper.Infof("获取用户信息: %s", accountId)
	tu, err := u.repo.FindUserByAccountID(ctx, accountId)
	if err != nil {
		return nil, err
	}
	if tu == nil {
		return nil, pkg.InvalidArgumentError("用户不存在")
	}
	userAddress, err := u.repo.GetAddress(ctx, tu.CityID)
	if err != nil {
		return nil, err
	}
	if userAddress == nil {
		userAddress = &user.UserAddress{}
	}
	profile := &user.UserProfile{
		UserId:       tu.UID,
		AccountId:    tu.AccountID,
		NickName:     tu.NickName,
		PersonalDesc: tu.PersonalDesc,
		Avatar:       tu.AvatarURL,
		UserAddress:  userAddress,
	}
	return profile, nil
}
func (u *UserBiz) ModifyProfile(ctx context.Context, user *user.UserProfile) error {
	ur, err := u.repo.FindUserByUID(ctx, user.UserId)
	if err != nil {
		return err
	}
	if ur == nil {
		return pkg.InvalidArgumentError("用户不存在")
	}
	address := user.UserAddress
	if address != nil {
		ur.CityID = address.CityId
	}
	if user.NickName != "" {
		ur.NickName = user.NickName
	}
	if user.PersonalDesc != "" {
		ur.PersonalDesc = user.PersonalDesc
	}
	err = u.repo.UpdateUser(ctx, ur)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserBiz) ModifyAccountID(ctx context.Context, uid int64, accountId string) error {
	user, err := u.repo.FindUserByUID(ctx, uid)
	if err != nil {
		return err
	}
	if user == nil {
		return pkg.InvalidArgumentError("用户不存在")
	}
	if !pkg.CompareTime(user.Expire) {
		return pkg.PermissionDeniedError("一年内只能修改一次，您已经修改过了")
	}
	expire := strconv.FormatInt(time.Now().Add(24*time.Hour*365).Unix(), 10)
	err = u.repo.UpdateAccountID(ctx, uid, accountId, expire)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserBiz) AddressList(ctx context.Context) (*user.AddressList, error) {
	//先从redis中获取
	result, err := u.redisCli.Get("addressList")
	if err != nil && err != redis.Nil {
		return nil, pkg.InternalError("获取地址列表失败: %s", err)
	}
	if result != "" {
		var addressList user.AddressList
		err = json.Unmarshal([]byte(result), &addressList)
		if err != nil {
			return nil, pkg.InternalError("解析地址列表失败: %s", err)
		}
		u.helper.Info("从redis中获取地址列表")
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
		return nil, pkg.InternalError("解析地址列表失败: %s", err)
	}
	err = u.redisCli.Set("addressList", string(addressListBytes), 0)
	if err != nil {
		return nil, pkg.InternalError("存入地址列表失败: %s", err)
	}
	return addressList, nil
}
func (u *UserBiz) BindPhone(ctx context.Context, uid int64, phone, code string) error {
	user, err := u.repo.FindUserByUID(ctx, uid)
	if err != nil {
		return err
	}
	if user == nil {
		return pkg.InvalidArgumentError("用户不存在")
	}
	res, err := u.redisCli.Get(sms.SmsCodePrefix + phone)
	if err != nil {
		return pkg.UnauthenticatedError("请先获取验证码")
	}
	smsCode := model.SmsCode{}
	err = json.Unmarshal([]byte(res), &smsCode)
	if err != nil {
		return pkg.InternalError("解析验证码失败: %s", err)
	}
	if smsCode.Code != code {
		return pkg.InvalidArgumentError("验证码错误")
	}
	err = u.repo.UpdatePhone(ctx, uid, phone)
	if err != nil {
		return err
	}
	err = u.redisCli.Del(sms.SmsCodePrefix + phone)
	if err != nil {
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
	reader := &pkg.UploadAvatarReaderWriter{
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
	userProfile, err := u.Profile(ctx, accountId)
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
	reader, err := u.ossClient.Bucket.GetObject(avatarUrl)
	if err != nil {
		return pkg.InternalError("获取图片失败: %s", err)
	}
	defer reader.Close()
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

func (u *UserBiz) GetProfiles(ctx context.Context, uids []int64) ([]*model.User, error) {
	profiles, err := u.repo.FindProfiles(ctx, uids)
	if err != nil {
		u.helper.Error("获取用户信息失败: %s", err)
		return nil, err
	}
	return profiles, nil
}

func (u *UserBiz) GetAddressAndDesc(ctx context.Context, uid int64) (*model.User, *user.UserAddress, error) {
	us, err := u.repo.FindUserByUID(ctx, uid)
	address, err := u.repo.GetAddress(ctx, us.CityID)
	if err != nil {
		u.helper.Error("获取用户信息失败: %s", err)
		return us, nil, err
	}
	return us, address, nil
}
