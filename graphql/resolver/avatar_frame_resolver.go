package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/avatar_frame"
)

func GetAvatarFrameByUserId(p graphql.ResolveParams) (i interface{}, e error){
	userId := p.Args["userId"].(int)
	frames := avatar_frame.GetByUserId(userId)
	return frames, nil
}
