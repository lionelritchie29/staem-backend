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

type MiniProfileBackground struct {
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
	db.DropTableIfExists(&MiniProfileBackground{})
	db.AutoMigrate(&MiniProfileBackground{})

	p := &MiniProfileBackground{}
	p.seed(db)
}

func (p *MiniProfileBackground) seed(db *gorm.DB) {
	for i:=1; i<= 12; i++ {
		db.Create(&MiniProfileBackground{
			Name:      faker.App().Name(),
			Url:       strconv.FormatInt(int64(i), 10) + ".jpg",
			Price:     rand.Intn(2500) + 500,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		})
	}
}
