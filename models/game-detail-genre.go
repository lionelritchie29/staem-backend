package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GameDetailGenre struct{
	GameID uint `gorm:"primary_key;autoIncrement:false"`
	GameGenreID uint `gorm:"primary_key;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameDetailGenre{})
	db.AutoMigrate(&GameDetailGenre{})

	p := &GameDetailGenre{}
	p.seed(db)
}

func (p *GameDetailGenre) seed(db *gorm.DB) {
	db.Create(&GameDetailGenre{
		GameID:      1,
		GameGenreID: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      1,
		GameGenreID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      1,
		GameGenreID: 3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      2,
		GameGenreID: 5,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      3,
		GameGenreID: 6,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      3,
		GameGenreID: 7,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      3,
		GameGenreID: 8,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      3,
		GameGenreID: 9,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      4,
		GameGenreID: 8,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      4,
		GameGenreID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      4,
		GameGenreID: 10,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      4,
		GameGenreID: 11,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      4,
		GameGenreID: 12,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      5,
		GameGenreID: 8,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      5,
		GameGenreID: 7,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      5,
		GameGenreID: 11,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      6,
		GameGenreID: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      6,
		GameGenreID: 7,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      6,
		GameGenreID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      6,
		GameGenreID: 4,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      7,
		GameGenreID: 3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      8,
		GameGenreID: 9,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      9,
		GameGenreID: 9,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      9,
		GameGenreID: 3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      10,
		GameGenreID: 3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      11,
		GameGenreID: 3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      11,
		GameGenreID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      12,
		GameGenreID: 5,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      12,
		GameGenreID: 6,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      12,
		GameGenreID: 10,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      13,
		GameGenreID: 5,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      13,
		GameGenreID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      13,
		GameGenreID: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      14,
		GameGenreID: 3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      14,
		GameGenreID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailGenre{
		GameID:      15,
		GameGenreID: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
}
