package _type

import "github.com/graphql-go/graphql"

var roleType *graphql.Object

func GetRoleType() *graphql.Object {
	if roleType == nil {
		roleType = graphql.NewObject(graphql.ObjectConfig{
			Name: "role",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return roleType
}
