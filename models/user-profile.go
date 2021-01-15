package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type UserProfile struct {
	ID                       uint `gorm:"primary_key"`
	Level                    int
	Point                    int
	DisplayName              string
	CustomURL                string
	AvatarFrameUrl           string
	ProfilePictureUrl        string
	ProfileBackgroundUrl     string
	MiniProfileBackgroundUrl string
	Theme                    string
	FeaturedBadgeUrl         string
	Country                  string
	CreatedAt                time.Time
	UpdatedAt                time.Time
	DeletedAt                *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&UserProfile{})
	db.AutoMigrate(&UserProfile{})

	p := &UserProfile{}
	p.seed(db)
}

func (p *UserProfile) seed(db *gorm.DB) {
	db.Create(&UserProfile{
		Level:                    99,
		Point:                    999,
		DisplayName:              "Lionel Ritchie",
		CustomURL:                "lionelskuy",
		AvatarFrameUrl:           "test.jpg",
		ProfilePictureUrl:        "test.jpg",
		ProfileBackgroundUrl:     "test.jpg",
		MiniProfileBackgroundUrl: "test.jpg",
		Theme:                    "#999999",
		FeaturedBadgeUrl:         "test.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    29,
		Point:                    299,
		DisplayName:              "Cindy Margareta",
		CustomURL:                "cindymargaretz",
		AvatarFrameUrl:           "test.jpg",
		ProfilePictureUrl:        "test.jpg",
		ProfileBackgroundUrl:     "test.jpg",
		MiniProfileBackgroundUrl: "test.jpg",
		Theme:                    "#CCCCCC",
		FeaturedBadgeUrl:         "test.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})
}
