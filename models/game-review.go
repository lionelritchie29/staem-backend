package models

import (
	"math/rand"
	"syreclabs.com/go/faker"
	//"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GameReview struct {
	ID uint `gorm:"primary_key"`
	IsRecommended bool
	GameID uint
	UserID uint
	Content string
	UpvoteCount int
	DownvoteCount int
	ReviewDateTime time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameReview{})
	db.AutoMigrate(&GameReview{})

	p := &GameReview{}
	p.seed(db)

}

func (p *GameReview) seed(db *gorm.DB) {

	for i := 0; i < 150; i++ {
		startDate := time.Date(2020, time.December, 17, 12, 0, 0, 0, time.UTC)
		booleanRandom :=  rand.Intn(10) + 1
		var isRecommended bool

		if booleanRandom == 1 {
			isRecommended = true
		}else{
			isRecommended = false
		}

		db.Create(&GameReview{
			IsRecommended: 	isRecommended,
			GameID:         uint( rand.Intn(15) + 1 ),
			UserID:         uint( rand.Intn(10) + 1 ),
			Content:        faker.Lorem().Paragraph(rand.Intn(10) + 3),
			UpvoteCount:    rand.Intn(100 + 1 ),
			DownvoteCount:  rand.Intn(100+ 1 ),
			ReviewDateTime: faker.Time().Between(startDate, time.Now()),
			CreatedAt:      time.Time{},
			UpdatedAt:      time.Time{},
			DeletedAt:      nil,
		})
	}
}

