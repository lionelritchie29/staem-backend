package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/animated_avatar"
)

func GetAnimatedAvatars(p graphql.ResolveParams) (i interface{}, e error){
	avatars := animated_avatar.GetAll()
	return avatars, nil
}

func GetAnimatedAvatarByUserId(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	stickers := animated_avatar.GetByUserId(userId)
	return stickers, nil
}

func CreateAnimatedAvatarTransaction(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	itemId := p.Args["itemId"].(int)
	isSuccess := animated_avatar.AddTransaction(userId, itemId)
	return isSuccess, nil
}
