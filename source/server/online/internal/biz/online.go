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
	o.helper.Infof("register device: %v", m)
	bytes, err := json.Marshal(&m)
	if err != nil {
		o.helper.Errorf("json marshal failed: %v", err)
		return err
	}
	err = o.redisCli.Set(redis.OnlineDeviceKey+pkg.FormatInt(uid), string(bytes), 0)
	if err != nil {
		o.helper.Errorf("redis set failed: %v", err)
		return err
	}
	return nil
}

func (o *OnlineBiz) UnregisterDevice(ctx context.Context, uid int64) error {
	err := o.redisCli.Del(redis.OnlineDeviceKey + pkg.FormatInt(uid))
	if err != nil {
		o.helper.Errorf("redis del failed: %v", err)
		return err
	}
	return nil
}

func (o *OnlineBiz) GetOnlineDevice(ctx context.Context, uid int64) (*model.Device, error) {
	v, err := o.redisCli.Get(redis.OnlineDeviceKey + pkg.FormatInt(uid))
	if err != nil {
		return nil, err
	}
	var m model.Device
	err = json.Unmarshal([]byte(v), &m)
	if err != nil {
		o.helper.Errorf("json unmarshal failed: %v", err)
		return nil, err
	}
	return &m, nil
}

func (o *OnlineBiz) GetOnlineDevices(ctx context.Context) ([]*online.OnlineDevice, error) {
	values, err := o.redisCli.GetPrefix(redis.OnlineDeviceKey)
	if err != nil {
		o.helper.Errorf("redis get prefix failed: %v", err)
		return nil, err
	}
	var devices []*online.OnlineDevice
	for _, v := range values {
		var m online.OnlineDevice
		err = json.Unmarshal([]byte(v), &m)
		if err != nil {
			o.helper.Errorf("json unmarshal failed: %v", err)
			return nil, err
		}
		devices = append(devices, &m)
	}
	return devices, nil
}

func NewOnlineBiz(helper *log.Helper, redisCli *redis.Redis) *OnlineBiz {
	return &OnlineBiz{helper: helper, redisCli: redisCli}
}
