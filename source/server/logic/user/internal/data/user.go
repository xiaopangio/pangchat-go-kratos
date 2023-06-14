package data

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	v1 "user/api/v1/user"
	"user/internal/biz"
	"user/internal/data/orm/dal"
	"user/internal/data/orm/model"
	"user/pkg"
)

type UserRepoImpl struct {
	data   *Data
	uidGen *snowflake.Node
	helper *log.Helper
}

func (u *UserRepoImpl) FindProfiles(ctx context.Context, uids []int64) ([]*model.User, error) {
	users, err := dal.User.WithContext(ctx).
		Where(dal.User.UID.In(uids...)).
		Select(
			dal.User.UID,
			dal.User.NickName,
			dal.User.AvatarURL,
		).
		Find()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		u.helper.Errorf("查询用户失败: %v", err)
		return nil, err
	}
	return users, nil
}

func (u *UserRepoImpl) UpdateAvatar(ctx context.Context, accountId, avatar string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	_, err := dal.User.WithContext(ctx).
		Where(dal.User.AccountID.Eq(accountId)).
		Updates(
			&model.User{
				AvatarURL: avatar,
			},
		)
	if err != nil {
		u.helper.Errorf("更新avatar失败: %v", err)
		return pkg.InternalError("更新avatar失败: %v", err)
	}
	return nil
}

func (u *UserRepoImpl) UpdatePhone(ctx context.Context, uid int64, phone string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	_, err := dal.User.WithContext(ctx).
		Where(dal.User.UID.Eq(uid)).
		Updates(
			&model.User{
				Phone: phone,
			},
		)
	if err != nil {
		u.helper.Errorf("更新phone失败: %v", err)
		return pkg.InternalError("更新phone失败: %v", err)
	}
	return nil
}

func (u *UserRepoImpl) GetAddressList(ctx context.Context) (*v1.AddressList, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	// 查询省份
	provinces, err := dal.Province.WithContext(ctx).Find()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		u.helper.Errorf("查询省份失败: %v", err)
		return nil, err
	}
	if len(provinces) == 0 {
		u.helper.Errorf("省份为空")
		return nil, pkg.InternalError("province is empty")
	}
	if err = pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	cities, err := dal.City.WithContext(ctx).
		Find()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		u.helper.Errorf("查询城市失败: %v", err)
		return nil, err
	}
	if len(cities) == 0 {
		u.helper.Errorf("城市为空")
		return nil, pkg.InternalError("city is empty")
	}
	addressList := &v1.AddressList{
		Provinces: make([]*v1.Province, len(provinces)),
	}
	for i, province := range provinces {
		addressList.Provinces[i] = &v1.Province{
			ProvinceId:   province.ProvinceID,
			ProvinceName: province.ProvinceName,
			Cities:       make([]*v1.City, 0),
		}
		for _, city := range cities {
			if city.ParentID == province.ProvinceID {
				addressList.Provinces[i].Cities = append(addressList.Provinces[i].Cities, &v1.City{
					CityId:   city.CityID,
					CityName: city.CityName,
				})
			}
		}
	}
	return addressList, nil
}

func (u *UserRepoImpl) UpdateAccountID(ctx context.Context, uid int64, accountID, expire string) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	_, err := dal.User.WithContext(ctx).
		Where(dal.User.UID.Eq(uid)).
		Updates(
			&model.User{
				AccountID: accountID,
				Expire:    expire,
			},
		)
	if err != nil {
		u.helper.Errorf("更新accountID失败: %v", err)
		return pkg.InternalError("更新accountID失败: %v", err)
	}
	return nil
}

func (u *UserRepoImpl) UpdateUser(ctx context.Context, user *model.User) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	_, err := dal.User.WithContext(ctx).
		Where(dal.User.UID.Eq(user.UID)).
		Updates(user)
	if err != nil {
		u.helper.Errorf("更新user失败: %v", err)
		return pkg.InternalError("更新user失败: %v", err)
	}
	return nil
}

func (u *UserRepoImpl) GetAddress(ctx context.Context, cityId string) (*v1.UserAddress, error) {
	if cityId == "" {
		return &v1.UserAddress{}, nil
	}
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	city, err := dal.City.WithContext(ctx).
		Where(dal.City.CityID.Eq(cityId)).
		First()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		u.helper.Errorf("查询城市失败: %v", err)
		return nil, err
	}
	if err = pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	province, err := dal.Province.WithContext(ctx).
		Where(dal.Province.ProvinceID.Eq(city.ParentID)).
		First()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		u.helper.Errorf("查询省份失败: %v", err)
		return nil, err
	}
	return &v1.UserAddress{
		CityId:   cityId,
		City:     city.CityName,
		Province: province.ProvinceName,
	}, nil
}

func (u *UserRepoImpl) ModifyPassword(ctx context.Context, uid int64, password string) error {
	cryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), pkg.PasswordCost)
	if err != nil {
		u.helper.Errorf("密码加密失败: %v", err)
		return pkg.InternalError("密码加密失败")
	}
	if err = pkg.ContextErr(ctx); err != nil {
		return err
	}
	_, err = dal.User.WithContext(ctx).Where(dal.User.UID.Eq(uid)).Update(dal.User.Password, cryptPassword)
	if err != nil {
		u.helper.Errorf("密码更新失败: %v", err)
		return pkg.InternalError("密码更新失败")
	}
	return nil
}

func (u *UserRepoImpl) FindUserByAccountID(ctx context.Context, accountID string) (*model.User, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	user, err := dal.User.WithContext(ctx).
		Where(dal.User.AccountID.Eq(accountID)).
		First()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		u.helper.Errorf("查询失败: %v", err)
		return nil, pkg.InternalError("查询失败: %v", err)
	}
	return user, nil
}

func (u *UserRepoImpl) FindUserByPhone(ctx context.Context, phone string) (*model.User, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	user, err := dal.User.WithContext(ctx).
		Where(dal.User.Phone.Eq(phone)).
		First()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		u.helper.Errorf("查询失败: %v", err)
		return nil, pkg.InternalError("查询失败: %v", err)
	}
	return user, nil
}

func (u *UserRepoImpl) FindUserByUID(ctx context.Context, uid int64) (*model.User, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	user, err := dal.User.WithContext(ctx).
		Where(dal.User.UID.Eq(uid)).
		First()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		u.helper.Errorf("查询失败: %v", err)
		return nil, pkg.InternalError("查询失败: %v", err)
	}
	return user, nil
}

func (u *UserRepoImpl) CreateUser(ctx context.Context, user *model.User) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	return dal.User.WithContext(ctx).Create(user)
}

func NewUserRepo(data *Data, uidGen *snowflake.Node, helper *log.Helper) biz.UserRepo {
	return &UserRepoImpl{data: data, uidGen: uidGen, helper: helper}
}
