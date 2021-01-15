package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GameImage struct {
	GameID uint `gorm:"primary_key;autoIncrement:false"`
	Url string `gorm:"primary_key;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameImage{})
	db.AutoMigrate(&GameImage{})

	p := &GameImage{}
	p.seed(db)
}

func (p *GameImage) seed(db *gorm.DB) {
	db.Create(&GameImage{
		GameID:    1,
		Url:       "1.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    1,
		Url:       "2.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    1,
		Url:       "3.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    1,
		Url:       "4.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    1,
		Url:       "header.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    2,
		Url:       "1.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    2,
		Url:       "2.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    2,
		Url:       "3.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    2,
		Url:       "4.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    2,
		Url:       "header.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    3,
		Url:       "1.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    3,
		Url:       "2.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    3,
		Url:       "3.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    3,
		Url:       "4.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    3,
		Url:       "header.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    4,
		Url:       "1.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    4,
		Url:       "2.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    4,
		Url:       "3.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    4,
		Url:       "4.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    4,
		Url:       "header.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}