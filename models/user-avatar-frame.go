package models

import (
	//"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"time"
)

type UserAvatarFrame struct {
	UserID uint `gorm:"primary_key"`
	AvatarFrameID uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&UserAvatarFrame{})
	db.AutoMigrate(&UserAvatarFrame{})

	p := &UserAvatarFrame{}
	p.seed(db)
}

func (p *UserAvatarFrame) seed(db *gorm.DB) {
	for i:=1; i<= 10; i++ {
		stepCount := rand.Intn(5) + 1
		for j := 1; j <= 12; j += stepCount {
			db.Create(&UserAvatarFrame{
				UserID:        uint(i),
				AvatarFrameID: uint(j),
				CreatedAt:     time.Time{},
				UpdatedAt:     time.Time{},
				DeletedAt:     nil,
			})
		}
	}
}
