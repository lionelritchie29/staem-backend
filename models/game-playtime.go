package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"strconv"
	"syreclabs.com/go/faker"
	"time"
)

type GamePlaytime struct {
	GameID uint `gorm:"primary_key;autoIncrement:false"`
	PlayHour int `gorm:"primary_key;autoIncrement:false"`
	Date string `gorm:"primary_key;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GamePlaytime{})
	db.AutoMigrate(&GamePlaytime{})

	p := &GamePlaytime{}
	p.seed(db)
}

func (p *GamePlaytime) seed(db *gorm.DB) {

	startDate := time.Date(2020, time.November, 1, 0, 0, 0, 0, time.UTC)
	for i:=1; i<=20; i++ {
		randomCount := rand.Intn(30) + 1;
		for j:=1; j<= randomCount; j++ {
			randDate := faker.Date().Between(startDate, time.Now())
			date := strconv.FormatInt(int64(randDate.Year()), 10) + "-" + strconv.FormatInt(int64(randDate.Month()),10) + "-" + strconv.FormatInt(int64(randDate.Day()), 10)
			db.Create(&GamePlaytime{
				GameID:    uint(i),
				PlayHour:  rand.Intn(500) + 100,
				Date:      date,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: nil,
			})
		}
	}
}