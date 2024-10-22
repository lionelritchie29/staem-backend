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
		Point:                    9999,
		DisplayName:              "Gomay",
		RealName:                 "Lionel Ritchie",
		CustomURL:                "lionelskuy",
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "1.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "1.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#222431",
		FeaturedBadgeUrl:         "1.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    29,
		Point:                    1499,
		DisplayName:              "Cindy Margareta",
		RealName:                 "Cindy Margareta",
		CustomURL:                "cindymargaretz",
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "2.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "1.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#222431",
		FeaturedBadgeUrl:         "2.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    15,
		Point:                    1000,
		DisplayName:              "Arthur Morgan",
		RealName:                 "Arthur Morgan",
		CustomURL:                "arthurz",
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "3.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "1.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#222431",
		FeaturedBadgeUrl:         "3.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    15,
		Point:                    2500,
		DisplayName:              "Abigail Roberts",
		RealName:                 "Abigail Roberts",
		CustomURL:                "abigailz",
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "4.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "1.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#222431",
		FeaturedBadgeUrl:         "4.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    23,
		Point:                    2750,
		DisplayName:              "Aloy Dawn",
		RealName:                 "Aloy Dawn",
		CustomURL:                "aloyz",
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "5.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "1.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#222431",
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
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "6.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "1.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#222431",
		FeaturedBadgeUrl:         "6.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    59,
		Point:                    1855,
		DisplayName:              "Hosea Matthews",
		RealName:                 "Hosea Matthews",
		CustomURL:                "hoseaz",
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "7.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "1.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#222431",
		FeaturedBadgeUrl:         "7.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    65,
		Point:                    1497,
		DisplayName:              "Patrick Star",
		RealName:                 "Patrick Star",
		CustomURL:                "patrickz",
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "8.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "1.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#222431",
		FeaturedBadgeUrl:         "8.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    87,
		Point:                    1555,
		DisplayName:              "Spongebob Squarepants",
		RealName:                 "Spongebob Squarepants",
		CustomURL:                "spongebobz",
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "9.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "1.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#222431",
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
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "10.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "1.jpg",
		Summary:                  faker.Lorem().Sentence(15),
		Theme:                    "#222431",
		FeaturedBadgeUrl:         "10.jpg",
		Country:                  "United States of America",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	//admin
	db.Create(&UserProfile{
		Level:                    99,
		Point:                    9999,
		DisplayName:              "Admin User",
		RealName:                 "Admin User",
		CustomURL:                "adminuser",
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "admin.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "admin.jpg",
		Summary:                  "admin",
		Theme:                    "#222431",
		FeaturedBadgeUrl:         "admin.jpg",
		Country:                  "Admin gapunya negara :D",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    99,
		Point:                    9999,
		DisplayName:              "Admin Promo",
		RealName:                 "Admin Promo",
		CustomURL:                "adminpromo",
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "admin.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "admin.jpg",
		Summary:                  "admin",
		Theme:                    "#222431",
		FeaturedBadgeUrl:         "admin.jpg",
		Country:                  "Admin gapunya negara :D",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})

	db.Create(&UserProfile{
		Level:                    99,
		Point:                    9999,
		DisplayName:              "Admin Game",
		RealName:                 "Admin Game",
		CustomURL:                "admingame",
		AvatarFrameUrl:           "1.png",
		ProfilePictureUrl:        "admin.jpg",
		ProfileBackgroundUrl:     "1.jpg",
		MiniProfileBackgroundUrl: "admin.jpg",
		Summary:                  "admin",
		Theme:                    "#222431",
		FeaturedBadgeUrl:         "admin.jpg",
		Country:                  "Admin gapunya negara :D",
		CreatedAt:                time.Time{},
		UpdatedAt:                time.Time{},
		DeletedAt:                nil,
	})
}
