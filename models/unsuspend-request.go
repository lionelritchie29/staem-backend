package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type UnsuspendRequest struct{
	UserID uint `gorm:"primary_key"`
	Reason string
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&UnsuspendRequest{})
	db.AutoMigrate(&UnsuspendRequest{})

	p := &UnsuspendRequest{}
	p.seed(db)
}

func (p *UnsuspendRequest) seed(db *gorm.DB) {
	db.Create(&UnsuspendRequest{
		UserID:    5,
		Reason:    "PLZZZ UNBAN MEEEEH",
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

