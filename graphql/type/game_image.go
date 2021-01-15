package _type

import "github.com/graphql-go/graphql"

var gameImage *graphql.Object

func GetGameImageType() *graphql.Object {
	if gameImage == nil {
		gameImage = graphql.NewObject(graphql.ObjectConfig{
			Name: "image",
			Fields: graphql.Fields{
				"url": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return gameImage
}