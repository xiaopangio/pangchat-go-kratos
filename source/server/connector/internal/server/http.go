package server

import (
	"connector/internal/conf"
	"connector/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cf *conf.Bootstrap, connectorService *service.ConnectorServiceService) *http.Server {
	g := gin.Default()
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	server := cf.Server.Http
	service.RegisterConnectorService(g, connectorService)
	if server.Network != "" {
		opts = append(opts, http.Network(server.Network))
	}
	if server.Addr != "" {
		opts = append(opts, http.Address(server.Addr))
	}
	if server.Timeout != nil {
		opts = append(opts, http.Timeout(server.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", g)
	return srv
}
