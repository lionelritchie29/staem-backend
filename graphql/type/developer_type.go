package _type

import "github.com/graphql-go/graphql"

var developerType *graphql.Object

func GetDeveloperType() *graphql.Object {
	if developerType == nil {
		developerType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "developer",
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

	return developerType
}