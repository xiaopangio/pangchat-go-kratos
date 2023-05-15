// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"api-gateway/internal/components/auth"
	"api-gateway/internal/components/client"
	"api-gateway/internal/components/logger"
	"api-gateway/internal/components/registry"
	"api-gateway/internal/conf"
	"api-gateway/internal/server"
	"api-gateway/internal/service/connector"
	"api-gateway/internal/service/logic"
	"api-gateway/internal/service/relationship"
	"api-gateway/internal/service/user"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confRegistry *conf.Registry, service *conf.Service, jwt *conf.Jwt, logLogger log.Logger) (*kratos.App, func(), error) {
	helper := logger.NewHelper(logLogger)
	clientv3Client, err := registry.NewEtcdClient(confRegistry, confServer, helper)
	if err != nil {
		return nil, nil, err
	}
	userRegistry := registry.NewEtcdUserRegistry(service, clientv3Client)
	userClient, err := client.NewUserClient(userRegistry, helper, service)
	if err != nil {
		return nil, nil, err
	}
	jwtManager := auth.NewJwtManager(jwt)
	userService := service_user.NewUserService(userClient, helper, jwtManager)
	connectorRegistry := registry.NewEtcdConnectorRegistry(service, clientv3Client)
	connectorServiceClient, err := client.NewConnectorClient(connectorRegistry, helper, service)
	if err != nil {
		return nil, nil, err
	}
	connectorService := service_connector.NewConnectorService(connectorServiceClient, helper, jwtManager)
	logicRegistry := registry.NewEtcdLogicRegistry(service, clientv3Client)
	logicClient, err := client.NewLogicClient(logicRegistry, helper, service)
	if err != nil {
		return nil, nil, err
	}
	logicService := service_logic.NewLogicService(logicClient, helper, jwtManager)
	relationshipRegistry := registry.NewEtcdRelationshipRegistry(service, clientv3Client)
	relationShipClient, err := client.NewRelationshipClient(relationshipRegistry, helper, service)
	if err != nil {
		return nil, nil, err
	}
	relationship := service_relationship.NewRelationship(helper, relationShipClient, jwtManager)
	httpServer := server.NewHTTPServer(confServer, userService, connectorService, logicService, relationship)
	app := newApp(logLogger, httpServer)
	return app, func() {
	}, nil
}
