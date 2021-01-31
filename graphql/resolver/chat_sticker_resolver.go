package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/chat_sticker"
)

func GetChatStickers(p graphql.ResolveParams) (i interface{}, e error){
	stickers := chat_sticker.GetAll()
	return stickers, nil
}

func GetChatStickerByUserId(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	stickers := chat_sticker.GetByUserId(userId)
	return stickers, nil
}

func CreateChatStickerTransaction(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	itemId := p.Args["itemId"].(int)
	isSuccess := chat_sticker.AddTransaction(userId, itemId)
	return isSuccess, nil
}