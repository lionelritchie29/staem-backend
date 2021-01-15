package _type

import "github.com/graphql-go/graphql"

var gameCategory *graphql.Object

func GetGameCategoryType() *graphql.Object {
	if gameCategory == nil {
		gameCategory = graphql.NewObject(graphql.ObjectConfig{
			Name: "category",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return gameCategory
}
