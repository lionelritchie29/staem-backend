package models

import (
	//"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type Publisher struct {
	ID uint `gorm:"primaryKey"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&Publisher{})
	db.AutoMigrate(&Publisher{})

	p := &Publisher{}
	p.seed(db)

}

func (p *Publisher) seed(db *gorm.DB) {
	db.Create(&Publisher{
		//ID: uuid.Must(uuid.NewRandom()).String(),
		Name: "XBox Ltd. Game Studios",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Publisher{
		//ID: uuid.Must(uuid.NewRandom()).String(),
		Name: "2K",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Publisher{
		//ID: uuid.Must(uuid.NewRandom()).String(),
		Name: "Rockstar Games",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	db.Create(&Publisher{
		//ID: uuid.Must(uuid.NewRandom()).String(),
		Name: "Valve",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})
}

