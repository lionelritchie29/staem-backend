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

type ChatSticker struct {
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
	db.DropTableIfExists(&ChatSticker{})
	db.AutoMigrate(&ChatSticker{})

	p := &ChatSticker{}
	p.seed(db)
}

func (p *ChatSticker) seed(db *gorm.DB) {
	for i:=1; i<= 12; i++ {
		db.Create(&ChatSticker{
			Name:      faker.App().Name(),
			Url:       strconv.FormatInt(int64(i), 10) + ".png",
			Price:     rand.Intn(2500) + 500,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		})
	}
}
