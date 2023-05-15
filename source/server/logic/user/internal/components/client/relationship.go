package client

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/filter"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"user/api/v1/relationship"
	"user/internal/components/registry"
	"user/internal/conf"
)

func NewRelationshipClient(registry *registry.UserRegistry, logger *log.Helper, cf *conf.Bootstrap) (relationship.RelationShipClient, error) {
	version := filter.Version("1.0")
	selector.SetGlobalSelector(wrr.NewBuilder())
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+cf.Service.RelationshipService),
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
	return relationship.NewRelationShipClient(conn), nil
}
