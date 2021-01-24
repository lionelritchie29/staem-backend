package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/developer"
	"github.com/lionelritchie29/staem-backend/models/game"
	"github.com/lionelritchie29/staem-backend/models/game_sale"
	"github.com/lionelritchie29/staem-backend/models/publisher"
)

var gameType *graphql.Object

func GetGameType() *graphql.Object {
	if gameType == nil {
		gameType = graphql.NewObject(graphql.ObjectConfig{
			Name: "game",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title" : &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"releasedate": &graphql.Field{
					Type: graphql.String,
				},
				"publisher": &graphql.Field{
					Type: GetPublisherType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						game := params.Source.(models.Game)
						pub := publisher.Get(game.Publisher)
						return pub, nil
					},
				},
				"developer": &graphql.Field{
					Type: GetDeveloperType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						game := params.Source.(models.Game)
						dev := developer.Get(game.Publisher)
						return dev, nil
					},
				},
				"genres": &graphql.Field{
					Type: graphql.NewList(GetGameGenreType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						gameP := params.Source.(models.Game)
						genres := game.GetGenres(int(gameP.ID))
						return genres, nil
					},
				},
				"tags": &graphql.Field{
					Type: graphql.NewList(GetGameTagType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						gameP := params.Source.(models.Game)
						tags := game.GetTags(int(gameP.ID))
						return tags, nil
					},
				},
				"images": &graphql.Field{
					Type: graphql.NewList(GetGameImageType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						gameP := params.Source.(models.Game)
						images := game.GetImages(int(gameP.ID))
						return images, nil
					},
				},
				"sale": &graphql.Field{
					Type: GetGameSaleType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						gameP := params.Source.(models.Game)
						sale := game_sale.GetByGameId(int(gameP.ID))
						return sale, nil
					},
				},
				"systemReq": &graphql.Field{
					Type: graphql.NewList(GetGameSystemReqType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						gameP := params.Source.(models.Game)
						systemReq := game.GetSystemRequirements(int(gameP.ID))
						return systemReq, nil
					},
				},
			},
		})
	}

	return gameType
}
