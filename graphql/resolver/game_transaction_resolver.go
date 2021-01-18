package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/game_transaction"
)

func GetGameTransactions(p graphql.ResolveParams) (interface{}, error) {
	transactions := game_transaction.GetAll()
	return transactions, nil
}
