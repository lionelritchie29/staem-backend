package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"strconv"
	"syreclabs.com/go/faker"
	"time"
)

type GameItem struct {
	ID uint `gorm:"primary_key"`
	Name string
	Game int
	GameItemCategory int
	Description string
	Image string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameItem{})
	db.AutoMigrate(&GameItem{})

	p := &GameItem{}
	p.seed(db)
}

func (p *GameItem) seed(db *gorm.DB) {
	for i:=1; i<=25; i++ {
		for j:=1; j<=10; j ++ {
			db.Create(&GameItem{
				Name:      faker.App().Name(),
				Game:      i,
				GameItemCategory: rand.Intn(10) + 1,
				Image:     strconv.FormatInt(int64(j), 10) + ".png",
				Description: faker.Lorem().Sentence(8),
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: nil,
			})

		}
	}
}