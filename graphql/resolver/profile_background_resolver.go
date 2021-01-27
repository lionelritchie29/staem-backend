package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/profile_background"
)

func GetProfileBackgroundByUserId(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	bg := profile_background.GetByUserId(userId)
	return bg, nil
}

