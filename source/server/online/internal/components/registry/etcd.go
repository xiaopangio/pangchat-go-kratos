package registry

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"online/internal/conf"
)

func NewEtcdClient(cf *conf.Bootstrap, logger *log.Helper) (*clientv3.Client, error) {
	r := cf.Registry
	client, err := clientv3.New(
		clientv3.Config{
			Endpoints: r.Etcd.Addrs,
		},
	)
	if err != nil {
		logger.Errorf("NewEtcdClient err:%v", err)
		return nil, err
	}
	return client, nil
}

type OnlineRegistry struct {
	*etcd.Registry
}

func NewOnlineRegistry(cf *conf.Bootstrap, client *clientv3.Client) *OnlineRegistry {
	service := cf.Service
	return &OnlineRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.OnlineService)),
	}
}
