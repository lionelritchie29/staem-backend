package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/mini_profile"
)

func GetMiniProfileBackgroundByUserId(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	bg := mini_profile.GetByUserId(userId)
	return bg, nil
}