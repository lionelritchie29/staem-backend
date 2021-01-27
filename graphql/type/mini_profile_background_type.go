package _type

import "github.com/graphql-go/graphql"

var miniProfileBackgroundType *graphql.Object

func GetMiniProfileBackgroundType() *graphql.Object {
	if miniProfileBackgroundType == nil {
		miniProfileBackgroundType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "miniProfileBackground",
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

	return miniProfileBackgroundType
}