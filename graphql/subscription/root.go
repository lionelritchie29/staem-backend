package subscription

import (
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootSubscription",
		Fields: graphql.Fields{
			"message": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					return "HAI", nil
				},
			},

		},
	})
}
