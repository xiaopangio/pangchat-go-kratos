// Package mysql  @Author xiaobaiio 2023/3/17 17:31:00
package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func InitDB() {
	root := viper.GetString("mysql.root")
	pass := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	dbName := viper.GetString("mysql.dbName")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", root, pass, host, dbName)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	DB = db
}
