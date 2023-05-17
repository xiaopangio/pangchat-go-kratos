package registry

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"user/internal/conf"
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

type UserRegistry struct {
	*etcd.Registry
}

func NewEtcdUserRegistry(cf *conf.Bootstrap, client *clientv3.Client) *UserRegistry {
	service := cf.Service
	return &UserRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.UserService)),
	}
}

type RelationshipRegistry struct {
	*etcd.Registry
}

func NewEtcdRelationshipRegistry(cf *conf.Bootstrap, client *clientv3.Client) *RelationshipRegistry {
	service := cf.Service
	return &RelationshipRegistry{
		Registry: etcd.New(client, etcd.Namespace(service.RelationshipService)),
	}
}
