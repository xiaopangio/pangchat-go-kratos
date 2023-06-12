package registry

import (
	"api-gateway/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewEtcdClient(r *conf.Registry, s *conf.Server, logger *log.Helper) (*clientv3.Client, error) {
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

func NewEtcdUserRegistry(service *conf.Service, client *clientv3.Client) *UserRegistry {
	return &UserRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.UserService)),
	}
}

type ConnectorRegistry struct {
	*etcd.Registry
}

func NewEtcdConnectorRegistry(service *conf.Service, client *clientv3.Client) *ConnectorRegistry {
	return &ConnectorRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.ConnectorService)),
	}
}

type LogicRegistry struct {
	*etcd.Registry
}

func NewEtcdLogicRegistry(service *conf.Service, client *clientv3.Client) *LogicRegistry {
	return &LogicRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.LogicService)),
	}
}

type RelationshipRegistry struct {
	*etcd.Registry
}

func NewEtcdRelationshipRegistry(service *conf.Service, client *clientv3.Client) *RelationshipRegistry {
	return &RelationshipRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.RelationshipService)),
	}
}

type OnlineRegistry struct {
	*etcd.Registry
}

func NewOnlineRegistry(service *conf.Service, client *clientv3.Client) *OnlineRegistry {
	return &OnlineRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.OnlineService)),
	}
}

type MessageRegistry struct {
	*etcd.Registry
}

func NewMessageRegistry(service *conf.Service, client *clientv3.Client) *MessageRegistry {
	return &MessageRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.MessageService)),
	}
}
