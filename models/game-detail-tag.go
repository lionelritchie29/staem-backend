package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GameDetailTag struct{
	GameID uint `gorm:"primary_key;autoIncrement:false"`
	GameTagID uint `gorm:"primary_key;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameDetailTag{})
	db.AutoMigrate(&GameDetailTag{})

	p := &GameDetailTag{}
	p.seed(db)
}

func (p *GameDetailTag) seed(db *gorm.DB) {
	db.Create(&GameDetailTag{
		GameID:      1,
		GameTagID: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      1,
		GameTagID: 3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      2,
		GameTagID: 4,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      2,
		GameTagID: 7,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      2,
		GameTagID: 9,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      2,
		GameTagID: 10,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      3,
		GameTagID: 5,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      3,
		GameTagID: 8,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      3,
		GameTagID: 9,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      3,
		GameTagID: 11,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      3,
		GameTagID: 10,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      4,
		GameTagID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      5,
		GameTagID: 6,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      5,
		GameTagID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      5,
		GameTagID: 7,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      6,
		GameTagID: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      6,
		GameTagID: 7,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      7,
		GameTagID: 5,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      7,
		GameTagID: 8,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      7,
		GameTagID: 4,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      8,
		GameTagID: 4,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      8,
		GameTagID: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      9,
		GameTagID: 7,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      10,
		GameTagID: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      10,
		GameTagID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      10,
		GameTagID: 3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      10,
		GameTagID: 4,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      11,
		GameTagID: 	  3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      11,
		GameTagID: 	  1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      11,
		GameTagID: 	  2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      12,
		GameTagID: 	  3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      12,
		GameTagID: 	  7,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      13,
		GameTagID: 	  2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      14,
		GameTagID: 	  8,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      15,
		GameTagID: 	  9,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      16,
		GameTagID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
	db.Create(&GameDetailTag{
		GameID:      16,
		GameTagID: 3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      17,
		GameTagID: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      17,
		GameTagID: 2,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      17,
		GameTagID: 3,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      18,
		GameTagID: 6,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      18,
		GameTagID: 7,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      18,
		GameTagID: 8,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      19,
		GameTagID: 9,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      19,
		GameTagID: 10,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      20,
		GameTagID: 6,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&GameDetailTag{
		GameID:      20,
		GameTagID: 8,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
}
