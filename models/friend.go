package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"time"
)

type Friend struct{
	UserID uint `gorm:"primaryKey"`
	FriendID uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&Friend{})
	db.AutoMigrate(&Friend{})

	p := &Friend{}
	p.seed(db)
}

func (p *Friend) seed(db *gorm.DB) {
	for i:=1; i<=10; i++ {
		skipCount := rand.Intn(5) + 1
		for j:=1; j<=10; j+=skipCount {
			if j == i {
				continue
			} else {
				db.Create(&Friend{
					UserID:    uint(i),
					FriendID:  uint(j),
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: nil,
				})
			}
		}
	}
}

