package _type

import "github.com/graphql-go/graphql"

var gameGenre *graphql.Object

func GetGameGenreType() *graphql.Object {
	if gameGenre == nil {
		gameGenre = graphql.NewObject(graphql.ObjectConfig{
			Name: "genre",
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

	return gameGenre
}
