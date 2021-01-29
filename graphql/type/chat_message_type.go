package _type

import "github.com/graphql-go/graphql"

var chatMessageType *graphql.Object

func GetChatMessageType() *graphql.Object {
	if chatMessageType == nil {
		chatMessageType = graphql.NewObject(graphql.ObjectConfig{
			Name: "chatMessage",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"senderId": &graphql.Field{
					Type: graphql.Int,
				},
				"recipientId": &graphql.Field{
					Type: graphql.Int,
				},
				"message": &graphql.Field{
					Type: graphql.String,
				},
				"createdAt": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}

	return chatMessageType
}
