package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GameItemCategory struct {
	ID uint `gorm:"primary_key;autoIncrement:false"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameItemCategory{})
	db.AutoMigrate(&GameItemCategory{})

	p := &GameItemCategory{}
	p.seed(db)
}

func (p *GameItemCategory) seed(db *gorm.DB) {
	db.Create(&GameItemCategory{
		Name:      "Arm",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameItemCategory{
		Name:      "Shoulder",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameItemCategory{
		Name:      "Off-Hand",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameItemCategory{
		Name:      "Tail",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameItemCategory{
		Name:      "Weapon",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameItemCategory{
		Name:      "Cursor Pack",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameItemCategory{
		Name:      "Belt",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameItemCategory{
		Name:      "Back",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameItemCategory{
		Name:      "Shoulder",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameItemCategory{
		Name:      "Music",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}