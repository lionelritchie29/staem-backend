package models

import (
	//"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type MarketTransaction struct {
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
	db.DropTableIfExists(&MarketTransaction{})
	db.AutoMigrate(&MarketTransaction{})

	p := &MarketTransaction{}
	p.seed(db)
}

func (p *MarketTransaction) seed(db *gorm.DB) {
	//
}
