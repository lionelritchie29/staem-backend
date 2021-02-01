package resolver

import (
	"github.com/graphql-go/graphql"
	image_video_post "github.com/lionelritchie29/staem-backend/models/image-video-post"
)

func GetImageVideoPosts(p graphql.ResolveParams) (interface{}, error) {
	posts := image_video_post.GetAll()
	return posts, nil
}
