package resolver

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/models/game_sale"
)

func CreateSale(p graphql.ResolveParams) (i interface{}, e error){
	gameId := p.Args["gameId"].(int)
	discount := p.Args["discount"].(int)
	validTo := p.Args["validTo"].(string)
	sale := game_sale.Create(gameId, discount, validTo)
	return sale, nil
}

func DeleteSale(p graphql.ResolveParams) (i interface{}, e error){
	gameId := p.Args["gameId"].(int)
	status := game_sale.Delete(gameId)
	fmt.Println(status)
	return status, nil
}

func UpdateSale(p graphql.ResolveParams) (i interface{}, e error){
	gameId := p.Args["gameId"].(int)
	discount := p.Args["discount"].(int)
	validTo := p.Args["validTo"].(string)
	sale := game_sale.Update(gameId, discount, validTo)
	return sale, nil
}

