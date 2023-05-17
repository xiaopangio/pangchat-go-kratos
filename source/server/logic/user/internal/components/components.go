package components

import (
	"github.com/google/wire"
	"user/internal/components/client"
	"user/internal/components/logger"
	"user/internal/components/mysql"
	"user/internal/components/oss"
	"user/internal/components/redis"
	"user/internal/components/registry"
	"user/internal/components/sms"
	"user/internal/components/uid"
)

var ProviderSet = wire.NewSet(
	logger.NewHelper,
	mysql.NewMysql,
	redis.NewRedisClient,
	registry.NewEtcdClient,
	registry.NewEtcdUserRegistry,
	registry.NewEtcdRelationshipRegistry,
	client.NewRelationshipClient,
	sms.NewSmsClient,
	oss.NewOSS,
	oss.NewSTS,
	uid.NewUidGenerator,
)
