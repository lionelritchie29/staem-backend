package _type

import (
	"github.com/graphql-go/graphql"
)

var imageVideoPostType *graphql.Object

func GetImageVideoPostType() *graphql.Object {
	if imageVideoPostType == nil {
		imageVideoPostType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "imageVideoPost",
			Fields: graphql.Fields{
				"id": {
					Type: graphql.Int,
				},
				"fileUrl": &graphql.Field{
					Type: graphql.String,
				},
				"type": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"likeCount": {
					Type: graphql.Int,
				},
				"dislikeCount": {
					Type: graphql.Int,
				},
				"userId": {
					Type: graphql.Int,
				},
				"createdAt": {
					Type: graphql.String,
				},
			},
		})
	}

	return imageVideoPostType
}
