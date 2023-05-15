package registry

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"job/internal/conf"
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

func NewEtcdConnectorRegistry(cf *conf.Bootstrap, client *clientv3.Client) *ConnectorRegistry {
	service := cf.Service
	return &ConnectorRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.ConnectorService)),
	}
}

type JobRegistry struct {
	*etcd.Registry
}

func NewEtcdJobRegistry(cf *conf.Bootstrap, client *clientv3.Client) *JobRegistry {
	service := cf.Service
	return &JobRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.JobService)),
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

type RelationshipRegistry struct {
	*etcd.Registry
}

func NewEtcdRelationshipRegistry(cf *conf.Bootstrap, client *clientv3.Client) *RelationshipRegistry {
	return &RelationshipRegistry{
		Registry: etcd.New(client, etcd.Namespace(cf.Service.RelationshipService)),
	}
}

type UserRegistry struct {
	*etcd.Registry
}

func NewEtcdUserRegistry(cf *conf.Bootstrap, client *clientv3.Client) *UserRegistry {
	return &UserRegistry{
		Registry: etcd.New(client, etcd.Namespace(cf.Service.UserService)),
	}
}
