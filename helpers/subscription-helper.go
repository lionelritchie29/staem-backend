package helpers

import (
	"github.com/functionalfoundry/graphqlws"
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/globals"
)

func Broadcast() {
	for conn, _ := range globals.SubscriptionManager.Subscriptions() {
		for _, sub := range globals.SubscriptionManager.Subscriptions()[conn] {
			params := graphql.Params{
				Schema:         globals.Schema,
				RequestString:  sub.Query,
				VariableValues: sub.Variables,
				OperationName:  sub.OperationName,
			}

			result := graphql.Do(params)

			data := graphqlws.DataMessagePayload{
				Data: result.Data,
				Errors: graphqlws.ErrorsFromGraphQLErrors(
					result.Errors),
			}

			sub.SendData(&data)
		}
	}
}
