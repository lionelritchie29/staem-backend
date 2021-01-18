package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/user"
)

var userType *graphql.Object

func GetUserType() *graphql.Object{
	if(userType == nil) {
		userType = graphql.NewObject(graphql.ObjectConfig{
			Name: "user",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"accountName": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"walletAmount": &graphql.Field{
					Type: graphql.Int,
				},
				"suspended": &graphql.Field{
					Type: graphql.DateTime,
				},
				"status": &graphql.Field{
					Type: graphql.String,
				},
				"role": &graphql.Field{
					Type: GetRoleType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						userAcc := params.Source.(models.UserAccount)
						role := user.GetRole(userAcc.Role)
						return role, nil
					},
				},
				"profile": &graphql.Field{
					Type: GetUserProfileType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						userAcc := params.Source.(models.UserAccount)
						profile := user.GetProfile(int(userAcc.ID))
						return profile, nil
					},
				},
			},
		})
	}

	return userType
}
