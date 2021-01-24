package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"syreclabs.com/go/faker"
	"time"
)

type UserReport struct {
	UserID uint `gorm:"primary_key"`
	ReporterID uint `gorm:"primary_key"`
	Reason string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&UserReport{})
	db.AutoMigrate(&UserReport{})

	p := &UserReport{}
	p.seed(db)
}

func (p *UserReport) seed(db *gorm.DB) {
	db.Create(&UserReport{
		UserID:     3,
		ReporterID: 5,
		Reason: faker.Lorem().Sentence(rand.Intn(5) + 3),
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})

	db.Create(&UserReport{
		UserID:     3,
		ReporterID: 6,
		Reason: faker.Lorem().Sentence(rand.Intn(5) + 3),
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})

	db.Create(&UserReport{
		UserID:     3,
		ReporterID: 7,
		Reason: faker.Lorem().Sentence(rand.Intn(5) + 3),
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})

	db.Create(&UserReport{
		UserID:     5,
		ReporterID: 9,
		Reason: faker.Lorem().Sentence(rand.Intn(5) + 3),
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})

	db.Create(&UserReport{
		UserID:     5,
		ReporterID: 10,
		Reason: faker.Lorem().Sentence(rand.Intn(5) + 3),
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	})

}