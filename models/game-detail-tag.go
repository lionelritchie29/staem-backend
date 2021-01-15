package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GameDetailTag struct{
	GameID uint `gorm:"primary_key;autoIncrement:false"`
	GameTagID uint `gorm:"primary_key;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameDetailTag{})
	db.AutoMigrate(&GameDetailTag{})

	p := &GameDetailTag{}
	p.seed(db)
}

func (p *GameDetailTag) seed(db *gorm.DB) {
	db.Create(&GameDetailTag{
		GameID:      1,
		GameTagID: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      1,
		GameTagID: 3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      2,
		GameTagID: 4,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      2,
		GameTagID: 7,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      2,
		GameTagID: 9,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      2,
		GameTagID: 10,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      3,
		GameTagID: 5,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      3,
		GameTagID: 8,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      3,
		GameTagID: 9,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      3,
		GameTagID: 11,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      3,
		GameTagID: 10,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      4,
		GameTagID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
}
