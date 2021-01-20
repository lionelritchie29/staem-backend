package input_type

import (
	"github.com/graphql-go/graphql"
)

var gameDetailTransactionInputType *graphql.InputObject

func GetGameTransactionDetailInputType() *graphql.InputObject {
	if gameDetailTransactionInputType == nil {
		gameDetailTransactionInputType = graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "NewTransactionDetail",
				Fields: graphql.InputObjectConfigFieldMap{
					"game": &graphql.InputObjectFieldConfig{
						Type: graphql.Int,
					},
					"price": &graphql.InputObjectFieldConfig{
						Type: graphql.Int,
					},
					"quantity": &graphql.InputObjectFieldConfig{
						Type: graphql.Int,
					},
				},
			})
	}

	return gameDetailTransactionInputType
}