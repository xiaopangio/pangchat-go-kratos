//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"api-gateway/internal/components"
	"api-gateway/internal/conf"
	"api-gateway/internal/server"
	"api-gateway/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Service, *conf.Jwt, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(components.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}
