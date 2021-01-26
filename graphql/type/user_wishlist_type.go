package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/game"
)

var userWishlistType *graphql.Object

func GetUserWishlistType() *graphql.Object {
	if userWishlistType == nil {
		userWishlistType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "userWishlist",
			Fields: graphql.Fields{
				"userId": &graphql.Field{
					Type: graphql.Int,
				},
				"game": &graphql.Field{
					Type: GetGameType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.UserWishlist)
						gameTemp := game.Get(int(parent.GameID))
						return gameTemp, nil
					},
				},
				"createdAt": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		})
	}

	return userWishlistType
}