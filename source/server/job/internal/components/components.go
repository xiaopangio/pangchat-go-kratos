package components

import (
	"github.com/google/wire"
	"job/internal/components/client"
	"job/internal/components/endpoints"
	"job/internal/components/logger"
	"job/internal/components/redis"
	"job/internal/components/registry"
)

var ProviderSet = wire.NewSet(
	logger.NewHelper,
	redis.NewRedisClient,
	registry.NewEtcdClient,
	registry.NewEtcdConnectorRegistry,
	registry.NewEtcdJobRegistry,
	registry.NewOnlineRegistry,
	registry.NewEtcdRelationshipRegistry,
	registry.NewEtcdUserRegistry,
	registry.NewMessageRegistry,
	client.NewConnectorClient,
	client.NewOnlineClient,
	client.NewRelationshipClient,
	client.NewUserClient,
	client.NewMessageClient,
	endpoints.NewEndPoints,
)
