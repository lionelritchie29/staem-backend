package _type

import "github.com/graphql-go/graphql"

var avatarFrameType *graphql.Object

func GetAvatarFrameType() *graphql.Object {
	if avatarFrameType == nil {
		avatarFrameType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "avatarFrame",
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

	return avatarFrameType
}