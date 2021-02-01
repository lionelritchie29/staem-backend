package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/game_discussion"
)

func GetGameDiscussions(p graphql.ResolveParams) (i interface{}, e error) {
	posts := game_discussion.GetAll()
	return posts, nil
}

func GetGameDiscussionById(p graphql.ResolveParams) (i interface{}, e error) {
	postId := p.Args["postId"].(int)
	post := game_discussion.Get(postId)
	return post, nil
}