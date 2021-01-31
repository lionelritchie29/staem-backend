package _type

import "github.com/graphql-go/graphql"

var animatedAvatarFrameType *graphql.Object

func GetAnimatedAvatarFrameType() *graphql.Object {
	if animatedAvatarFrameType == nil {
		animatedAvatarFrameType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "animatedAvatar",
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

	return animatedAvatarFrameType
}