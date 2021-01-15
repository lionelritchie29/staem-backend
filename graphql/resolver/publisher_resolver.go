package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/publisher"
)


func GetPublishers(p graphql.ResolveParams) (i interface{}, e error){
	publishers := publisher.GetAll()
	return publishers, nil
}

func GetPublisher(p graphql.ResolveParams) (i interface{}, e error) {
	id := p.Args["id"].(int)
	publisher := publisher.Get(id)
	return publisher, nil
}