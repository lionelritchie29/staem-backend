package models

import (
	//"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type Developer struct {
	ID uint `gorm:"primaryKey"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&Developer{})
	db.AutoMigrate(&Developer{})

	p := &Developer{}
	p.seed(db)
}

func (p *Developer) seed(db *gorm.DB) {
	db.Create(&Developer{
		//ID: uuid.Must(uuid.NewRandom()).String(),
		Name: "Valve",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Developer{
		//ID: uuid.Must(uuid.NewRandom()).String(),
		Name: "Rockstar Games",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Developer{
		//ID: uuid.Must(uuid.NewRandom()).String(),
		Name: "Hangar 13",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Developer{
		//ID: uuid.Must(uuid.NewRandom()).String(),
		Name: "Ubisoft",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Developer{
		//ID: uuid.Must(uuid.NewRandom()).String(),
		Name: "CD Projekt Red",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
}
