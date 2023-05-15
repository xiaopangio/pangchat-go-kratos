package client

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/filter"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"relationship/api/v1/user"
	"relationship/internal/components/registry"
	"relationship/internal/conf"
)

func NewUserClient(registry *registry.UserRegistry, logger *log.Helper, cf *conf.Bootstrap) (user.UserClient, error) {
	version := filter.Version("1.0")
	selector.SetGlobalSelector(wrr.NewBuilder())
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+cf.Service.UserService),
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
	return user.NewUserClient(conn), nil
}
