package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/user"
)

var userCommentType *graphql.Object

func GetUserCommentType() *graphql.Object{
	if(userCommentType == nil) {
		userCommentType = graphql.NewObject(graphql.ObjectConfig{
			Name: "userComment",
			Fields: graphql.Fields{
				"srcUser": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						comment := p.Source.(models.UserComment)
						userTarget := user.Get(int(comment.SrcUserId))
						return userTarget, nil
					},
				},
				"destUser": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						comment := p.Source.(models.UserComment)
						userTarget := user.Get(int(comment.DestUserId))
						return userTarget, nil
					},
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
				"createdAt": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		})
	}

	return userCommentType
}
