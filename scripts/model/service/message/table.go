// Package message  @Author xiaobaiio 2023/3/17 18:10:00
package message

import "gorm.io/gen"

func GenerateHistoryMessage(g *gen.Generator) {
	message := g.GenerateModel("message")
	groupMessage := g.GenerateModel("group_message")
	g.ApplyBasic(message, groupMessage)
	g.Execute()
}
