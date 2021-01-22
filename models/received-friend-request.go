package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type ReceivedFriendRequest struct{
	UserID uint `gorm:"primary_key"`
	FriendID uint `gorm:"primary_key"`
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&ReceivedFriendRequest{})
	db.AutoMigrate(&ReceivedFriendRequest{})

	p := &ReceivedFriendRequest{}
	p.seed(db)
}

func (p *ReceivedFriendRequest) seed(db *gorm.DB) {
	db.Create(&ReceivedFriendRequest{
		UserID:    1,
		FriendID:  5,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&ReceivedFriendRequest{
		UserID:    1,
		FriendID:  7,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&ReceivedFriendRequest{
		UserID:    2,
		FriendID:  5,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&ReceivedFriendRequest{
		UserID:    3,
		FriendID:  9,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&ReceivedFriendRequest{
		UserID:    4,
		FriendID:  7,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&ReceivedFriendRequest{
		UserID:    4,
		FriendID:  8,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&ReceivedFriendRequest{
		UserID:    2,
		FriendID:  1,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}

