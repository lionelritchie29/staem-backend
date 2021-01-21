package mutation

import (
	"github.com/graphql-go/graphql"
	res "github.com/lionelritchie29/staem-backend/graphql/resolver"
	typ "github.com/lionelritchie29/staem-backend/graphql/type"
	inp_typ "github.com/lionelritchie29/staem-backend/graphql/input_type"
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

			"createNewUser": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"newUser": &graphql.ArgumentConfig{
						Type: inp_typ.GetUserInputType(),
					},
				},
				Resolve: res.CreateUser,
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
		},
	})
}