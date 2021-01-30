package _type

import (
	"github.com/graphql-go/graphql"
)

var marketSellListingGroupedByPrice *graphql.Object

func GetMarketSellListingGroupedByPrice() *graphql.Object {
	if marketSellListingGroupedByPrice == nil {
		marketSellListingGroupedByPrice =  graphql.NewObject(graphql.ObjectConfig{
			Name: "marketSellListingGroupedByPrice",
			Fields: graphql.Fields{
				"price": {
					Type: graphql.Int,
				},
				"quantity": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return marketSellListingGroupedByPrice
}
