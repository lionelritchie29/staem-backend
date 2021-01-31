package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/user"
)

var recentMarketActivityType *graphql.Object

func GetMarketRecentActivityType() *graphql.Object{
	if(recentMarketActivityType == nil) {
		recentMarketActivityType = graphql.NewObject(graphql.ObjectConfig{
			Name: "marketRecentActivity",
			Fields: graphql.Fields{
				"type": &graphql.Field{
					Type: graphql.String,
				},
				"seller": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.RecentMarketActivity)
						userP := user.Get(parent.UserID)
						return userP, nil
					},
				},
				"buyer": &graphql.Field{
					Type: GetUserType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						parent := p.Source.(models.RecentMarketActivity)
						userP := user.Get(parent.BuyerID)
						return userP, nil
					},
				},
				"gameItemId": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"quantity": &graphql.Field{
					Type: graphql.Int,
				},
				"transactionDate": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		})
	}

	return recentMarketActivityType
}
