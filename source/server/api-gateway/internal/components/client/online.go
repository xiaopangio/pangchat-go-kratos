package client

import (
	"api-gateway/api/v1/online"
	"api-gateway/internal/components/registry"
	"api-gateway/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/selector/filter"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewOnlineClient(registry *registry.OnlineRegistry, logger *log.Helper, service *conf.Service) (online.OnlineClient, error) {
	version := filter.Version("1.0")
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+service.OnlineService),
		grpc.WithDiscovery(registry),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		grpc.WithNodeFilter(version),
	)
	if err != nil {
		logger.Errorw("kind", "grpc-client", "reason", "GRPC_CLIENT_INIT_ERROR", "err", err)
		return nil, err
	}
	return online.NewOnlineClient(conn), nil
}
