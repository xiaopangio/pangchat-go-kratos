// Package relationship  @Author xiaobaiio 2023/3/17 18:10:00
package relationship

import (
	"gorm.io/gen"
)

func GenerateRelationship(g *gen.Generator) {
	friends := g.GenerateModel("friends")
	friendGroups := g.GenerateModel("friend_groups")
	friendRequests := g.GenerateModel("friend_requests")
	g.ApplyBasic(friends, friendRequests, friendGroups)
	g.Execute()
}
