package _type

import "github.com/graphql-go/graphql"

var paymentMethodType *graphql.Object

func GetPaymentMethodType() *graphql.Object {
	if paymentMethodType == nil {
		paymentMethodType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "paymentMethod",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"hasCard": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		})
	}

	return paymentMethodType
}