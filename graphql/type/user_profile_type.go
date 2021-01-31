package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/animated_avatar"
	"github.com/lionelritchie29/staem-backend/models/avatar_frame"
	"github.com/lionelritchie29/staem-backend/models/chat_sticker"
	"github.com/lionelritchie29/staem-backend/models/mini_profile"
	"github.com/lionelritchie29/staem-backend/models/profile_background"
)

var profileType *graphql.Object

func GetUserProfileType() *graphql.Object {
	if profileType == nil {
		profileType = graphql.NewObject(graphql.ObjectConfig{
			Name: "profile",
			Fields: graphql.Fields{
				"displayName": &graphql.Field{
					Type: graphql.String,
				},
				"realName": &graphql.Field{
					Type: graphql.String,
				},
				"level": &graphql.Field{
					Type: graphql.Int,
				},
				"point": &graphql.Field{
					Type: graphql.Int,
				},
				"customURL": &graphql.Field{
					Type: graphql.String,
				},
				"avatarFrameUrl": &graphql.Field{
					Type: graphql.String,
				},
				"profilePictureUrl": &graphql.Field{
					Type: graphql.String,
				},
				"profileBackgroundUrl": &graphql.Field{
					Type: graphql.String,
				},
				"miniProfileBackgroundUrl": &graphql.Field{
					Type: graphql.String,
				},
				"summary": &graphql.Field{
					Type: graphql.String,
				},
				"theme": &graphql.Field{
					Type: graphql.String,
				},
				"featuredBadgeUrl": &graphql.Field{
					Type: graphql.String,
				},
				"country": &graphql.Field{
					Type: graphql.String,
				},
				"avatarFrames": &graphql.Field{
					Type: graphql.NewList(GetAvatarFrameType()),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.UserProfile)
						frames := avatar_frame.GetByUserId(int(parent.ID))
						return frames, nil
					},
				},
				"profileBackgrounds": &graphql.Field{
					Type: graphql.NewList(GetProfileBackgroundType()),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.UserProfile)
						bg := profile_background.GetByUserId(int(parent.ID))
						return bg, nil
					},
				},
				"miniProfileBackgrounds": &graphql.Field{
					Type: graphql.NewList(GetMiniProfileBackgroundType()),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.UserProfile)
						bg := mini_profile.GetByUserId(int(parent.ID))
						return bg, nil
					},
				},
				"animatedAvatars": &graphql.Field{
					Type: graphql.NewList(GetAnimatedAvatarFrameType()),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.UserProfile)
						avatars := animated_avatar.GetByUserId(int(parent.ID))
						return avatars, nil
					},
				},
				"chatStickers": &graphql.Field{
					Type: graphql.NewList(GetChatStickerType()),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.UserProfile)
						stickers := chat_sticker.GetByUserId(int(parent.ID))
						return stickers, nil
					},
				},
			},
		})
	}

	return profileType
}
