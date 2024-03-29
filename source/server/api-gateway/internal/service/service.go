package service

import (
	service_connector "api-gateway/internal/service/connector"
	service_logic "api-gateway/internal/service/logic"
	service_message "api-gateway/internal/service/message"
	"api-gateway/internal/service/online"
	service_relationship "api-gateway/internal/service/relationship"
	"api-gateway/internal/service/user"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(service_user.NewUserService, service_connector.NewConnectorService, service_logic.NewLogicService, service_relationship.NewRelationship, service_online.NewOnline, service_message.NewMessage)
