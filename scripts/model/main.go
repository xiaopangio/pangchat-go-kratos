// Package main  @Author xiaobaiio 2023/3/17 17:19:00
package main

import (
	"flag"
	"model/component/generate"
	"model/component/mysql"
	"model/config"
)

var dir *string

func init() {
	dir = flag.String("service", "user", "The name of the service to generate a table (user,message,history-message,offline-message)")
}
func main() {
	flag.Parse()
	config.InitConfig()
	mysql.InitDB()
	generate.InitGen(*dir)
	generate.Execute()
}
