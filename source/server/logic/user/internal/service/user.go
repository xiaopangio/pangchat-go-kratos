package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "user/api/v1/user"
	"user/internal/biz"
	"user/pkg"
)

// UserService 用户服务
type UserService struct {
	pb.UnimplementedUserServer
	bz     *biz.UserBiz
	helper *log.Helper
}

// NewUserService 创建用户服务
func NewUserService(bz *biz.UserBiz, helper *log.Helper) *UserService {
	return &UserService{
		bz:     bz,
		helper: helper,
	}
}

// RegisterUser 注册用户
func (u *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserReply, error) {
	if req.Password != req.PasswordConfirm {
		u.helper.Error("两次密码不一致，请重新输入")
		return nil, pkg.InvalidArgumentError("两次密码不一致，请重新输入")
	}
	err := u.bz.Register(ctx, req.Type, req.Username, req.Password, req.NickName, req.AvatarUrl)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterUserReply{}, nil
}

// SendSmsCode 发送短信验证码
func (u *UserService) SendSmsCode(ctx context.Context, req *pb.SendSmsCodeRequest) (*pb.SendSmsCodeReply, error) {
	err := u.bz.SendSmsCode(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	return &pb.SendSmsCodeReply{}, nil
}

// ModifyPasswd 修改密码
func (u *UserService) ModifyPasswd(ctx context.Context, req *pb.ModifyPasswdRequest) (*pb.ModifyPasswdReply, error) {
	if req.NewPassword != req.NewPasswordConfirm {
		u.helper.Error("密码不一致")
		return nil, pkg.InvalidArgumentError("两次密码不一致，请重新输入")
	}
	err := u.bz.ModifyPasswd(ctx, req.Uid, req.OldPassword, req.NewPassword)
	if err != nil {
		return nil, err
	}
	return &pb.ModifyPasswdReply{}, nil
}

// ResetPassword 重置密码
func (u *UserService) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordReply, error) {
	if req.Password != req.PasswordConfirm {
		u.helper.Error("密码不一致")
		return nil, pkg.InvalidArgumentError("两次密码不一致，请重新输入")
	}
	if err := u.bz.ResetPassword(ctx, req.Uid, req.Phone, req.Password); err != nil {
		return nil, err
	}
	return &pb.ResetPasswordReply{}, nil
}

// ModifyAccountID 修改账号ID
func (u *UserService) ModifyAccountID(ctx context.Context, req *pb.ModifyAccountIDRequest) (*pb.ModifyAccountIDReply, error) {
	if err := u.bz.ModifyAccountID(ctx, req.Uid, req.AccountId); err != nil {
		return nil, err
	}
	return &pb.ModifyAccountIDReply{}, nil
}

// ModifyProfile 修改用户资料
func (u *UserService) ModifyProfile(ctx context.Context, req *pb.ModifyProfileRequest) (*pb.ModifyProfileReply, error) {
	err := u.bz.ModifyProfile(ctx, req.User)
	if err != nil {
		return nil, err
	}
	return &pb.ModifyProfileReply{}, nil
}

// Profile 获取用户资料
func (u *UserService) Profile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileReply, error) {
	//先根据手机号获取用户资料
	user, err := u.bz.GetProfileByPhone(ctx, req.AccountId)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return &pb.ProfileReply{
			User: user,
		}, nil
	}
	//再根据账号ID获取用户资料
	user, err = u.bz.GetProfileByAccountID(ctx, req.AccountId)
	if err != nil {
		return nil, err
	}
	return &pb.ProfileReply{
		User: user,
	}, nil
}

// AddressList 获取地址列表
func (u *UserService) AddressList(ctx context.Context, _ *pb.AddressListRequest) (*pb.AddressListReply, error) {
	addressList, err := u.bz.AddressList(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.AddressListReply{
		AddressList: addressList,
	}, nil
}

// BindPhone 绑定手机号
func (u *UserService) BindPhone(ctx context.Context, req *pb.BindPhoneRequest) (*pb.BindPhoneReply, error) {
	err := u.bz.BindPhone(ctx, req.Uid, req.Phone, req.SmsCode)
	if err != nil {
		return nil, err
	}
	return &pb.BindPhoneReply{}, nil
}

// UploadAvatar 上传头像
func (u *UserService) UploadAvatar(stream pb.User_UploadAvatarServer) error {
	err := u.bz.UploadAvatar(stream.Context(), stream)
	if err != nil {
		return err
	}
	return nil
}

// GetAvatar 获取头像
func (u *UserService) GetAvatar(req *pb.GetAvatarRequest, stream pb.User_GetAvatarServer) error {
	err := u.bz.GetAvatar(stream.Context(), req.AvatarUrl, stream)
	if err != nil {
		return err
	}
	return nil
}

// VerifyCode 验证短信验证码
func (u *UserService) VerifyCode(ctx context.Context, req *pb.VerifyCodeRequest) (*pb.VerifyCodeReply, error) {
	err := u.bz.VerifySmsCode(ctx, req.Phone, req.SmsCode)
	if err != nil {
		return nil, err
	}
	return &pb.VerifyCodeReply{}, nil
}

// Ping ping
func (u *UserService) Ping(context.Context, *pb.PingRequest) (*pb.PingReply, error) {
	u.helper.Info("ping")
	return &pb.PingReply{}, nil
}

// GetProfiles 获取用户资料
func (u *UserService) GetProfiles(ctx context.Context, req *pb.GetProfilesRequest) (*pb.GetProfilesReply, error) {
	profiles, err := u.bz.GetProfiles(ctx, req.Uids)
	if err != nil {
		return nil, err
	}
	var users []*pb.ShortProfile
	for _, profile := range profiles {
		users = append(users, &pb.ShortProfile{
			Uid:      profile.UID,
			NickName: profile.NickName,
			Avatar:   profile.AvatarURL,
		})
	}
	return &pb.GetProfilesReply{
		Profiles: users,
	}, nil
}

// GetAddressAndDesc 获取地址和描述
func (u *UserService) GetAddressAndDesc(ctx context.Context, req *pb.GetAddressAndDescRequest) (*pb.GetAddressAndDescReply, error) {
	us, address, err := u.bz.GetAddressAndDesc(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &pb.GetAddressAndDescReply{
		CityName:     address.City,
		ProvinceName: address.Province,
		Desc:         us.PersonalDesc,
		AccountId:    us.AccountID,
	}, nil
}

// GetProfileByUID 获取用户信息by uid
func (u *UserService) GetProfileByUID(ctx context.Context, req *pb.GetProfileByUIDRequest) (*pb.GetProfileByUIDResponse, error) {
	profile, err := u.bz.GetProfileByUid(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileByUIDResponse{
		User: profile,
	}, nil
}
