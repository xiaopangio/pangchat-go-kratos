package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user/internal/conf"
	"user/internal/data/orm/dal"
)

func NewMysql(c *conf.Bootstrap) *gorm.DB {
	cf := c.Data.Database
	db, err := gorm.Open(mysql.Open(cf.Source), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dal.SetDefault(db)
	return db
}
