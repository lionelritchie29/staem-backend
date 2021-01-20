package mutation

import (
	"github.com/graphql-go/graphql"
	res "github.com/lionelritchie29/staem-backend/graphql/resolver"

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
		},
	})
}
