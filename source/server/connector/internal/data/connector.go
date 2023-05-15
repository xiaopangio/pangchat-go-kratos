package data

import (
	"connector/api/v1/connector"
	"connector/internal/biz"
	"connector/internal/data/orm/dal"
	"connector/internal/data/orm/model"
	"connector/pkg"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ConnectorRepoImpl struct {
	Data   *Data
	helper *log.Helper
}

func NewConnectorRepoImpl(data *Data, helper *log.Helper) biz.ConnectorServiceRepo {
	return &ConnectorRepoImpl{Data: data, helper: helper}
}

func (c *ConnectorRepoImpl) FindUserByPhone(ctx context.Context, phone string) (*model.User, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	user, err := dal.User.WithContext(ctx).Where(dal.User.Phone.Eq(phone)).First()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		return nil, pkg.InternalError(err.Error())
	}
	if user == nil {
		return nil, pkg.NotFoundError("user not found")
	}
	return user, nil
}

func (c *ConnectorRepoImpl) FindUserByAccountId(ctx context.Context, accountId string) (*model.User, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	c.helper.Infof("FindUserByAccountId request: %v", accountId)
	user, err := dal.User.WithContext(ctx).Where(dal.User.AccountID.Eq(accountId)).First()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		return nil, pkg.InternalError(err.Error())
	}
	if user == nil {
		return nil, pkg.NotFoundError("user not found")
	}
	return user, nil
}

func (c *ConnectorRepoImpl) UpdateLoginStatus(ctx context.Context, uid int64, status int) error {
	_, err := dal.User.WithContext(ctx).Where(dal.User.UID.Eq(uid)).Update(dal.User.Status, status)
	if err != nil {
		return pkg.InternalError(err.Error())
	}
	return nil
}

func (c *ConnectorRepoImpl) FindAddressByCityId(ctx context.Context, cityId string) (*connector.UserAddress, error) {
	city, err := dal.City.WithContext(ctx).Where(dal.City.CityID.Eq(cityId)).First()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		return nil, pkg.InternalError(err.Error())
	}
	if city == nil {
		return nil, pkg.NotFoundError("city not found")
	}
	province, err := dal.Province.WithContext(ctx).Where(dal.Province.ProvinceID.Eq(city.ParentID)).First()
	if err = pkg.IsNotRecordNotFoundError(err); err != nil {
		return nil, pkg.InternalError(err.Error())
	}
	if province == nil {
		return nil, pkg.NotFoundError("province not found")
	}
	return &connector.UserAddress{
		CityId:   cityId,
		City:     city.CityName,
		Province: province.ProvinceName,
	}, nil
}
