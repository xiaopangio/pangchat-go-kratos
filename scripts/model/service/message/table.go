// Package message  @Author xiaobaiio 2023/3/17 18:10:00
package message

import "gorm.io/gen"

func GenerateMessage(g *gen.Generator) {
	message := g.GenerateModel("messages")
	g.ApplyBasic(message)
	g.Execute()
}
