package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
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
	for i:=1; i<=25; i++ {
		stepCount := rand.Intn(4)+1
		for j:=1; j<=12; j+=stepCount {
			db.Create(&GameDetailGenre{
				GameID:      uint(i),
				GameGenreID: uint(j),
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
				DeletedAt:   nil,
			})
		}
	}

}
