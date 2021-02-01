package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/input_models"
	"github.com/lionelritchie29/staem-backend/models/game_transaction"
	"github.com/mitchellh/mapstructure"
)

func GetGameTransactions(p graphql.ResolveParams) (interface{}, error) {
	transactions := game_transaction.GetAll()
	return transactions, nil
}

func AddTransaction(p graphql.ResolveParams) (interface{}, error) {
	newTransactionRaw := p.Args["newTransaction"]
	var newTransaction input_models.NewGameTransaction
	mapstructure.Decode(newTransactionRaw, &newTransaction)
	
	if !game_transaction.AddTransaction(newTransaction) {
		return false, nil
	}

	return true, nil
}