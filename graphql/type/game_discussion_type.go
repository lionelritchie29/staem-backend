package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/game_discussion"
	"github.com/lionelritchie29/staem-backend/models/user"
)

var gameDiscussionType *graphql.Object

func GetGameDiscussionType() *graphql.Object {
	if gameDiscussionType == nil {
		gameDiscussionType = graphql.NewObject(graphql.ObjectConfig{
			Name: "gameDiscussion",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
				"gameId": &graphql.Field{
					Type: graphql.Int,
				},
				"user": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.GameDiscussionPost)
						userP := user.Get(int(parent.UserID))
						return userP, nil
					},
				},
				"createdAt": &graphql.Field{
					Type: graphql.DateTime,
				},
				"comments": &graphql.Field{
					Type: graphql.NewList(GetDiscussionCommentType()),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.GameDiscussionPost)
						comments := game_discussion.GetCommentsByPostId(int(parent.ID))
						return comments, nil
					},
				},
			},
		})
	}

	return gameDiscussionType
}
