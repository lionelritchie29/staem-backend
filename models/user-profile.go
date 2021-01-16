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

	db.Create(&UserProfile{
		Level:                    15,
		Point:                    499,
		DisplayName:              "Arthur Morgan",
		CustomURL:                "arthurz",
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

	db.Create(&UserProfile{
		Level:                    15,
		Point:                    499,
		DisplayName:              "Abigail Roberts",
		CustomURL:                "abigailz",
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

	db.Create(&UserProfile{
		Level:                    23,
		Point:                    799,
		DisplayName:              "Aloy Dawn",
		CustomURL:                "aloyz",
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

	db.Create(&UserProfile{
		Level:                    120,
		Point:                    1455,
		DisplayName:              "Peter Parker",
		CustomURL:                "peterz",
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

	db.Create(&UserProfile{
		Level:                    59,
		Point:                    1345,
		DisplayName:              "Hosea Matthews",
		CustomURL:                "hoseaz",
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

	db.Create(&UserProfile{
		Level:                    65,
		Point:                    461,
		DisplayName:              "Patrick Star",
		CustomURL:                "patrickz",
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

	db.Create(&UserProfile{
		Level:                    87,
		Point:                    1547,
		DisplayName:              "Spongebob Squarepants",
		CustomURL:                "spongebobz",
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

	db.Create(&UserProfile{
		Level:                    89,
		Point:                    2451,
		DisplayName:              "Sadie Adler",
		CustomURL:                "sadiez",
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
