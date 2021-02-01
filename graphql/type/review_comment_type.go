package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/user"
)

var reviewCommentType *graphql.Object

func GetReviewCommentType() *graphql.Object {
	if reviewCommentType == nil {
		reviewCommentType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "reviewComment",
			Fields: graphql.Fields{
				"postId": {
					Type: graphql.Int,
				},
				"user": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.ReviewComment)
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

	return reviewCommentType
}
