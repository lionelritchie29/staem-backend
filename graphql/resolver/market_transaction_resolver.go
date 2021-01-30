package resolver

import (
	"github.com/graphql-go/graphql"
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