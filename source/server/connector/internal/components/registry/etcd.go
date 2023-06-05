package registry

import (
	"connector/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
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

type ConnectorRegistry struct {
	*etcd.Registry
}

func NewConnectorRegistry(cf *conf.Bootstrap, client *clientv3.Client) *ConnectorRegistry {
	service := cf.Service
	return &ConnectorRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.ConnectorService)),
	}
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

type MessageRegistry struct {
	*etcd.Registry
}

func NewMessageRegistry(cf *conf.Bootstrap, client *clientv3.Client) *MessageRegistry {
	service := cf.Service
	return &MessageRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.MessageService)),
	}
}
