package input_type

import (
	"github.com/graphql-go/graphql"
)

var gameTransactionInputType *graphql.InputObject

func GetGameTransactionInputType() *graphql.InputObject {
	if gameTransactionInputType == nil {
		gameTransactionInputType = graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "NewTransaction",
				Fields: graphql.InputObjectConfigFieldMap{
					"user": &graphql.InputObjectFieldConfig{
						Type: graphql.Int,
					},
					"cardNo": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"cardExp": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"billingAddress": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"billingCity": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"postalCode": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"phoneNumber": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"country": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"paymentMethod": &graphql.InputObjectFieldConfig{
						Type: graphql.Int,
					},
					"details": &graphql.InputObjectFieldConfig{
						Type: graphql.NewList(GetGameTransactionDetailInputType()),
					},
				},
			})
	}

	return gameTransactionInputType
}