package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"relationship/internal/conf"
	"relationship/internal/data/orm/dal"
)

func NewMysql(cf *conf.Bootstrap) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cf.Data.Database.Source), &gorm.Config{})
	fmt.Println(cf.Data.Database.Source)
	if err != nil {
		panic(err)
	}
	dal.SetDefault(db)
	return db
}
