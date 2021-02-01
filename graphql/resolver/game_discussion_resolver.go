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

func CreateGameDiscussionComment(p graphql.ResolveParams) (i interface{}, e error) {
	postId := p.Args["postId"].(int)
	userId := p.Args["userId"].(int)
	content := p.Args["content"].(string)
	post := game_discussion.CreateComment(postId, userId, content)
	return post, nil
}

func CreateNewGameDiscussion(p graphql.ResolveParams) (i interface{}, e error) {
	gameId := p.Args["gameId"].(int)
	userId := p.Args["userId"].(int)
	title := p.Args["title"].(string)
	content := p.Args["content"].(string)
	id := game_discussion.CreateDiscussion(userId, gameId, title, content)
	return id, nil
}