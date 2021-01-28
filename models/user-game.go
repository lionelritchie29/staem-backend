package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"time"
)

type UserGame struct {
	UserID    uint `gorm:"primaryKey"`
	GameID    uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&UserGame{})
	db.AutoMigrate(&UserGame{})

	p := &UserGame{}
	p.seed(db)
}

func (p *UserGame) seed(db *gorm.DB) {
	for i := 1; i <= 10; i++ {
		skipCount := rand.Intn(5) + 1
		for j := 1; j <= 20; j += skipCount {
			db.Create(&UserGame{
				UserID:    uint(i),
				GameID:    uint(j),
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: nil,
			})

			skipCount2 := rand.Intn(2) + 1
			for k := 1; k <= 10; k += skipCount2 {
				db.Create(&UserGameItem{
					UserID:     uint(i),
					GameItemID: uint((j-1) * 10 + k),
					CreatedAt:  time.Time{},
					UpdatedAt:  time.Time{},
					DeletedAt:  nil,
				})
			}

		}

	}
}
