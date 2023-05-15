// Package group  @Author xiaobaiio 2023/3/17 18:10:00
package group

import "gorm.io/gen"

func GenerateGroup(g *gen.Generator) {
	group := g.GenerateModel("group")
	groupRequest := g.GenerateModel("group_request")
	groupMember := g.GenerateModel("group_member")
	g.ApplyBasic(group, groupMember, groupRequest)
	g.Execute()
}
