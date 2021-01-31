package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"syreclabs.com/go/faker"
	"time"
)

type GameSale struct {
	GameID uint `gorm:"primary_key;autoIncrement:false"`
	Discount int
	ValidTo time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameSale{})
	db.AutoMigrate(&GameSale{})

	p := &GameSale{}
	p.seed(db)
}

func (p *GameSale) seed(db *gorm.DB) {
	startDate := time.Date(2020, time.November, 17, 12, 0, 0, 0, time.UTC)

	db.Create(&GameSale{
		GameID:    1,
		Discount:  55,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    2,
		Discount:  65,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    3,
		Discount:  35,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    5,
		Discount:  70,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    7,
		Discount:  20,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    9,
		Discount:  45,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    12,
		Discount:  65,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    13,
		Discount:  45,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    15,
		Discount:  30,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    19,
		Discount:  45,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    21,
		Discount:  30,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GameSale{
		GameID:    22,
		Discount:  65,
		ValidTo:   faker.Time().Between(startDate, time.Now()),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})


}
