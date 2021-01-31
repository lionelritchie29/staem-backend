package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/lionelritchie29/staem-backend/helpers"
	"github.com/lionelritchie29/staem-backend/models/market_transaction"
)

func GetSellListings(p graphql.ResolveParams) (interface{}, error) {
	listings := market_transaction.GetAllSellListing()
	return listings, nil
}

func GetBuyListings(p graphql.ResolveParams) (interface{}, error) {
	listings := market_transaction.GetAllBuyListing()
	return listings, nil
}

func GetSellListingsPaginate(p graphql.ResolveParams) (interface{}, error) {
	offset := p.Args["offset"].(int)
	limit := p.Args["limit"].(int)

	listings := market_transaction.GetPaginateSellListing(limit, offset)
	return listings, nil
}

func GetSellListingsGroupedByPrice(p graphql.ResolveParams) (interface{}, error) {
	gameItemId := p.Args["gameItemId"].(int)

	listings := market_transaction.GetSellListingsGroupedByPrice(gameItemId)
	return listings, nil
}

func GetBuyListingsGroupedByPrice(p graphql.ResolveParams) (interface{}, error) {
	gameItemId := p.Args["gameItemId"].(int)

	listings := market_transaction.GetBuyListingsGroupedByPrice(gameItemId)
	return listings, nil
}

func GetSellListingByUserIdAndGameId(p graphql.ResolveParams) (interface{}, error) {
	userId := p.Args["userId"].(int)
	gameItemId := p.Args["gameItemId"].(int)

	listings := market_transaction.GetSellListingByUserIdAndGameItemId(userId, gameItemId)
	return listings, nil
}

func GetBuyListingByUserIdAndGameId(p graphql.ResolveParams) (interface{}, error) {
	userId := p.Args["userId"].(int)
	gameItemId := p.Args["gameItemId"].(int)

	listings := market_transaction.GetBuyListingByUserIdAndGameItemId(userId, gameItemId)
	return listings, nil
}

func CreateSellListing(p graphql.ResolveParams) (interface{}, error) {
	userId := p.Args["userId"].(int)
	gameItemId := p.Args["gameItemId"].(int)
	price := p.Args["price"].(int)
	quantity := p.Args["quantity"].(int)
	isSuccess := market_transaction.CreateSellListing(userId, gameItemId, price, quantity)

	helpers.Broadcast()
	return isSuccess, nil
}

func CreateBuyListing(p graphql.ResolveParams) (interface{}, error) {
	userId := p.Args["userId"].(int)
	gameItemId := p.Args["gameItemId"].(int)
	price := p.Args["price"].(int)
	quantity := p.Args["quantity"].(int)

	isSuccess := market_transaction.CreateBuyListing(userId, gameItemId, price, quantity)
	helpers.Broadcast()
	return isSuccess, nil
}

func GetMarketRecentActivites(p graphql.ResolveParams) (interface{}, error) {
	gameId := p.Args["gameItemId"].(int)
	activites := market_transaction.GetLatestRecentActvities(gameId)
	return activites, nil
}