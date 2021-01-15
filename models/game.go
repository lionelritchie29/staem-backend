package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type Game struct {
	ID uint `gorm:"primaryKey"`
	Publisher int // foreign key to Publisher
	Developer int // foreign key to Developer
	Title string
	Description string
	Price int
	ReleaseDate string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&Game{})
	db.AutoMigrate(&Game{})

	p := &Game{}
	p.seed(db)
}

func (p *Game) seed(db *gorm.DB) {
	db.Create(&Game{
		Publisher:   3,
		Developer:   2,
		Title:       "Red Dead Redemption 2",
		Description: "Winner of over 175 Game of the Year Awards and recipient of over 250 perfect scores, RDR2 is the epic tale of outlaw Arthur Morgan and the infamous Van der Linde gang, on the run across America at the dawn of the modern age. Also includes access to the shared living world of Red Dead Online.",
		Price:       690000,
		ReleaseDate: "2019-12-06",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   4,
		Developer:   4,
		Title:       "Cyberpunk 2077",
		Description: "Cyberpunk 2077 is an open-world, action-adventure story set in Night City, a megalopolis obsessed with power, glamour and body modification. You play as V, a mercenary outlaw going after a one-of-a-kind implant that is the key to immortality.",
		Price:       699999,
		ReleaseDate: "2020-12-10",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   2,
		Developer:   4,
		Title:       "Watch Dogs 2",
		Description: "Welcome to San Francisco. Play as Marcus, a brilliant young hacker, and join the most notorious hacker group, DedSec. Your objective: execute the biggest hack of history.",
		Price:       619000,
		ReleaseDate: "2016-11-28",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})

	db.Create(&Game{
		Publisher:   3,
		Developer:   4,
		Title:       "Overcooked 2",
		Description: "Overcooked returns with a brand-new helping of chaotic cooking action! Journey back to the Onion Kingdom and assemble your team of chefs in classic couch co-op or online play for up to four players. Hold onto your aprons… it’s time to save the world again!",
		Price:       199000,
		ReleaseDate: "2018-08-07",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	})
}


