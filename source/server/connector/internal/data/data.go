package data

import (
	"connector/internal/components/redis"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewConnectorRepoImpl)

// Data .
type Data struct {
	RedisCli *redis.Redis
	MysqlCli *gorm.DB
}

// NewData .
func NewData(RedisCli *redis.Redis, MysqlCli *gorm.DB, helper *log.Helper) (*Data, func(), error) {
	cleanup := func() {
		helper.Info("closing the data resources")
	}
	return &Data{
		RedisCli: RedisCli,
		MysqlCli: MysqlCli,
	}, cleanup, nil
}
