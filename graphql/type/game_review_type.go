package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/game_review"
	"github.com/lionelritchie29/staem-backend/models/user"
)

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
				"user": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						gameP := params.Source.(models.GameReview)
						user := user.Get(int(gameP.UserID))
						return user, nil
					},
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
				"comments": &graphql.Field{
					Type: graphql.NewList(GetReviewCommentType()),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.GameReview)
						comments := game_review.GetComments(int(parent.ID))
						return comments, nil
					},
				},
			},
		})
	}

	return gameReview
}
