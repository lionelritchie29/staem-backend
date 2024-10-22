package mutation

import (
	"github.com/graphql-go/graphql"
	inp_typ "github.com/lionelritchie29/staem-backend/graphql/input_type"
	res "github.com/lionelritchie29/staem-backend/graphql/resolver"
	typ "github.com/lionelritchie29/staem-backend/graphql/type"
)

func GetRoot() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createNewSelfTransaction": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"newTransaction": &graphql.ArgumentConfig{
						Type: inp_typ.GetGameTransactionInputType(),
					},
				},
				Resolve: res.AddTransaction,
			},

			"createFriendRequest": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"fromId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"toId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CreateFriendRequest,
			},

			"createGame": &graphql.Field{
				Type: typ.GetGameType(),
				Args: graphql.FieldConfigArgument{
					"newGame": &graphql.ArgumentConfig{
						Type: inp_typ.GetGameInputType(),
					},
				},
				Resolve: res.CreateGame,
			},
			"deleteGame": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteGame,
			},

			"createPromo": &graphql.Field{
				Type: typ.GetGameSaleType(),
				Args: graphql.FieldConfigArgument{
					"gameId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"discount": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"validTo": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateSale,
			},
			"deletePromo": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"gameId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteSale,
			},
			"updatePromo": &graphql.Field{
				Type: typ.GetGameSaleType(),
				Args: graphql.FieldConfigArgument{
					"gameId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"discount": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"validTo": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateSale,
			},


			"acceptFriendRequest": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"fromId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"toId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.AcceptFriendRequest,
			},

			"rejectFriendRequest": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"fromId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"toId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.RejectFriendRequest,
			},

			"ignoreFriendRequest": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"fromId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"toId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.IgnoreFriendRequest,
			},

			"createNewUser": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"newUser": &graphql.ArgumentConfig{
						Type: inp_typ.GetUserInputType(),
					},
				},
				Resolve: res.CreateUser,
			},
			"suspendUser": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.SuspendUser,
			},
			"unsuspendUser": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.UnsuspendUser,
			},
			"createUnsuspendRequest": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"reason": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateUnsuspendRequest,
			},
			"denyUnsuspendRequest": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DenyUnsuspendRequest,
			},

			"createNewComment": &graphql.Field{
				Type: typ.GetUserCommentType(),
				Args: graphql.FieldConfigArgument{
					"srcUserId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"destUserId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateComment,
			},

			"createWishlist": &graphql.Field{
				Type: typ.GetUserWishlistType(),
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"gameId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CreateWishlist,
			},
			"deleteWishlist": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"gameId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.DeleteWishlist,
			},

			"redeemWallet": &graphql.Field{
				Type: typ.GetWalletVoucherType(),
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.RedeemWallet,
			},

			"logout": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.LogoutResolver,
			},

			"createReport": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"reporterId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"reason": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateReport,
			},

			"createGameReview": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"gameId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"isRecommended": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
				},
				Resolve: res.CreateGameReview,
			},
			"increseReviewUpvote": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"reviewId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.IncreaseReviewUpvoteCount,
			},
			"increseReviewDownvote": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"reviewId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.IncreaseReviewDownvoteCount,
			},

			// Profile
			"updateProfileInfo": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"newProfile": &graphql.ArgumentConfig{
						Type: inp_typ.GetNewProfileInputType(),
					},
				},
				Resolve: res.UpdateProfileInfo,
			},
			"updateProfileAvatar": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"imageString": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateProfileAvatar,
			},
			"updateProfileAvatarFrame": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"avatarFrameUrl": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateProfilAvatarFrame,
			},
			"updateProfileBackground": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"profileBackgroundUrl": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateProfileBackground,
			},
			"updateMiniProfileBackground": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"miniProfileBackgroundUrl": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateMiniProfileBackground,
			},
			"updateProfileTheme": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"hexCode": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateProfileTheme,
			},
			"createAvatarFrameTransaction": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"itemId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CreateAvatarFrameTransaction,
			},
			"createProfileBgTransaction": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"itemId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CreateProfileBgTransaction,
			},
			"createMiniProfileTransaction": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"itemId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CreateMiniProfileTransaction,
			},
			"createChatStickerTransaction": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"itemId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CreateChatStickerTransaction,
			},
			"createAnimatedAvatarTransaction": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"itemId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CreateAnimatedAvatarTransaction,
			},

			"createOTP": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"accountName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateOTP,
			},
			"verifyOTP": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.VerifyOTP,
			},
			"addChatMessage": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"senderId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"recipientId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"message": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateChat,
			},
			"createSellListing": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"gameItemId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"quantity": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CreateSellListing,
			},
			"createBuyListing": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"gameItemId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"quantity": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CreateBuyListing,
			},

			"createPostComment": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"postId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"comment": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.AddImageVideoComment,
			},
			"addImageVideoLike": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"postId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.AddImageVideoLike,
			},
			"addImageVideoDislike": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"postId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.AddImageVideoDislike,
			},
			"addReviewComment": &graphql.Field{
				Type: typ.GetReviewCommentType(),
				Args: graphql.FieldConfigArgument{
					"postId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"comment": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateReviewComment,
			},
			"createGameDiscussionComment": &graphql.Field{
				Type: typ.GetDiscussionCommentType(),
				Args: graphql.FieldConfigArgument{
					"postId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateGameDiscussionComment,
			},
			"createNewGameDiscussion": &graphql.Field{
				Type: graphql.Int,
				Args: graphql.FieldConfigArgument{
					"gameId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateNewGameDiscussion,
			},
		},
	})
}
