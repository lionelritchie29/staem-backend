package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/game"
)

var gameTransactionDetailType *graphql.Object

func GetGameTransactionDetailType() *graphql.Object {
	if gameTransactionDetailType == nil {
		gameTransactionDetailType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "gameTransactionDetail",
			Fields: graphql.Fields{
				"game": &graphql.Field{
					Type: GetGameType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						detail := p.Source.(models.GameTransactionDetail)
						gam := game.Get(detail.Game)
						return gam, nil
					},
				},
				"quantity": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return gameTransactionDetailType
}