package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GameTag struct {
	ID uint `gorm:"primaryKey"`
	Name string
	IsAdult bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameTag{})
	db.AutoMigrate(&GameTag{})

	p := &GameTag{}
	p.seed(db)
}

func (p *GameTag) seed(db *gorm.DB) {
	db.Create(&GameTag{
		Name:      "Building",
		IsAdult:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "Sandbox",
		IsAdult:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "Simulation",
		IsAdult:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "Indie",
		IsAdult:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "Funny",
		IsAdult:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "War",
		IsAdult:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "Strategy",
		IsAdult:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "Violence",
		IsAdult:   true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "Gore",
		IsAdult:   true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "Nudity",
		IsAdult:   true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "Multiplayer",
		IsAdult:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "Open World",
		IsAdult:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameTag{
		Name:      "Single Player",
		IsAdult:   false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}
