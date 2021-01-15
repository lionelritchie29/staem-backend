package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GameDetailCategory struct {
	GameID         uint `gorm:"primary_Key;autoIncrement:false"`
	GameCategoryID uint `gorm:"primary_key;autoIncrement:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameDetailCategory{})
	db.AutoMigrate(&GameDetailCategory{})

	p := &GameDetailCategory{}
	p.seed(db)
}

func (p *GameDetailCategory) seed(db *gorm.DB) {
	db.Create(&GameDetailCategory{
		GameID:         1,
		GameCategoryID: 1,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

	db.Create(&GameDetailCategory{
		GameID:         1,
		GameCategoryID: 2,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

	db.Create(&GameDetailCategory{
		GameID:         2,
		GameCategoryID: 3,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

	db.Create(&GameDetailCategory{
		GameID:         2,
		GameCategoryID: 4,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

	db.Create(&GameDetailCategory{
		GameID:         3,
		GameCategoryID: 1,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})
}
