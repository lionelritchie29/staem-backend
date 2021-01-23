package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"syreclabs.com/go/faker"
	"time"
)

type UserProfile struct {
	ID                       uint `gorm:"primary_key"`
	Level                    int
	Point                    int
	DisplayName              string
	RealName                 string
	CustomURL                string
	AvatarFrameUrl           string
	ProfilePictureUrl        string
	ProfileBackgroundUrl     string
	MiniProfileBackgroundUrl string
	Summary                  string
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
		DisplayName:              "Gomay",
		RealName:                 "Lionel Ritchie",
		CustomURL:                "lionelskuy",
		AvatarFrameUrl:           "1.jpg",
		ProfilePictureUrl:        "1.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "1.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#999999",
		FeaturedBadgeUrl:         "1.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    29,
		Point:                    299,
		DisplayName:              "Cindy Margareta",
		RealName:                 "Cindy Margareta",
		CustomURL:                "cindymargaretz",
		AvatarFrameUrl:           "2.jpg",
		ProfilePictureUrl:        "2.jpg",
		ProfileBackgroundUrl:     "2.jpg",
		MiniProfileBackgroundUrl: "2.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#CCCCCC",
		FeaturedBadgeUrl:         "2.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    15,
		Point:                    499,
		DisplayName:              "Arthur Morgan",
		RealName:                 "Arthur Morgan",
		CustomURL:                "arthurz",
		AvatarFrameUrl:           "3.jpg",
		ProfilePictureUrl:        "3.jpg",
		ProfileBackgroundUrl:     "3.jpg",
		MiniProfileBackgroundUrl: "3.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#CCCCCC",
		FeaturedBadgeUrl:         "3.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    15,
		Point:                    499,
		DisplayName:              "Abigail Roberts",
		RealName:                 "Abigail Roberts",
		CustomURL:                "abigailz",
		AvatarFrameUrl:           "4.jpg",
		ProfilePictureUrl:        "4.jpg",
		ProfileBackgroundUrl:     "4.jpg",
		MiniProfileBackgroundUrl: "4.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#CCCCCC",
		FeaturedBadgeUrl:         "4.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    23,
		Point:                    799,
		DisplayName:              "Aloy Dawn",
		RealName:                 "Aloy Dawn",
		CustomURL:                "aloyz",
		AvatarFrameUrl:           "5.jpg",
		ProfilePictureUrl:        "5.jpg",
		ProfileBackgroundUrl:     "5.jpg",
		MiniProfileBackgroundUrl: "5.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#CCCCCC",
		FeaturedBadgeUrl:         "5.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    120,
		Point:                    1455,
		DisplayName:              "Peter Parker",
		RealName:                 "Peter Parker",
		CustomURL:                "peterz",
		AvatarFrameUrl:           "6.jpg",
		ProfilePictureUrl:        "6.jpg",
		ProfileBackgroundUrl:     "6.jpg",
		MiniProfileBackgroundUrl: "6.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#CCCCCC",
		FeaturedBadgeUrl:         "6.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    59,
		Point:                    1345,
		DisplayName:              "Hosea Matthews",
		RealName:                 "Hosea Matthews",
		CustomURL:                "hoseaz",
		AvatarFrameUrl:           "7.jpg",
		ProfilePictureUrl:        "7.jpg",
		ProfileBackgroundUrl:     "7.jpg",
		MiniProfileBackgroundUrl: "7.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#CCCCCC",
		FeaturedBadgeUrl:         "7.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    65,
		Point:                    461,
		DisplayName:              "Patrick Star",
		RealName:                 "Patrick Star",
		CustomURL:                "patrickz",
		AvatarFrameUrl:           "8.jpg",
		ProfilePictureUrl:        "8.jpg",
		ProfileBackgroundUrl:     "8.jpg",
		MiniProfileBackgroundUrl: "8.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#CCCCCC",
		FeaturedBadgeUrl:         "8.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    87,
		Point:                    1547,
		DisplayName:              "Spongebob Squarepants",
		RealName:                 "Spongebob Squarepants",
		CustomURL:                "spongebobz",
		AvatarFrameUrl:           "9.jpg",
		ProfilePictureUrl:        "9.jpg",
		ProfileBackgroundUrl:     "9.jpg",
		MiniProfileBackgroundUrl: "9.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#CCCCCC",
		FeaturedBadgeUrl:         "9.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    89,
		Point:                    2451,
		DisplayName:              "Sadie Adler",
		RealName:                 "Sadie Adler",
		CustomURL:                "sadiez",
		AvatarFrameUrl:           "10.jpg",
		ProfilePictureUrl:        "10.jpg",
		ProfileBackgroundUrl:     "10.jpg",
		MiniProfileBackgroundUrl: "10.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#CCCCCC",
		FeaturedBadgeUrl:         "10.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})
}
