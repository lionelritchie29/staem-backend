package _type

import "github.com/graphql-go/graphql"

var gamePlaytime *graphql.Object

func GetGamePlaytimeType() *graphql.Object {
	if gamePlaytime == nil {
		gamePlaytime = graphql.NewObject(graphql.ObjectConfig{
			Name: "playtime",
			Fields: graphql.Fields{
				"gameId": &graphql.Field{
					Type: graphql.Int,
				},
				"playHour": &graphql.Field{
					Type: graphql.Int,
				},
				"Date": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return gamePlaytime
}