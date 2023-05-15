//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"online/internal/biz"
	"online/internal/components"
	"online/internal/conf"
	"online/internal/server"
	"online/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(biz.ProviderSet, components.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}
