package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "user/api/v1/user"
	"user/internal/biz"
	"user/pkg"
)

type UserService struct {
	pb.UnimplementedUserServer
	bz     *biz.UserBiz
	helper *log.Helper
}

func NewUserService(bz *biz.UserBiz, helper *log.Helper) *UserService {
	return &UserService{
		bz:     bz,
		helper: helper,
	}
}
func (u *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserReply, error) {
	if req.Type != 1 && req.Type != 2 {
		err := pkg.InvalidArgumentError("request: type is not invalid")
		u.helper.Error(err.Error())
		return nil, err
	}
	if req.Password != req.PasswordConfirm {
		err := pkg.InvalidArgumentError("request: the two passwords don't match")
		u.helper.Error(err.Error())
		return nil, err
	}
	err := u.bz.Register(ctx, req.Type, req.Username, req.Password, req.NickName, req.AvatarUrl)
	if err != nil {
		u.helper.Error(err.Error())
		return nil, err
	}
	return &pb.RegisterUserReply{}, nil
}
func (u *UserService) SendSmsCode(ctx context.Context, req *pb.SendSmsCodeRequest) (*pb.SendSmsCodeReply, error) {
	err := u.bz.SendSmsCode(ctx, req.Phone)
	if err != nil {
		u.helper.Error(err.Error())
		return nil, err
	}
	return &pb.SendSmsCodeReply{}, nil
}
func (u *UserService) ModifyPasswd(ctx context.Context, req *pb.ModifyPasswdRequest) (*pb.ModifyPasswdReply, error) {
	if req.NewPassword != req.NewPasswordConfirm {
		err := pkg.InvalidArgumentError("request: the two passwords don't match")
		u.helper.Error(err.Error())
		return nil, err
	}
	err := u.bz.ModifyPasswd(ctx, req.Uid, req.OldPassword, req.NewPassword)
	if err != nil {
		u.helper.Error(err.Error())
		return nil, err
	}
	return &pb.ModifyPasswdReply{}, nil
}
func (u *UserService) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordReply, error) {
	if req.Password != req.PasswordConfirm {
		err := pkg.InvalidArgumentError("request: the two passwords don't match")
		u.helper.Error(err.Error())
		return nil, err
	}
	if err := u.bz.ResetPassword(ctx, req.Uid, req.Phone, req.Password); err != nil {
		u.helper.Error(err.Error())
		return nil, err
	}
	return &pb.ResetPasswordReply{}, nil
}
func (u *UserService) ModifyAccountID(ctx context.Context, req *pb.ModifyAccountIDRequest) (*pb.ModifyAccountIDReply, error) {
	if err := u.bz.ModifyAccountID(ctx, req.Uid, req.AccountId); err != nil {
		u.helper.Error(err.Error())
		return nil, err
	}
	return &pb.ModifyAccountIDReply{}, nil
}
func (u *UserService) ModifyProfile(ctx context.Context, req *pb.ModifyProfileRequest) (*pb.ModifyProfileReply, error) {
	err := u.bz.ModifyProfile(ctx, req.User)
	if err != nil {
		u.helper.Error(err.Error())
	}
	return &pb.ModifyProfileReply{}, nil
}
func (u *UserService) Profile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileReply, error) {
	user, err := u.bz.Profile(ctx, req.AccountId)
	if err != nil {
		u.helper.Error(err.Error())
		return nil, err
	}
	return &pb.ProfileReply{
		User: user,
	}, nil
}
func (u *UserService) AddressList(ctx context.Context, req *pb.AddressListRequest) (*pb.AddressListReply, error) {
	addressList, err := u.bz.AddressList(ctx)
	if err != nil {
		u.helper.Error(err.Error())
		return nil, err
	}
	return &pb.AddressListReply{
		AddressList: addressList,
	}, nil
}
func (u *UserService) BindPhone(ctx context.Context, req *pb.BindPhoneRequest) (*pb.BindPhoneReply, error) {
	err := u.bz.BindPhone(ctx, req.Uid, req.Phone, req.SmsCode)
	if err != nil {
		u.helper.Error(err.Error())
		return nil, err
	}
	return &pb.BindPhoneReply{}, nil
}
func (u *UserService) UploadAvatar(stream pb.User_UploadAvatarServer) error {
	err := u.bz.UploadAvatar(stream.Context(), stream)
	if err != nil {
		u.helper.Error(err.Error())
		return err
	}
	return nil
}
func (u *UserService) GetAvatar(req *pb.GetAvatarRequest, stream pb.User_GetAvatarServer) error {
	err := u.bz.GetAvatar(stream.Context(), req.AvatarUrl, stream)
	if err != nil {
		u.helper.Error(err.Error())
		return err
	}
	return nil
}
func (u *UserService) VerifyCode(ctx context.Context, req *pb.VerifyCodeRequest) (*pb.VerifyCodeReply, error) {
	err := u.bz.VerifySmsCode(ctx, req.Phone, req.SmsCode)
	if err != nil {
		u.helper.Error(err.Error())
		return nil, err
	}
	return &pb.VerifyCodeReply{}, nil
}
func (u *UserService) Ping(context.Context, *pb.PingRequest) (*pb.PingReply, error) {
	u.helper.Info("ping")
	return &pb.PingReply{}, nil
}
func (u *UserService) GetProfiles(ctx context.Context, req *pb.GetProfilesRequest) (*pb.GetProfilesReply, error) {
	profiles, err := u.bz.GetProfiles(ctx, req.Uids)
	if err != nil {
		u.helper.Error(err.Error())
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
func (u *UserService) GetAddressAndDesc(ctx context.Context, req *pb.GetAddressAndDescRequest) (*pb.GetAddressAndDescReply, error) {
	desc, address, err := u.bz.GetAddressAndDesc(ctx, req.Uid)
	if err != nil {
		u.helper.Error(err.Error())
		return nil, err
	}
	return &pb.GetAddressAndDescReply{
		CityName:     address.City,
		ProvinceName: address.Province,
		Desc:         desc,
	}, nil
}
