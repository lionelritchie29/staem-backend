package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GameCategory struct {
	ID uint `gorm:"primary_key"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameCategory{})
	db.AutoMigrate(&GameCategory{})

	p := &GameCategory{}
	p.seed(db)
}

func (p *GameCategory) seed(db *gorm.DB) {
	db.Create(&GameCategory{
		Name:      "Top Sellers",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameCategory{
		Name:      "New Releases",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameCategory{
		Name:      "Upcoming",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameCategory{
		Name:      "Specials",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}