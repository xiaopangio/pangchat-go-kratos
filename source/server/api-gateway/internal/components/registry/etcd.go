package registry

import (
	"api-gateway/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewEtcdClient(bc *conf.Bootstrap, logger *log.Helper) (*clientv3.Client, error) {
	r := bc.Registry
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

type UserRegistry struct {
	*etcd.Registry
}

func NewEtcdUserRegistry(bc *conf.Bootstrap, client *clientv3.Client) *UserRegistry {
	service := bc.Service
	return &UserRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.UserService)),
	}
}

type ConnectorRegistry struct {
	*etcd.Registry
}

func NewEtcdConnectorRegistry(bc *conf.Bootstrap, client *clientv3.Client) *ConnectorRegistry {
	service := bc.Service
	return &ConnectorRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.ConnectorService)),
	}
}

type LogicRegistry struct {
	*etcd.Registry
}

func NewEtcdLogicRegistry(bc *conf.Bootstrap, client *clientv3.Client) *LogicRegistry {
	service := bc.Service
	return &LogicRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.LogicService)),
	}
}

type RelationshipRegistry struct {
	*etcd.Registry
}

func NewEtcdRelationshipRegistry(bc *conf.Bootstrap, client *clientv3.Client) *RelationshipRegistry {
	service := bc.Service
	return &RelationshipRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.RelationshipService)),
	}
}

type OnlineRegistry struct {
	*etcd.Registry
}

func NewOnlineRegistry(bc *conf.Bootstrap, client *clientv3.Client) *OnlineRegistry {
	service := bc.Service
	return &OnlineRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.OnlineService)),
	}
}

type MessageRegistry struct {
	*etcd.Registry
}

func NewMessageRegistry(bc *conf.Bootstrap, client *clientv3.Client) *MessageRegistry {
	service := bc.Service
	return &MessageRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.MessageService)),
	}
}
