package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/game_item"
)

var marketSellListingOnlyItemIdAndQtyType *graphql.Object

func GetMarketSellListingOnlyItemIdAndQtyType() *graphql.Object {
	if marketSellListingOnlyItemIdAndQtyType == nil {
		marketSellListingOnlyItemIdAndQtyType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "marketSellListingOnlyItemIdAndQty",
			Fields: graphql.Fields{
				"gameItem": {
					Type: GetGameItemType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.MarketSellListingOnlyItemIdAndQty)
						gameItem := game_item.Get(int(parent.GameItemID))
						return gameItem, nil
					},
				},
				"quantity": &graphql.Field{
					Type: graphql.Int,
				},
				"lowestPrice": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}

	return marketSellListingOnlyItemIdAndQtyType
}
