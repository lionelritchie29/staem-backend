package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"syreclabs.com/go/faker"
	"time"
)

type GameDiscussionPost struct {
	ID uint `gorm:"primaryKey"`
	Title string
	Content string
	GameID uint
	UserID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameDiscussionPost{})
	db.AutoMigrate(&GameDiscussionPost{})

	p := &GameDiscussionPost{}
	p.seed(db)
}

func (p *GameDiscussionPost) seed(db *gorm.DB) {
	for i:= 1; i <= 100; i++ {
		startDate := time.Date(2020, 10, 1, 12, 0, 0, 0, time.UTC)
		db.Create(&GameDiscussionPost{
			Title:    	faker.Lorem().Sentence(rand.Intn(5) + 2),
			Content:   faker.Lorem().Sentence(rand.Intn(20) + 5),
			GameID:    uint(rand.Intn(25) + 1),
			UserID:    uint(rand.Intn(10) + 1),
			CreatedAt: faker.Time().Between(startDate, time.Now()),
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		})
	}
}