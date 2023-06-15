package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"online/api/v1/online"
	"online/internal/conf"
	"online/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, service *service.OnlineService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			validate.Validator(),
		),
	}
	server := c.Server.Grpc
	if server.Network != "" {
		opts = append(opts, grpc.Network(server.Network))
	}
	if server.Addr != "" {
		opts = append(opts, grpc.Address(server.Addr))
	}
	if server.Timeout != nil {
		opts = append(opts, grpc.Timeout(server.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	online.RegisterOnlineServer(srv, service)
	return srv
}
