package _type

import "github.com/graphql-go/graphql"

var marketSellListingType *graphql.Object

func GetMarketSellListingPaginateType() *graphql.Object {
	if marketSellListingType == nil {
		marketSellListingType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "marketSellListingPaginate",
			Fields: graphql.Fields{
				"totalItems": &graphql.Field{
					Type: graphql.Int,
				},
				"lowestPrice": &graphql.Field{
					Type: graphql.Int,
				},
				"listings": &graphql.Field{
					Type: graphql.NewList(GetMarketSellListingOnlyItemIdAndQtyType()),
				},
			},
		})
	}

	return marketSellListingType
}