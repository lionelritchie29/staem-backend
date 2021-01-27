package _type

import "github.com/graphql-go/graphql"

var profileBackgroundType *graphql.Object

func GetProfileBackgroundType() *graphql.Object {
	if profileBackgroundType == nil {
		profileBackgroundType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "profileBackground",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"url": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"createdAt": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		})
	}

	return profileBackgroundType
}