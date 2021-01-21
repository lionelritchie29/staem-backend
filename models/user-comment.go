package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"syreclabs.com/go/faker"
	"time"
)

type UserComment struct {
	SrcUserId uint `gorm:"primary_key"`
	DestUserId uint `gorm:"primary_key"`
	Content string
	CreatedAt time.Time `gorm:"primary_key"`
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&UserComment{})
	db.AutoMigrate(&UserComment{})

	p := &UserComment{}
	p.seed(db)
}

func (p *UserComment) seed(db *gorm.DB) {
	for i:=1; i<=10; i++ {
		randStep := rand.Intn(4) + 1
		for j:=1; j<=10; j += randStep {
			if i == j {
				continue
			} else {
				db.Create(&UserComment{
					SrcUserId:  uint(i),
					DestUserId: uint(j),
					Content:    faker.Lorem().Sentence(rand.Intn(35) + 4),
					CreatedAt:  time.Time{},
					UpdatedAt:  time.Time{},
					DeletedAt:  nil,
				})
			}
		}
	}
}
