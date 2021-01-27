package input_type

import (
	"github.com/graphql-go/graphql"
)

var newProfileInfoType *graphql.InputObject

func GetNewProfileInputType() *graphql.InputObject {
	if newProfileInfoType == nil {
		newProfileInfoType = graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "NewProfileInfo",
				Fields: graphql.InputObjectConfigFieldMap{
					"userId": &graphql.InputObjectFieldConfig{
						Type: graphql.Int,
					},
					"displayName": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"realName": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"customUrl": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"country": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"summary": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
				},
			})
	}

	return newProfileInfoType
}