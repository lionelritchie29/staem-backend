package _type

import (
	"github.com/graphql-go/graphql"
)

var userPaginateType *graphql.Object

func GetUserPaginateType() *graphql.Object {
	if userPaginateType == nil {
		userPaginateType = graphql.NewObject(graphql.ObjectConfig{
			Name: "userPaginate",
			Fields: graphql.Fields{
				"totalCount": &graphql.Field{
					Type: graphql.Int,
				},
				"users": &graphql.Field{
					Type: graphql.NewList(GetUserType()),
				},
			},
		})
	}

	return userPaginateType
}
