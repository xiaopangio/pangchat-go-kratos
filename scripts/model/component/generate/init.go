// Package generate  @Author xiaobaiio 2023/3/17 18:17:00
package generate

import (
	"fmt"
	"gorm.io/gen"
	"model/component/mysql"
	"model/service/connector"
	"model/service/group"
	"model/service/logic"
	"model/service/message"
	"model/service/relationship"
	"model/service/user"
	"os"
	"strings"
)

var g *gen.Generator
var workdir string

func InitGen(dir string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	wd = strings.TrimSuffix(wd, "\\scripts\\model")
	dir = strings.ReplaceAll(dir, ".", "\\")
	wd = fmt.Sprintf("%s\\source\\server\\%s\\internal\\data\\orm\\", wd, dir)
	fmt.Printf("workdir: %s\n", wd)
	generator := gen.NewGenerator(gen.Config{
		Mode:         gen.WithDefaultQuery,
		OutPath:      wd + "dal",
		ModelPkgPath: wd + "model",
	})
	generator.UseDB(mysql.DB)
	g = generator
	workdir = dir
}
func Execute() {
	switch workdir {
	case "logic\\user":
		user.GenerateUser(g)
	case "group":
		group.GenerateGroup(g)
	case "connector":
		connector.GenerateConnector(g)
	case "logic\\relationship":
		relationship.GenerateRelationship(g)
	case "logic\\logic":
		logic.GenerateLogic(g)
	case "logic\\message":
		message.GenerateMessage(g)
	default:
		panic("unknown workdir")
	}
}
