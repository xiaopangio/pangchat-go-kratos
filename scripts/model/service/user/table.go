// Package user  @Author xiaobaiio 2023/3/17 18:10:00
package user

import (
	"gorm.io/gen"
)

func GenerateUser(g *gen.Generator) {
	user := g.GenerateModel("users")
	province := g.GenerateModel("provinces")
	city := g.GenerateModel("cities")
	g.ApplyBasic(user, province, city)
	g.Execute()
}
