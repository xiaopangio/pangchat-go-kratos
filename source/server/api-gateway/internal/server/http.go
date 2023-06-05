package server

import (
	"api-gateway/internal/conf"
	"api-gateway/internal/service/connector"
	service_logic "api-gateway/internal/service/logic"
	service_message "api-gateway/internal/service/message"
	serivice_online "api-gateway/internal/service/online"
	service_relationship "api-gateway/internal/service/relationship"
	"api-gateway/internal/service/user"
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, messageService *service_message.Message, userService *service_user.UserService, connectorService *service_connector.ConnectorService, logicService *service_logic.LogicService, relationship *service_relationship.Relationship, onlineService *serivice_online.Online) *http.Server {
	g := gin.Default()
	g.Use(kgin.Middlewares(recovery.Recovery()), middleware.Cors())
	v1 := g.Group("/api/v1")
	service_user.RegisterUserService(v1, userService)                  //注册用户
	service_connector.RegisterConnectorService(v1, connectorService)   //注册连接器
	service_logic.RegisterLogicService(v1, logicService)               //注册逻辑服务
	service_relationship.RegisterRelationshipService(v1, relationship) //注册关系服务
	serivice_online.RegisterOnlineService(v1, onlineService)           //注册在线服务
	service_message.RegisterMessageService(v1, messageService)         //注册消息服务
	var opts []http.ServerOption
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", g)
	return srv
}
