// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"job/internal/biz"
	"job/internal/components/client"
	"job/internal/components/endpoints"
	"job/internal/components/logger"
	"job/internal/components/redis"
	"job/internal/components/registry"
	"job/internal/conf"
	"job/internal/server"
	"job/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, logLogger log.Logger) (*kratos.App, func(), error) {
	helper := logger.NewHelper(logLogger)
	clientv3Client, err := registry.NewEtcdClient(bootstrap, helper)
	if err != nil {
		return nil, nil, err
	}
	jobRegistry := registry.NewEtcdJobRegistry(bootstrap, clientv3Client)
	v := endpoints.NewEndPoints(bootstrap)
	redisRedis := redis.NewRedisClient(bootstrap, helper)
	onlineRegistry := registry.NewOnlineRegistry(bootstrap, clientv3Client)
	onlineClient, err := client.NewOnlineClient(onlineRegistry, helper, bootstrap)
	if err != nil {
		return nil, nil, err
	}
	relationshipRegistry := registry.NewEtcdRelationshipRegistry(bootstrap, clientv3Client)
	relationShipClient, err := client.NewRelationshipClient(relationshipRegistry, helper, bootstrap)
	if err != nil {
		return nil, nil, err
	}
	userRegistry := registry.NewEtcdUserRegistry(bootstrap, clientv3Client)
	userClient, err := client.NewUserClient(userRegistry, helper, bootstrap)
	if err != nil {
		return nil, nil, err
	}
	jobBiz := biz.NewJobBiz(helper, redisRedis, onlineClient, relationShipClient, userClient)
	jobService := service.NewJobService(helper, jobBiz)
	kafkaServer, cleanup := server.NewKafkaConsumerServer(bootstrap, jobService)
	app := newApp(logLogger, bootstrap, jobRegistry, v, kafkaServer)
	return app, func() {
		cleanup()
	}, nil
}
