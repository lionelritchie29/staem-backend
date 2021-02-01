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

type ImageVideoPost struct {
	ID uint `gorm:"primary_key"`
	FileURL string
	Type string
	Description string
	LikeCount int
	DislikeCount int
	UserID int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&ImageVideoPost{})
	db.AutoMigrate(&ImageVideoPost{})

	p := &ImageVideoPost{}
	p.seed(db)
}

func (p *ImageVideoPost) seed(db *gorm.DB) {
	for i := 1; i<= 10; i++ {
		db.Create(&ImageVideoPost{
			FileURL:      strconv.FormatInt(int64(i), 10) + ".jpg",
			Type:         "image",
			Description:  faker.Lorem().Sentence(rand.Intn(15) + 5),
			LikeCount:    rand.Intn(30) + 1,
			DislikeCount: rand.Intn(30) + 1,
			UserID:       rand.Intn(10) + 1,
			CreatedAt:    time.Time{},
			UpdatedAt:    time.Time{},
			DeletedAt:    nil,
		})
	}
}
