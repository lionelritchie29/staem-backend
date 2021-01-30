package _type

import "github.com/graphql-go/graphql"

var marketTransactionType *graphql.Object

func GetMarketTransactionType() *graphql.Object {
	if marketTransactionType == nil {
		marketTransactionType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "marketTransactionType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"userId": &graphql.Field{
					Type: graphql.Int,
				},
				"gameItemId": &graphql.Field{
					Type: graphql.Int,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"quantity": &graphql.Field{
					Type: graphql.Int,
				},
				"createdAt": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return marketTransactionType
}