package _type

import "github.com/graphql-go/graphql"

var chatStickerType *graphql.Object

func GetChatStickerType() *graphql.Object {
	if chatStickerType == nil {
		chatStickerType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "chatSticker",
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

	return chatStickerType
}