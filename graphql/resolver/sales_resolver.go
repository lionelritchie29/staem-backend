package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/game_sale"
)

func GetGameSales(p graphql.ResolveParams) (i interface{}, e error){
	sales := game_sale.GetAll()
	return sales, nil
}

func GetSpecialGameSales (p graphql.ResolveParams) (i interface{}, e error){
	sales := game_sale.GetMoreThanFiftyPercentDiscount()
	return sales, nil
}