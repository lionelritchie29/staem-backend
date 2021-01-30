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

			//tags and genres
			"tags": &graphql.Field{
				Type: graphql.NewList(typ.GetGameTagType()),
				Resolve: res.GetTags,
			},
			"genres": &graphql.Field{
				Type: graphql.NewList(typ.GetGameGenreType()),
				Resolve: res.GetGenre,
			},

			// Games
			"games": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Resolve: res.GetGames,
			},
			"gamesLimitOffset": &graphql.Field{
				Type: typ.GetGamePaginateType(),
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGamesLimitOffset,
			},
			"gamesByTitlelimitOffset": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"query": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetGamesByTitleLimitOffset,
			},
			"gamesByTitleLimit5": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"query": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetGamesByTitleLimit5,
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
			"communityRecommendedGames": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Resolve: res.GetCommunityRecommendedGames,
			},
			"topSellerGames": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Resolve: res.GetTopSellerGames,
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
			"mostHelpfulReviewByGameId": &graphql.Field{
				Type: graphql.NewList(typ.GetGameReviewType()),
				Args: graphql.FieldConfigArgument{
					"gameId" : &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetPastMonthGameReviewById,
			},
			"recentlyPostedReviewByGameId": &graphql.Field{
				Type: graphql.NewList(typ.GetGameReviewType()),
				Args: graphql.FieldConfigArgument{
					"gameId" : &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetRecentlyPostedReviewById,
			},

			// Transactions
			"gameTransactions": &graphql.Field{
				Type: graphql.NewList(typ.GetGameTransactionType()),
				Resolve: res.GetGameTransactions,
			},

			// User
			"users": &graphql.Field{
				Type: graphql.NewList(typ.GetUserType()),
				Resolve: res.GetUsers,
			},
			"user": &graphql.Field{
				Type: typ.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetUser,
			},
			"userByAccountName": &graphql.Field{
				Type: typ.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"url": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetUserByAccountName,
			},
			"userByCode": &graphql.Field{
				Type: typ.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetUserByCode,
			},
			"friendsByUserId": &graphql.Field{
				Type: graphql.NewList(typ.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetFriends,
			},
			"gamesByUserId": &graphql.Field{
				Type: graphql.NewList(typ.GetGameType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetUserGames,
			},
			"commentsByUserId": &graphql.Field{
				Type: graphql.NewList(typ.GetUserCommentType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetUserComments,
			},"friendRequestByUserId": &graphql.Field{
				Type: graphql.NewList(typ.GetFriendRequestType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetReceivedFriendRequest,
			},"sentFriendRequestById": &graphql.Field{
				Type: graphql.NewList(typ.GetSentFriendRequestType()),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetSentFriendRequest,
			},
			"reportsByUserId": &graphql.Field{
				Type: graphql.NewList(typ.GetUserReportType()),
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetReportsByUserId,
			},
			"unsuspendRequests": &graphql.Field{
				Type: graphql.NewList(typ.GetUnsuspendRequestType()),
				Resolve: res.GetAllSuspendedRequests,
			},

			//wishlist
			"wishlists": &graphql.Field{
				Type: graphql.NewList(typ.GetUserWishlistType()),
				Resolve: res.GetWishlists,
			},
			"wishlistByUserId": &graphql.Field{
				Type: graphql.NewList(typ.GetUserWishlistType()),
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetWishlistByUserId,
			},

			// Auth
			"login": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"accountName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.LoginResolver,
			},

			//
			"avatarFrameById": &graphql.Field{
				Type: graphql.NewList(typ.GetAvatarFrameType()),
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetAvatarFrameByUserId,
			},
			"profileBackgroundById": &graphql.Field{
				Type: graphql.NewList(typ.GetProfileBackgroundType()),
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetProfileBackgroundByUserId,
			},
			"miniProfileBackgroundById": &graphql.Field{
				Type: graphql.NewList(typ.GetMiniProfileBackgroundType()),
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetMiniProfileBackgroundByUserId,
			},
			"avatarFrames": &graphql.Field{
				Type: graphql.NewList(typ.GetAvatarFrameType()),
				Resolve: res.GetAvatarFrames,
			},
			"profileBackgrounds": &graphql.Field{
				Type: graphql.NewList(typ.GetProfileBackgroundType()),
				Resolve: res.GetAllProfileBackground,
			},
			"miniProfileBackgrounds": &graphql.Field{
				Type: graphql.NewList(typ.GetMiniProfileBackgroundType()),
				Resolve: res.GetAllMiniProfileBackgrounds,
			},

			"chats": &graphql.Field{
				Type: graphql.NewList(typ.GetChatMessageType()),
				Args: graphql.FieldConfigArgument{
					"senderId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"recipientId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetChats,
			},

			"sellListings": &graphql.Field{
				Type: graphql.NewList(typ.GetMarketTransactionType()),
				Resolve: res.GetSellListings,
			},

			"buyListings": &graphql.Field{
				Type: graphql.NewList(typ.GetMarketTransactionType()),
				Resolve: res.GetBuyListings,
			},

			"sellListingsPaginate": &graphql.Field{
				Type: typ.GetMarketSellListingPaginateType(),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetSellListingsPaginate,
			},
		},
	})
}