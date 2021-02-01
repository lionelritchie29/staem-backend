package _type

import (
	"github.com/graphql-go/graphql"
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
				"userId": &graphql.Field{
					Type: graphql.Int,
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
