package _type

import (
	"github.com/graphql-go/graphql"
)

var gameItemType *graphql.Object

func GetGameItemType() *graphql.Object {
	if gameItemType == nil {
		gameItemType = graphql.NewObject(graphql.ObjectConfig{
			Name: "gameItem",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"game": &graphql.Field{
					Type: graphql.Int,
				},
				"gameItemCategory": &graphql.Field{
					Type: graphql.Int,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"image": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return gameItemType
}