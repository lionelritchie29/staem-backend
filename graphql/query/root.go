package query


import (
	"github.com/graphql-go/graphql"
	res "github.com/lionelritchie29/staem-backend/graphql/resolver"
	typ "github.com/lionelritchie29/staem-backend/graphql/type"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			// Publishers
			"publishers" : &graphql.Field{
				Type: graphql.NewList(typ.GetPublisherType()),
				Resolve: res.GetPublishers,
			},
			"publisher": &graphql.Field{
				Type: typ.GetPublisherType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetPublisher,
			},

			// Developers
			"developers" : &graphql.Field{
				Type: graphql.NewList(typ.GetPublisherType()),
				Resolve: res.GetDevelopers,
			},
			"developer": &graphql.Field{
				Type: typ.GetDeveloperType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetDeveloper,
			},

			// Games
			"games": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Resolve: res.GetGames,
			},
			"gamesByTitle": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"query": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetGamesByTitle,
			},
			"game": &graphql.Field{
				Type: typ.GetGameType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGame,
			},
			"playtimes": &graphql.Field{
				Type: graphql.NewList(typ.GetGamePlaytimeType()),
				Resolve: res.GetPlaytimes,
			},
			"playtimeByGameId": &graphql.Field{
				Type: graphql.NewList(typ.GetGamePlaytimeType()),
				Args: graphql.FieldConfigArgument{
					"id" : &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetPlaytimesByGameId,
			},
			"gameSales": &graphql.Field{
				Type: graphql.NewList(typ.GetGameSaleType()),
				Resolve: res.GetGameSales,
			},
			"gamesOnSale": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Resolve: res.GetSpecialOfferGames,
			},
			"featuredGames": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Resolve: res.GetFeaturedRecommendedGames,
			},
			"recentlyPublishedGames": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Resolve: res.GetRecentlyPublishedGames,
			},
			"specialGames": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Resolve: res.GetSpecialCategoryGames,
			},

			//Game Review
			"gameReviewById": &graphql.Field{
				Type: graphql.NewList(typ.GetGameReviewType()),
				Args: graphql.FieldConfigArgument{
					"gameId" : &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGameReviewByGameId,
			},

			// User
			"users": &graphql.Field{
				Type: graphql.NewList(typ.GetUserType()),
				Resolve: res.GetUsers,
			},
		},
	})
}