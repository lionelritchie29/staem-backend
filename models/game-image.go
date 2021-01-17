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
		Url:       "5.jpg",
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
		Url:       "5.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    2,
		Url:       "6.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    2,
		Url:       "7.jpg",
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

	//
	db.Create(&GameImage{
		GameID:    5,
		Url:       "1.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    5,
		Url:       "2.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    5,
		Url:       "3.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    5,
		Url:       "4.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    5,
		Url:       "header.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	//
	db.Create(&GameImage{
		GameID:    6,
		Url:       "1.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    6,
		Url:       "2.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    6,
		Url:       "3.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    6,
		Url:       "4.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    6,
		Url:       "header.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	//
	db.Create(&GameImage{
		GameID:    7,
		Url:       "1.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    7,
		Url:       "2.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    7,
		Url:       "3.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    7,
		Url:       "4.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    7,
		Url:       "header.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	//
	db.Create(&GameImage{
		GameID:    8,
		Url:       "1.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    8,
		Url:       "2.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    8,
		Url:       "3.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    8,
		Url:       "4.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    8,
		Url:       "header.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	//
	db.Create(&GameImage{
		GameID:    9,
		Url:       "1.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    9,
		Url:       "2.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    9,
		Url:       "3.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    9,
		Url:       "4.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    9,
		Url:       "header.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	//
	db.Create(&GameImage{
		GameID:    10,
		Url:       "1.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    10,
		Url:       "2.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    10,
		Url:       "3.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    10,
		Url:       "4.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameImage{
		GameID:    10,
		Url:       "header.jpg",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}