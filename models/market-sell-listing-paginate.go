package models

type MarketSellListingPaginate struct {
	TotalItems int
	LowestPrice int
	Listings []MarketSellListingOnlyItemIdAndQty
}
