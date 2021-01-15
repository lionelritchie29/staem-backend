package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type UserRole struct {
	ID uint `gorm:"primary_key"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&UserRole{})
	db.AutoMigrate(&UserRole{})

	p := &UserRole{}
	p.seed(db)
}

func (p *UserRole) seed(db *gorm.DB) {
	db.Create(&UserRole{
		Name:      "user-default",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&UserRole{
		Name:      "admin-game",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&UserRole{
		Name:      "admin-promo",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&UserRole{
		Name:      "admin-user",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&UserRole{
		Name:      "admin-master",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}
