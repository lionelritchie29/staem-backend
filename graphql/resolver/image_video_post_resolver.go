package resolver

import (
	"github.com/graphql-go/graphql"
	image_video_post "github.com/lionelritchie29/staem-backend/models/image-video-post"
)

func GetImageVideoPosts(p graphql.ResolveParams) (interface{}, error) {
	posts := image_video_post.GetAll()
	return posts, nil
}

func GetImageVideoPostById(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(int)
	post := image_video_post.GetById(id)
	return post, nil
}

func AddImageVideoComment(p graphql.ResolveParams) (interface{}, error) {
	userId := p.Args["userId"].(int)
	postId := p.Args["postId"].(int)
	comment := p.Args["comment"].(string)
	isSuccess := image_video_post.AddComments(postId, userId, comment)
	return isSuccess, nil
}

func AddImageVideoLike(p graphql.ResolveParams) (interface{}, error) {
	postId := p.Args["postId"].(int)
	isSuccess := image_video_post.AddLike(postId)
	return isSuccess, nil
}

func AddImageVideoDislike(p graphql.ResolveParams) (interface{}, error) {
	postId := p.Args["postId"].(int)
	isSuccess := image_video_post.AddDislike(postId)
	return isSuccess, nil
}