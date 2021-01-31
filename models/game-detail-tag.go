package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
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
	for i:=1; i<=25; i++ {
		stepCount := rand.Intn(4)+1
		for j:=1; j<=12; j+=stepCount {
			db.Create(&GameDetailTag{
				GameID:      uint(i),
				GameTagID: uint(j),
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
				DeletedAt:   nil,
			})
		}
	}
}
