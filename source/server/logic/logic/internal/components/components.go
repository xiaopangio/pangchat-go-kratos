package components

import (
	"github.com/google/wire"
	"logic/internal/components/endpoints"
	"logic/internal/components/loadbalance"
	"logic/internal/components/logger"
	"logic/internal/components/mysql"
	"logic/internal/components/oss"
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
	endpoints.NewEndPoints,
	mysql.NewMysql,
	oss.NewSTS,
	oss.NewOSS,
)
