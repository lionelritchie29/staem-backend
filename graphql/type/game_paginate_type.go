package _type

import (
	"github.com/graphql-go/graphql"
)

var gamePaginateType *graphql.Object

func GetGamePaginateType() *graphql.Object {
	if gamePaginateType == nil {
		gamePaginateType = graphql.NewObject(graphql.ObjectConfig{
			Name: "gamesPaginate",
			Fields: graphql.Fields{
				"totalCount": &graphql.Field{
					Type: graphql.Int,
				},
				"games": &graphql.Field{
					Type: graphql.NewList(GetGameType()),
				},
			},
		})
	}

	return gamePaginateType
}
