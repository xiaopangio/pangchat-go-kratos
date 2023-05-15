//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"relationship/internal/biz"
	"relationship/internal/components"
	"relationship/internal/conf"
	"relationship/internal/data"
	"relationship/internal/server"
	"relationship/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(biz.ProviderSet, components.ProviderSet, server.ProviderSet, service.ProviderSet, data.ProviderSet, newApp))
}
