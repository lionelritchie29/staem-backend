package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/helpers"
	chat_message "github.com/lionelritchie29/staem-backend/models/chat-message"
)

func CreateChat(p graphql.ResolveParams) (i interface{}, e error){
	senderId := p.Args["senderId"].(int)
	recipientId := p.Args["recipientId"].(int)
	message := p.Args["message"].(string)

	isSuccess := chat_message.Add(senderId, recipientId, message)
	helpers.Broadcast()
	return isSuccess, nil
}

func GetChats(p graphql.ResolveParams) (i interface{}, e error){
	senderId := p.Args["senderId"].(int)
	recipientId := p.Args["recipientId"].(int)

	chatMessages := chat_message.GetChats(senderId, recipientId)
	return chatMessages, nil
}

func GetLatestChat(p graphql.ResolveParams) (i interface{}, e error){
	senderId := p.Args["senderId"].(int)
	recipientId := p.Args["recipientId"].(int)

	chatMessage := chat_message.GetLatestChat(senderId, recipientId)
	return chatMessage, nil
}

