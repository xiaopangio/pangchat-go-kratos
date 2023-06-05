package data

import (
	"gorm.io/gorm"
	"logic/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewLogicRepoImpl)

// Data .
type Data struct {
	Mysql *gorm.DB
}

// NewData .
func NewData(cf *conf.Bootstrap, logger log.Logger, Mysql *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		Mysql: Mysql,
	}, cleanup, nil
}
