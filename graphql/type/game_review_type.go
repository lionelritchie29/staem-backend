package _type

import "github.com/graphql-go/graphql"

var gameReview *graphql.Object

func GetGameReviewType() *graphql.Object {
	if gameReview == nil {
		gameReview = graphql.NewObject(graphql.ObjectConfig{
			Name: "gameReview",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"isRecommended": &graphql.Field{
					Type: graphql.Boolean,
				},
				"gameId": &graphql.Field{
					Type: graphql.Int,
				},
				"userId": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
				"upvoteCount": &graphql.Field{
					Type: graphql.Int,
				},
				"downvoteCount": &graphql.Field{
					Type: graphql.Int,
				},
				"reviewDateTime": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		})
	}

	return gameReview
}
