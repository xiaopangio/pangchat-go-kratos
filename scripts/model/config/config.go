// Package config  @Author xiaobaiio 2023/3/17 17:21:00
package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(fmt.Sprintf("%s\\config", workdir))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
