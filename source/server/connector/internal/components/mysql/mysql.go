package mysql

import (
	"connector/internal/conf"
	"connector/internal/data/orm/dal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(c *conf.Bootstrap) *gorm.DB {

	db, err := gorm.Open(mysql.Open(c.Data.Database.Source), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dal.SetDefault(db)
	return db
}
