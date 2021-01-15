package _type

import "github.com/graphql-go/graphql"

var gameSaleType *graphql.Object

func GetGameSaleType() *graphql.Object {
	if gameSaleType == nil {
		gameSaleType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "gamesale",
			Fields: graphql.Fields{
				"gameId": &graphql.Field{
					Type: graphql.Int,
				},
				"discount": &graphql.Field{
					Type: graphql.Int,
				},
				"validTo": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return gameSaleType
}