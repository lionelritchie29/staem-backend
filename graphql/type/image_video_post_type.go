package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	image_video_post "github.com/lionelritchie29/staem-backend/models/image-video-post"
	"github.com/lionelritchie29/staem-backend/models/user"
)

var imageVideoPostType *graphql.Object

func GetImageVideoPostType() *graphql.Object {
	if imageVideoPostType == nil {
		imageVideoPostType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "imageVideoPost",
			Fields: graphql.Fields{
				"id": {
					Type: graphql.Int,
				},
				"fileUrl": &graphql.Field{
					Type: graphql.String,
				},
				"type": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"likeCount": {
					Type: graphql.Int,
				},
				"dislikeCount": {
					Type: graphql.Int,
				},
				"user": {
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.ImageVideoPost)
						userP := user.Get(int(parent.UserID))
						return userP, nil
					},
				},
				"createdAt": {
					Type: graphql.String,
				},
				"comments": {
					Type: graphql.NewList(GetPostCommentType()),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.ImageVideoPost)
						comments := image_video_post.GetComments(int(parent.ID))
						return comments, nil
					},
				},
			},
		})
	}

	return imageVideoPostType
}
