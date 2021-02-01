package models

import (
	//"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"syreclabs.com/go/faker"
	"time"
)

type ImageVideoComment struct {
	PostID uint `gorm:"primary_key"`
	UserID uint `gorm:"primary_key"`
	Comments string
	CreatedAt time.Time `gorm:"primary_key"`
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&ImageVideoComment{})
	db.AutoMigrate(&ImageVideoComment{})

	p := &ImageVideoComment{}
	p.seed(db)
}

func (p *ImageVideoComment) seed(db *gorm.DB) {
	for i := 1; i<= 50; i++ {
		startTime := time.Date(2020, 10, 10, 12, 0, 0, 0, time.UTC)
		db.Create(&ImageVideoComment{
			PostID: uint(rand.Intn(10) + 1),
			UserID:           uint(rand.Intn(10)),
			Comments:         faker.Lorem().Sentence(rand.Intn(10) + 3),
			CreatedAt:        faker.Time().Between(startTime, time.Now()),
			UpdatedAt:        time.Time{},
			DeletedAt:        nil,
		})
	}
}
