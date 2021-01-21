package input_type

import (
	"github.com/graphql-go/graphql"
)

var userInputType *graphql.InputObject

func GetUserInputType() *graphql.InputObject {
	if userInputType == nil {
		userInputType = graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "NewUser",
				Fields: graphql.InputObjectConfigFieldMap{
					"accountName": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"email": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"password": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"country": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
				},
			})
	}

	return userInputType
}