package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type UserWishlist struct {
	UserID uint `gorm:"primary_key"`
	GameID uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&UserWishlist{})
	db.AutoMigrate(&UserWishlist{})

	p := &UserWishlist{}
	p.seed(db)
}

func (p *UserWishlist) seed(db *gorm.DB) {
	db.Create(&UserWishlist{
		UserID:    1,
		GameID:    1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&UserWishlist{
		UserID:    1,
		GameID:    11,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&UserWishlist{
		UserID:    2,
		GameID:    3,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}
