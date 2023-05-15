package connector

import (
	"gorm.io/gen"
)

func GenerateConnector(g *gen.Generator) {
	user := g.GenerateModel("users")
	province := g.GenerateModel("provinces")
	city := g.GenerateModel("cities")
	g.ApplyBasic(user, province, city)
	g.Execute()
}
