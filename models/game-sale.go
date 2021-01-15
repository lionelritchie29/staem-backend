package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GameSale struct {
	GameID uint `gorm:"primary_key;autoIncrement:false"`
	Discount int
	ValidTo string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameSale{})
	db.AutoMigrate(&GameSale{})

	p := &GameSale{}
	p.seed(db)
}

func (p *GameSale) seed(db *gorm.DB) {
	db.Create(&GameSale{
		GameID:    1,
		Discount:  55,
		ValidTo:   time.Now().Format("2006-01-02"),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    2,
		Discount:  65,
		ValidTo:   time.Now().Format("2006-01-02"),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

}
