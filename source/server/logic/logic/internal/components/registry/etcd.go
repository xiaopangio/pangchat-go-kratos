package registry

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"logic/internal/conf"
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

type LogicRegistry struct {
	*etcd.Registry
}

func NewEtcdLogicRegistry(cf *conf.Bootstrap, client *clientv3.Client) *LogicRegistry {
	service := cf.Service
	return &LogicRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.LogicService)),
	}
}

type ConnectorRegistry struct {
	*etcd.Registry
}

func NewEtcdConnectorRegistry(cf *conf.Bootstrap, client *clientv3.Client) *ConnectorRegistry {
	service := cf.Service
	return &ConnectorRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.ConnectorService)),
	}
}
