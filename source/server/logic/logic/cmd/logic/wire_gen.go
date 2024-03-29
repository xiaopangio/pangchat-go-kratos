// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"logic/internal/biz"
	"logic/internal/components/endpoints"
	"logic/internal/components/loadbalance"
	"logic/internal/components/logger"
	"logic/internal/components/mysql"
	"logic/internal/components/oss"
	"logic/internal/components/redis"
	"logic/internal/components/registry"
	"logic/internal/conf"
	"logic/internal/data"
	"logic/internal/server"
	"logic/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, logLogger log.Logger) (*kratos.App, func(), error) {
	helper := logger.NewHelper(logLogger)
	stsClient := oss.NewSTS(bootstrap)
	ossClient := oss.NewOSS(bootstrap, stsClient)
	redisRedis := redis.NewRedisClient(bootstrap, helper)
	client, err := registry.NewEtcdClient(bootstrap, helper)
	if err != nil {
		return nil, nil, err
	}
	connectorRegistry := registry.NewEtcdConnectorRegistry(bootstrap, client)
	loadBalance := loadbalance.NewRandomLoadBalance()
	db := mysql.NewMysql(bootstrap)
	dataData, cleanup, err := data.NewData(bootstrap, logLogger, db)
	if err != nil {
		return nil, nil, err
	}
	logicRepo := data.NewLogicRepoImpl(dataData, helper)
	logicBiz := biz.NewLogicBiz(helper, ossClient, redisRedis, connectorRegistry, loadBalance, bootstrap, logicRepo)
	logicService := service.NewLogicService(logicBiz, helper)
	grpcServer := server.NewGRPCServer(bootstrap, logLogger, logicService)
	v := endpoints.NewEndPoints(bootstrap)
	logicRegistry := registry.NewEtcdLogicRegistry(bootstrap, client)
	app := newApp(logLogger, bootstrap, grpcServer, v, logicRegistry)
	return app, func() {
		cleanup()
	}, nil
}
