package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/user"
)

var postCommentType *graphql.Object

func GetPostCommentType() *graphql.Object {
	if postCommentType == nil {
		postCommentType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "postComment",
			Fields: graphql.Fields{
				"postId": {
					Type: graphql.Int,
				},
				"user": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.ImageVideoComment)
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

	return postCommentType
}
