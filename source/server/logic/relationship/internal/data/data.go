package data

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewRelationshipRepoImpl)

// Data .
type Data struct {
	Mysql *gorm.DB
}

// NewData .
func NewData(Mysql *gorm.DB) *Data {
	return &Data{
		Mysql: Mysql,
	}
}
