package components

import (
	"github.com/google/wire"
	"logic/internal/components/loadbalance"
	"logic/internal/components/logger"
	"logic/internal/components/redis"
	"logic/internal/components/registry"
)

var ProviderSet = wire.NewSet(
	logger.NewHelper,
	redis.NewRedisClient,
	registry.NewEtcdClient,
	registry.NewEtcdLogicRegistry,
	registry.NewEtcdConnectorRegistry,
	loadbalance.NewRandomLoadBalance,
)
