package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/user"
)

var unsuspendRequestType *graphql.Object

func GetUnsuspendRequestType() *graphql.Object{
	if(unsuspendRequestType == nil) {
		unsuspendRequestType = graphql.NewObject(graphql.ObjectConfig{
			Name: "unsuspendRequest",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.UnsuspendRequest)
						userTarget := user.Get(int(parent.UserID))
						return userTarget, nil
					},
				},
				"reason": &graphql.Field{
					Type: graphql.String,
				},
				"status": &graphql.Field{
					Type: graphql.String,
				},
				"createdAt": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		})
	}

	return unsuspendRequestType
}


