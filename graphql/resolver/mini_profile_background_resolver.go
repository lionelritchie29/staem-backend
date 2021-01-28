package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/mini_profile"
)

func GetAllMiniProfileBackgrounds(p graphql.ResolveParams) (i interface{}, e error){
	bg := mini_profile.GetAll()
	return bg, nil
}

func GetMiniProfileBackgroundByUserId(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	bg := mini_profile.GetByUserId(userId)
	return bg, nil
}

func CreateMiniProfileTransaction(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	itemId := p.Args["itemId"].(int)
	isSuccess := mini_profile.AddTransaction(userId, itemId)
	return isSuccess, nil
}