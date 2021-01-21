package _type

import "github.com/graphql-go/graphql"

var profileType *graphql.Object

func GetUserProfileType() *graphql.Object {
	if profileType == nil {
		profileType = graphql.NewObject(graphql.ObjectConfig{
			Name: "profile",
			Fields: graphql.Fields{
				"displayName": &graphql.Field{
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
			},
		})
	}

	return profileType
}
