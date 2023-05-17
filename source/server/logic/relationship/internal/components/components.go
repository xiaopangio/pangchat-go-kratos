package components

import (
	"github.com/google/wire"
	"relationship/internal/components/broker"
	"relationship/internal/components/client"
	"relationship/internal/components/endpoints"
	"relationship/internal/components/logger"
	"relationship/internal/components/mysql"
	"relationship/internal/components/redis"
	"relationship/internal/components/registry"
	"relationship/internal/components/uid"
)

var ProviderSet = wire.NewSet(
	logger.NewHelper,
	redis.NewRedisClient,
	registry.NewEtcdClient,
	registry.NewEtcdRelationshipRegistry,
	registry.NewEtcdUserRegistry,
	client.NewUserClient,
	mysql.NewMysql,
	broker.NewKafkaBroker,
	uid.NewUidGenerator,
	endpoints.NewEndPoints,
)
