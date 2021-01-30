package market_transaction

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
)

func GetAllSellListing() []models.MarketSellListing {
	db := database.GetInstance()
	var listings []models.MarketSellListing
	db.Find(&listings)
	return listings
}

func GetAllBuyListing() []models.MarketBuyListing {
	db := database.GetInstance()
	var listings []models.MarketBuyListing
	db.Find(&listings)
	return listings
}

func GetPaginateSellListing(limit, offset int) models.MarketSellListingPaginate {
	db := database.GetInstance()
	var listings []models.MarketSellListingOnlyItemIdAndQty

	limitStr := strconv.FormatInt(int64(limit), 10)
	offsetStr := strconv.FormatInt(int64(offset), 10)

	db.Debug().Raw("SELECT a.game_item_id, SUM(quantity) AS quantity," +
					"(select price from market_sell_listings c where c.game_item_id = a.game_item_id order by c.game_item_id, price limit 1) AS lowest_price " +
					"FROM market_sell_listings a GROUP BY game_item_id ORDER BY SUM(quantity) DESC LIMIT " + limitStr + " OFFSET " + offsetStr).Scan(&listings)
	res := db.Raw("SELECT game_item_id, COUNT(quantity) AS quantity FROM market_sell_listings GROUP BY game_item_id").Scan(&models.MarketSellListingOnlyItemIdAndQty{})

	listingsPaginate := models.MarketSellListingPaginate{
		TotalItems:  int(res.RowsAffected),
		Listings:    listings,
	}

	return listingsPaginate
}
