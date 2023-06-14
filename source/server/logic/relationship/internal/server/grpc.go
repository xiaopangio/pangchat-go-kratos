package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"relationship/api/v1/relationship"
	"relationship/internal/conf"
	"relationship/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, logger log.Logger, service *service.RelationShipService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			//	log中间件
			logging.Server(logger),
		),
	}
	serverConf := c.Server.Grpc
	if serverConf.Network != "" {
		opts = append(opts, grpc.Network(serverConf.Network))
	}
	if serverConf.Addr != "" {
		opts = append(opts, grpc.Address(serverConf.Addr))
	}
	if serverConf.Timeout != nil {
		opts = append(opts, grpc.Timeout(serverConf.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	relationship.RegisterRelationShipServer(srv, service)
	return srv
}
