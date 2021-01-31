package subscription

import (
	"github.com/graphql-go/graphql"
	res "github.com/lionelritchie29/staem-backend/graphql/resolver"
	typ "github.com/lionelritchie29/staem-backend/graphql/type"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootSubscription",
		Fields: graphql.Fields{
			"latestChat": &graphql.Field{
				Type: typ.GetChatMessageType(),
				Args: graphql.FieldConfigArgument{
					"senderId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"recipientId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetLatestChat,
			},
			"marketRecentActivity": &graphql.Field{
				Type: graphql.NewList(typ.GetMarketRecentActivityType()),
				Args: graphql.FieldConfigArgument{
					"gameItemId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetMarketRecentActivites,
			},
		},
	})
}
