package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/profile_background"
)

func GetAllProfileBackground(p graphql.ResolveParams) (i interface{}, e error){
	bg := profile_background.GetAll()
	return bg, nil
}

func GetProfileBackgroundByUserId(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	bg := profile_background.GetByUserId(userId)
	return bg, nil
}

func CreateProfileBgTransaction(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	itemId := p.Args["itemId"].(int)
	isSuccess := profile_background.AddTransaction(userId, itemId)
	return isSuccess, nil
}