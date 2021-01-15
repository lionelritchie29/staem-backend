package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/developer"
)

func GetDevelopers(p graphql.ResolveParams) (i interface{}, e error){
	devs := developer.GetAll()
	return devs, nil
}

func GetDeveloper(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	dev := developer.Get(id)
	return dev, nil
}