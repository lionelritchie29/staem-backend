package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type SentFriendRequest struct{
	UserID uint `gorm:"primaryKey"`
	FriendID uint `gorm:"primaryKey"`
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&SentFriendRequest{})
	db.AutoMigrate(&SentFriendRequest{})

	p := &SentFriendRequest{}
	p.seed(db)
}

func (p *SentFriendRequest) seed(db *gorm.DB) {
	db.Create(&ReceivedFriendRequest{
		UserID:    5,
		FriendID:  1,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&ReceivedFriendRequest{
		UserID:    7,
		FriendID:  1,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&ReceivedFriendRequest{
		UserID:    5,
		FriendID:  2,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&ReceivedFriendRequest{
		UserID:    9,
		FriendID:  3,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&ReceivedFriendRequest{
		UserID:    7,
		FriendID:  4,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&ReceivedFriendRequest{
		UserID:    8,
		FriendID:  4,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

