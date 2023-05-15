package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
	"user/internal/components/redis"
	"user/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	RedisCli *redis.Redis
	MysqlCli *gorm.DB
}

// NewData .
func NewData(c *conf.Bootstrap, redisCli *redis.Redis, db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		RedisCli: redisCli,
		MysqlCli: db,
	}, cleanup, nil
}
