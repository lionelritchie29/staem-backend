package models

import (
	"syreclabs.com/go/faker"
	//"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GameSystemRequirement struct {
	ID uint `gorm:"primary_key"`
	GameID uint
	IsRecommended bool
	Note string
	OperatingSystem string
	Processor string
	Memory string
	Graphics string
	DirectX string
	Storage string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameSystemRequirement{})
	db.AutoMigrate(&GameSystemRequirement{})

	p := &GameSystemRequirement{}
	p.seed(db)

}

func (p *GameSystemRequirement) seed(db *gorm.DB) {
	for i := 1; i <=10; i++  {
		db.Create(&GameSystemRequirement{
			GameID:          uint(i),
			IsRecommended: true,
			Note:            faker.Lorem().Sentence(4),
			OperatingSystem: "Windows 10",
			Processor:       "FX-6350 or Equivalent; Core i5-3570 or Equivalent",
			Memory:          "8 GB RAM",
			Graphics:        "AMD: Radeon 7970/Radeon R9 280x or Equivalent; NVIDIA: GeForce GTX 760 or Equivalent",
			DirectX:         "Version 11",
			Storage:         "50 GB available space",
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
			DeletedAt:       nil,
		})

		db.Create(&GameSystemRequirement{
			GameID:          uint(i),
			IsRecommended: false,
			Note:            faker.Lorem().Sentence(4),
			OperatingSystem: "Windows 10",
			Processor:       "Ryzen 3 1300X or Equivalent; Core i7-4790 or Equivalent",
			Memory:          "16 GB RAM",
			Graphics:        "AMD: Radeon RX 480 or Equivalent; NVIDIA: GeForce GTX 1060 or Equivalent",
			DirectX:         "Version 11",
			Storage:         "50 GB available space",
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
			DeletedAt:       nil,
		})
	}
}

