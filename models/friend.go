package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
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
	db.Create(&Friend{
		UserID:    1,
		FriendID:  2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    2,
		FriendID:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    1,
		FriendID:  3,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    3,
		FriendID:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    1,
		FriendID:  4,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    4,
		FriendID:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    2,
		FriendID:  3,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    3,
		FriendID:  2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    2,
		FriendID:  4,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    4,
		FriendID:  2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    5,
		FriendID:  4,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    4,
		FriendID:  5,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    5,
		FriendID:  6,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    6,
		FriendID:  5,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    5,
		FriendID:  7,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    7,
		FriendID:  5,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    6,
		FriendID:  8,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    8,
		FriendID:  6,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    8,
		FriendID:  9,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    9,
		FriendID:  8,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    10,
		FriendID:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    1,
		FriendID:  10,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    9,
		FriendID:  10,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    10,
		FriendID:  9,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    1,
		FriendID:  8,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    8,
		FriendID:  1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    2,
		FriendID:  10,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    10,
		FriendID:  2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    2,
		FriendID:  9,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    9,
		FriendID:  2,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&Friend{
		UserID:    3,
		FriendID:  8,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
	db.Create(&Friend{
		UserID:    8,
		FriendID:  3,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

}

