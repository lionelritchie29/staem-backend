package _type

import "github.com/graphql-go/graphql"

var walletVoucherType *graphql.Object

func GetWalletVoucherType() *graphql.Object {
	if walletVoucherType == nil {
		walletVoucherType =  graphql.NewObject(graphql.ObjectConfig{
			Name: "walletVoucher",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"code": &graphql.Field{
					Type: graphql.String,
				},
				"amount": &graphql.Field{
					Type: graphql.Int,
				},
				"isValid": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		})
	}

	return walletVoucherType
}