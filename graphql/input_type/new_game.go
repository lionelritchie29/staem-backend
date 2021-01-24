package input_type

import (
	"github.com/graphql-go/graphql"
)

var gameInputType *graphql.InputObject

func GetGameInputType() *graphql.InputObject {
	if gameInputType == nil {
		gameInputType = graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name: "NewGame",
				Fields: graphql.InputObjectConfigFieldMap{
					"title": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"description": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"releaseDate": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"price": &graphql.InputObjectFieldConfig{
						Type: graphql.Int,
					},
					"tagIds": &graphql.InputObjectFieldConfig{
						Type: graphql.NewList(graphql.Int),
					},
					"genreIds": &graphql.InputObjectFieldConfig{
						Type: graphql.NewList(graphql.Int),
					},
					"publisherId": &graphql.InputObjectFieldConfig{
						Type: graphql.Int,
					},
					"developerId": &graphql.InputObjectFieldConfig{
						Type: graphql.Int,
					},
					"gameHeaderImage": &graphql.InputObjectFieldConfig{
						Type: graphql.String,
					},
					"gameImages": &graphql.InputObjectFieldConfig{
						Type: graphql.NewList(graphql.String),
					},
				},
			})
	}

	return gameInputType
}