package components

import (
	"github.com/google/wire"
	"online/internal/components/logger"
	"online/internal/components/redis"
	"online/internal/components/registry"
)

var ProviderSet = wire.NewSet(
	logger.NewHelper,
	redis.NewRedisClient,
	registry.NewEtcdClient,
	registry.NewOnlineRegistry,
)
