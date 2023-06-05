//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"logic/internal/biz"
	"logic/internal/components"
	"logic/internal/conf"
	"logic/internal/data"
	"logic/internal/server"
	"logic/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(biz.ProviderSet, components.ProviderSet, data.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}
