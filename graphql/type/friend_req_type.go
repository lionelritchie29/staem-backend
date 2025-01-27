package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/user"
)

var friendReqType *graphql.Object

func GetFriendRequestType() *graphql.Object {
	if friendReqType == nil {
		friendReqType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "friendRequest",
			Fields: graphql.Fields{
				//"user": &graphql.Field{
				//	Type: GetUserType(),
				//	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//		friendP := p.Source.(models.ReceivedFriendRequest)
				//		userP := user.Get(int(friendP.UserID))
				//		return userP, nil
				//	},
				//},
				"friend": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						friendP := p.Source.(models.ReceivedFriendRequest)
						userP := user.Get(int(friendP.FriendID))
						return userP, nil
					},
				},
				"status": &graphql.Field{
					Type: graphql.String,
				},
				"createdAt": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return friendReqType
}