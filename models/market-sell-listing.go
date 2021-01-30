package models

import (
	//"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type MarketSellListing struct {
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
	db.DropTableIfExists(&MarketSellListing{})
	db.AutoMigrate(&MarketSellListing{})

	p := &MarketSellListing{}
	p.seed(db)
}

func (p *MarketSellListing) seed(db *gorm.DB) {
//	Seeded at user-game.go
}
