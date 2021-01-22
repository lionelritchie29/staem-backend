package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/comments"
)

func CreateComment(p graphql.ResolveParams) (i interface{}, e error){
	srcUserId := p.Args["srcUserId"].(int)
	destUserId := p.Args["destUserId"].(int)
	content := p.Args["content"].(string)
	comment := comments.Create(srcUserId, destUserId, content)
	return comment, nil
}
