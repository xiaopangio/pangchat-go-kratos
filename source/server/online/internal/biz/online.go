package biz

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"online/api/v1/online"
	"online/internal/components/redis"
	"online/internal/data/model"
	"online/pkg"
)

type OnlineBiz struct {
	helper   *log.Helper
	redisCli *redis.Redis
}

func (o *OnlineBiz) RegisterDevice(ctx context.Context, uid int64, url string) error {
	m := model.Device{
		Uid: pkg.FormatInt(uid),
		Url: url,
	}
	bytes, err := json.Marshal(&m)
	if err != nil {
		err = pkg.InternalError("json编码失败: %v", err)
		o.helper.Errorf(err.Error())
		return err
	}
	if err = pkg.ContextErr(ctx); err != nil {
		return err
	}
	err = o.redisCli.Set(redis.OnlineDeviceKey+pkg.FormatInt(uid), string(bytes), 0)
	if err != nil {
		err = pkg.InternalError("redis set 失败: %v", err)
		o.helper.Errorf(err.Error())
		return err
	}
	return nil
}

func (o *OnlineBiz) UnregisterDevice(ctx context.Context, uid int64) error {
	if err := pkg.ContextErr(ctx); err != nil {
		return err
	}
	err := o.redisCli.Del(redis.OnlineDeviceKey + pkg.FormatInt(uid))
	if err != nil {
		err = pkg.InternalError("redis del 失败: %v", err)
		o.helper.Errorf(err.Error())
		return err
	}
	return nil
}

func (o *OnlineBiz) GetOnlineDevice(ctx context.Context, uid int64) (*model.Device, error) {
	if err := pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	v, err := o.redisCli.Get(redis.OnlineDeviceKey + pkg.FormatInt(uid))
	if err != nil {
		err = pkg.InternalError("redis get 失败: %v", err)
		o.helper.Errorf(err.Error())
		return nil, err
	}
	var m model.Device
	err = json.Unmarshal([]byte(v), &m)
	if err != nil {
		err = pkg.InternalError("json unmarshal 失败: %v", err)
		o.helper.Errorf(err.Error())
		return nil, err
	}
	return &m, nil
}

func (o *OnlineBiz) GetOnlineDevices(ctx context.Context) ([]*online.OnlineDevice, error) {
	values, err := o.redisCli.GetPrefix(redis.OnlineDeviceKey)
	if err != nil {
		err = pkg.InternalError("redis get prefix 失败: %v", err)
		o.helper.Errorf(err.Error())
		return nil, err
	}
	var devices []*online.OnlineDevice
	if err = pkg.ContextErr(ctx); err != nil {
		return nil, err
	}
	for _, v := range values {
		var m online.OnlineDevice
		err = json.Unmarshal([]byte(v), &m)
		if err != nil {
			err = pkg.InternalError("json unmarshal 失败: %v", err)
			o.helper.Errorf(err.Error())
			return nil, err
		}
		devices = append(devices, &m)
	}
	return devices, nil
}

func NewOnlineBiz(helper *log.Helper, redisCli *redis.Redis) *OnlineBiz {
	return &OnlineBiz{helper: helper, redisCli: redisCli}
}
