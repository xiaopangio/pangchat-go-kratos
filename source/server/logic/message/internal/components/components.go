package components

import (
	"github.com/google/wire"
	"message/internal/components/broker"
	"message/internal/components/endpoints"
	"message/internal/components/logger"
	"message/internal/components/mysql"
	"message/internal/components/redis"
	"message/internal/components/registry"
	"message/internal/components/uid"
)

var ProviderSet = wire.NewSet(
	logger.NewHelper,
	redis.NewRedisClient,
	registry.NewEtcdClient,
	registry.NewEtcdMessageRegistry,
	mysql.NewMysql,
	broker.NewKafkaBroker,
	uid.NewUidGenerator,
	endpoints.NewEndPoints,
)
