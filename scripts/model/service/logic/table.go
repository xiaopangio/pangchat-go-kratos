// Package logic  @Author xiaobaiio 2023/3/17 18:10:00
package logic

import (
	"gorm.io/gen"
)

func GenerateLogic(g *gen.Generator) {
	toolOptions := g.GenerateModel("tool_options")
	emojis := g.GenerateModel("emojis")
	g.ApplyBasic(toolOptions, emojis)
	g.Execute()
}
