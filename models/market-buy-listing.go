package models

import (
	//"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type MarketBuyListing struct {
	ID uint `gorm:"primary_key"`
	UserID uint
	GameItemID uint
	Price int
	Quantity int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&MarketBuyListing{})
	db.AutoMigrate(&MarketBuyListing{})

	p := &MarketBuyListing{}
	p.seed(db)
}

func (p *MarketBuyListing) seed(db *gorm.DB) {
//	Seeded at user-game.go
}
