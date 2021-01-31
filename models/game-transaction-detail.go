package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"time"
)

type GameTransactionDetail struct {
	ID uint `gorm:"primary_key"`
	GameTransactionID uint
	Game int
	Price int
	Quantity int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameTransactionDetail{})
	db.AutoMigrate(&GameTransactionDetail{})

	p := &GameTransactionDetail{}
	p.seed(db)
}

func (p *GameTransactionDetail) seed(db *gorm.DB) {
	for i := 0; i < 225; i++ {
		db.Create(&GameTransactionDetail{
			GameTransactionID: uint(rand.Intn(100) + 1),
			Game:    rand.Intn(25) + 1,
			Price: rand.Intn(1099999) + 169999,
			Quantity:  rand.Intn(20) + 1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		})
	}
}
