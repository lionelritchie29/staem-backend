package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/user"
)

var discussionCommentType *graphql.Object

func GetDiscussionCommentType() *graphql.Object {
	if discussionCommentType == nil {
		discussionCommentType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "discussionComment",
			Fields: graphql.Fields{
				"postId": {
					Type: graphql.Int,
				},
				"user": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.GameDiscussionComment)
						userP := user.Get(int(parent.UserID))
						return userP, nil
					},
				},
				"comments": &graphql.Field{
					Type: graphql.String,
				},
				"createdAt": {
					Type: graphql.String,
				},
			},
		})
	}

	return discussionCommentType
}
