package _type

import "github.com/graphql-go/graphql"

var gameTag *graphql.Object

func GetGameTagType() *graphql.Object {
	if gameTag == nil {
		gameTag = graphql.NewObject(graphql.ObjectConfig{
			Name: "tag",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"isAdult": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		})
	}

	return gameTag
}