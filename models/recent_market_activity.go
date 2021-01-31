package models

import "time"

type RecentMarketActivity struct {
	Type string
	UserID int
	BuyerID int
	GameItemID int
	Price int
	Quantity int
	TransactionDate time.Time
}
