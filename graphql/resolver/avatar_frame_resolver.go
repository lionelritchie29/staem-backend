package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/avatar_frame"
)

func GetAvatarFrames(p graphql.ResolveParams) (i interface{}, e error){
	frames := avatar_frame.GetAll()
	return frames, nil
}

func GetAvatarFrameByUserId(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	frames := avatar_frame.GetByUserId(userId)
	return frames, nil
}

func CreateAvatarFrameTransaction(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	itemId := p.Args["itemId"].(int)
	isSuccess := avatar_frame.AddTransaction(userId, itemId)
	return isSuccess, nil
}