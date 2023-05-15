package server

import (
	"connector/api/v1/connector"
	"connector/internal/conf"
	"connector/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, connectorService *service.ConnectorServiceService, helper *log.Helper) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
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
	connector.RegisterConnectorServiceServer(srv, connectorService)
	return srv
}
