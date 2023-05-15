// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"logic/internal/biz"
	"logic/internal/components/loadbalance"
	"logic/internal/components/logger"
	"logic/internal/components/redis"
	"logic/internal/components/registry"
	"logic/internal/conf"
	"logic/internal/server"
	"logic/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, logLogger log.Logger, confRegistry *conf.Registry, confService *conf.Service) (*kratos.App, func(), error) {
	helper := logger.NewHelper(logLogger)
	redisRedis := redis.NewRedisClient(data, helper)
	client, err := registry.NewEtcdClient(confRegistry, confServer, helper)
	if err != nil {
		return nil, nil, err
	}
	connectorRegistry := registry.NewEtcdConnectorRegistry(confService, client)
	loadBalance := loadbalance.NewRandomLoadBalance()
	logicBiz := biz.NewLogicBiz(helper, redisRedis, connectorRegistry, loadBalance, confService)
	logicService := service.NewLogicService(logicBiz, helper)
	grpcServer := server.NewGRPCServer(confServer, logLogger, logicService)
	logicRegistry := registry.NewEtcdLogicRegistry(confService, client)
	app := newApp(logLogger, confServer, grpcServer, logicRegistry)
	return app, func() {
	}, nil
}
