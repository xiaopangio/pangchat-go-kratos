package components

import (
	"connector/internal/components/broker"
	"connector/internal/components/cache"
	"connector/internal/components/client"
	"connector/internal/components/endpoints"
	"connector/internal/components/logger"
	"connector/internal/components/mysql"
	"connector/internal/components/redis"
	"connector/internal/components/registry"
	"connector/internal/components/websocket"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	logger.NewHelper,
	mysql.NewMysql,
	redis.NewRedisClient,
	registry.NewEtcdClient,
	registry.NewConnectorRegistry,
	registry.NewOnlineRegistry,
	registry.NewMessageRegistry,
	client.NewOnlineClient,
	client.NewMessageClient,
	cache.NewConnectionCache,
	websocket.NewUpgrader,
	broker.NewKafkaBroker,
	endpoints.NewEndPoints,
)
