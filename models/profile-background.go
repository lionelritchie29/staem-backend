package models

import (
	//"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"strconv"
	"syreclabs.com/go/faker"
	"time"
)

type ProfileBackground struct {
	ID uint `gorm:"primary_key"`
	Name string
	Url string
	Price int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&ProfileBackground{})
	db.AutoMigrate(&ProfileBackground{})

	p := &ProfileBackground{}
	p.seed(db)
}

func (p *ProfileBackground) seed(db *gorm.DB) {
	for i:=1; i<= 12; i++ {
		db.Create(&ProfileBackground{
			Name:      faker.App().Name(),
			Url:       strconv.FormatInt(int64(i), 10) + ".jpg",
			Price:     rand.Intn(2500) + 500,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		})
	}
}
