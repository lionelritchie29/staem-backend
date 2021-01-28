package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type UserGameItem struct {
	UserID     uint `gorm:"primaryKey"`
	GameItemID uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&UserGameItem{})
	db.AutoMigrate(&UserGameItem{})

	p := &UserGameItem{}
	p.seed(db)
}

func (p *UserGameItem) seed(db *gorm.DB) {
	//for i := 1; i <= 10; i++ {
	//	skipCount := rand.Intn(30) + 1
	//	for j := 1; j <= 200; j += skipCount {
	//		db.Create(&UserGameItem{
	//			UserID:     uint(i),
	//			GameItemID: uint(j),
	//			CreatedAt:  time.Time{},
	//			UpdatedAt:  time.Time{},
	//			DeletedAt:  nil,
	//		})
	//	}
	//}
}
