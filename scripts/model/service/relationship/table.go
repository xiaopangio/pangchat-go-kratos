// Package relationship  @Author xiaobaiio 2023/3/17 18:10:00
package relationship

import (
	"gorm.io/gen"
)

func GenerateRelationship(g *gen.Generator) {
	friends := g.GenerateModel("friends")
	friendGroups := g.GenerateModel("friend_groups")
	friendRequests := g.GenerateModel("friend_requests")
	groups := g.GenerateModel("groups")
	groupMembers := g.GenerateModel("group_members")
	groupRequests := g.GenerateModel("group_requests")
	groupAdmins := g.GenerateModel("group_admins")
	g.ApplyBasic(groups, groupMembers, groupRequests, groupAdmins)
	g.ApplyBasic(friends, friendRequests, friendGroups)
	g.Execute()
}
