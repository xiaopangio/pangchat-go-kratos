package components

import (
	"api-gateway/internal/components/auth"
	"api-gateway/internal/components/client"
	"api-gateway/internal/components/logger"
	"api-gateway/internal/components/registry"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	logger.NewHelper,
	registry.NewEtcdClient,
	registry.NewEtcdUserRegistry,
	registry.NewEtcdConnectorRegistry,
	registry.NewEtcdLogicRegistry,
	registry.NewEtcdRelationshipRegistry,
	client.NewUserClient,
	client.NewConnectorClient,
	client.NewLogicClient,
	client.NewRelationshipClient,
	auth.NewJwtManager,
)
