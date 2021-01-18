package _type

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/game_transaction"
)

var gameTransactionHeaderType *graphql.Object

func GetGameTransactionType() *graphql.Object {
	if gameTransactionHeaderType == nil {
		gameTransactionHeaderType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "gameTransactionHeader",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"user": &graphql.Field{
					Type: graphql.Int,
				},
				"cardNo": &graphql.Field{
					Type: graphql.String,
				},
				"cardExp": &graphql.Field{
					Type: graphql.String,
				},
				"billingAddress": &graphql.Field{
					Type: graphql.String,
				},
				"billingCity": &graphql.Field{
					Type: graphql.String,
				},
				"postalCode": &graphql.Field{
					Type: graphql.String,
				},
				"phoneNumber": &graphql.Field{
					Type: graphql.String,
				},
				"country": &graphql.Field{
					Type: graphql.String,
				},
				"transactionDate": &graphql.Field{
					Type: graphql.DateTime,
				},
				"paymentMethod": &graphql.Field{
					Type: GetPaymentMethodType(),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						tran := params.Source.(models.GameTransactionHeader)
						paymentMethod := game_transaction.GetPaymentMethodById(tran.PaymentMethod)
						return paymentMethod, nil
					},
				},
				"details": &graphql.Field{
					Type: graphql.NewList(GetGameTransactionDetailType()),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						tran := params.Source.(models.GameTransactionHeader)
						details := game_transaction.GetDetails(int(tran.ID))
						return details, nil
					},
				},
			},
		})
	}

	return gameTransactionHeaderType
}