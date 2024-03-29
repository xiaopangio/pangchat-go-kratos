// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/biz"
	"user/internal/components/client"
	"user/internal/components/logger"
	"user/internal/components/mysql"
	"user/internal/components/oss"
	"user/internal/components/redis"
	"user/internal/components/registry"
	"user/internal/components/sms"
	"user/internal/components/uid"
	"user/internal/conf"
	"user/internal/data"
	"user/internal/server"
	"user/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, logLogger log.Logger) (*kratos.App, func(), error) {
	helper := logger.NewHelper(logLogger)
	redisRedis := redis.NewRedisClient(bootstrap, helper)
	db := mysql.NewMysql(bootstrap)
	dataData, cleanup, err := data.NewData(bootstrap, redisRedis, db, logLogger)
	if err != nil {
		return nil, nil, err
	}
	node := uid.NewUidGenerator(bootstrap, helper)
	userRepo := data.NewUserRepo(dataData, node, helper)
	smsClient := sms.NewSmsClient(bootstrap, redisRedis)
	stsClient := oss.NewSTS(bootstrap)
	ossClient := oss.NewOSS(bootstrap, stsClient)
	clientv3Client, err := registry.NewEtcdClient(bootstrap, helper)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	relationshipRegistry := registry.NewEtcdRelationshipRegistry(bootstrap, clientv3Client)
	relationShipClient, err := client.NewRelationshipClient(relationshipRegistry, helper, bootstrap)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userBiz := biz.NewUserBiz(userRepo, helper, redisRedis, smsClient, ossClient, relationShipClient, node)
	userService := service.NewUserService(userBiz, helper)
	grpcServer := server.NewGRPCServer(bootstrap, userService, logLogger)
	userRegistry := registry.NewEtcdUserRegistry(bootstrap, clientv3Client)
	app := newApp(logLogger, bootstrap, grpcServer, userRegistry)
	return app, func() {
		cleanup()
	}, nil
}
