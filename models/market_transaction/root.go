package market_transaction

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
	"time"
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

	db.Raw("SELECT a.game_item_id, SUM(quantity) AS quantity," +
					"(select price from market_sell_listings c where c.game_item_id = a.game_item_id order by c.game_item_id, price limit 1) AS lowest_price " +
					"FROM market_sell_listings a GROUP BY game_item_id ORDER BY SUM(quantity) DESC LIMIT " + limitStr + " OFFSET " + offsetStr).Scan(&listings)
	res := db.Raw("SELECT game_item_id, COUNT(quantity) AS quantity FROM market_sell_listings GROUP BY game_item_id").Scan(&models.MarketSellListingOnlyItemIdAndQty{})

	listingsPaginate := models.MarketSellListingPaginate{
		TotalItems:  int(res.RowsAffected),
		Listings:    listings,
	}

	return listingsPaginate
}

func GetSellListingsGroupedByPrice(gameItemId int) []models.MarketSellListingGroupedByPrice {
	db := database.GetInstance()
	gameItemIdStr := strconv.FormatInt(int64(gameItemId), 10)
	var listings []models.MarketSellListingGroupedByPrice
	db.Raw("SELECT price, SUM(quantity) as quantity FROM market_sell_listings WHERE game_item_id = " + gameItemIdStr + " GROUP BY price").Scan(&listings)
	return listings
}

func GetBuyListingsGroupedByPrice(gameItemId int) []models.MarketSellListingGroupedByPrice {
	db := database.GetInstance()
	gameItemIdStr := strconv.FormatInt(int64(gameItemId), 10)
	var listings []models.MarketSellListingGroupedByPrice
	db.Raw("SELECT price, SUM(quantity) as quantity FROM market_buy_listings WHERE game_item_id = " + gameItemIdStr + " GROUP BY price").Scan(&listings)
	return listings
}

func GetSellListingByUserIdAndGameItemId(userId, gameItemId int) []models.MarketSellListing {
	db := database.GetInstance()
	var listings []models.MarketSellListing
	db.Debug().Where("user_id = ? AND game_item_id = ?", userId, gameItemId).Find(&listings)
	return listings
}

func GetBuyListingByUserIdAndGameItemId(userId, gameItemId int) []models.MarketBuyListing {
	db := database.GetInstance()
	var listings []models.MarketBuyListing
	db.Debug().Where("user_id = ? AND game_item_id = ?", userId, gameItemId).Find(&listings)
	return listings
}

func CreateSellListing(userId, gameItemId, price, qty int) bool {
	db := database.GetInstance()
	res := db.Create(&models.MarketSellListing{
		UserID:     uint(userId),
		GameItemID: uint(gameItemId),
		Price:      price,
		Quantity:   qty,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})

	if res.Error != nil {
		return false
	}

	return true
}