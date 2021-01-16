package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type GamePlaytime struct {
	GameID uint `gorm:"primary_key;autoIncrement:false"`
	PlayHour int
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
	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  79,
		Date:      "2020-12-01",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    2,
		PlayHour:  125,
		Date:      "2020-12-05",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    3,
		PlayHour:  142,
		Date:      "2020-12-03",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    4,
		PlayHour:  56,
		Date:      "2020-12-04",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  66,
		Date:      "2020-12-07",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    4,
		PlayHour:  27,
		Date:      "2020-12-15",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    3,
		PlayHour:  98,
		Date:      "2020-12-17",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    2,
		PlayHour:  211,
		Date:      "2020-12-18",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  144,
		Date:      "2020-12-13",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    3,
		PlayHour:  316,
		Date:      "2020-12-19",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  80,
		Date:      "2020-12-22",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    3,
		PlayHour:  91,
		Date:      "2020-12-18",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    4,
		PlayHour:  120,
		Date:      "2020-12-29",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  33,
		Date:      "2020-12-26",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    4,
		PlayHour:  77,
		Date:      "2020-12-28",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  79,
		Date:      "2020-12-30",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    3,
		PlayHour:  88,
		Date:      "2020-12-31",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  126,
		Date:      "2020-12-24",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    2,
		PlayHour:  279,
		Date:      "2020-12-01",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    3,
		PlayHour:  117,
		Date:      "2020-11-28",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    4,
		PlayHour:  123,
		Date:      "2020-11-24",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    2,
		PlayHour:  145,
		Date:      "2020-11-21",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    3,
		PlayHour:  164,
		Date:      "2020-11-24",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    4,
		PlayHour:  134,
		Date:      "2020-11-19",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  81,
		Date:      "2020-11-17",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  40,
		Date:      "2020-11-14",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    3,
		PlayHour:  126,
		Date:      "2020-11-16",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  140,
		Date:      "2020-11-09",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    2,
		PlayHour:  147,
		Date:      "2020-11-07",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    3,
		PlayHour:  215,
		Date:      "2020-11-02",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  154,
		Date:      "2020-11-07",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    2,
		PlayHour:  153,
		Date:      "2021-01-02",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    2,
		PlayHour:  153,
		Date:      "2021-01-05",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    3,
		PlayHour:  79,
		Date:      "2021-01-07",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    2,
		PlayHour:  102,
		Date:      "2021-01-08",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    3,
		PlayHour:  111,
		Date:      "2021-01-09",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  158,
		Date:      "2021-01-11",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  136,
		Date:      "2021-01-13",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    4,
		PlayHour:  155,
		Date:      "2021-01-10",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    1,
		PlayHour:  136,
		Date:      "2021-01-01",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    5,
		PlayHour:  154,
		Date:      "2020-11-01",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    5,
		PlayHour:  240,
		Date:      "2020-11-08",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    5,
		PlayHour:  89,
		Date:      "2020-11-15",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    5,
		PlayHour:  111,
		Date:      "2020-11-22",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    5,
		PlayHour:  17,
		Date:      "2020-11-29",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    5,
		PlayHour:  159,
		Date:      "2020-12-06",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    5,
		PlayHour:  159,
		Date:      "2020-12-13",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    5,
		PlayHour:  264,
		Date:      "2020-12-20",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    5,
		PlayHour:  310,
		Date:      "2020-12-27",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    5,
		PlayHour:  123,
		Date:      "2021-01-03",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	//
	db.Create(&GamePlaytime{
		GameID:    6,
		PlayHour:  124,
		Date:      "2020-11-01",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    6,
		PlayHour:  144,
		Date:      "2020-11-08",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    6,
		PlayHour:  177,
		Date:      "2020-11-15",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    6,
		PlayHour:  240,
		Date:      "2020-11-22",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    6,
		PlayHour:  112,
		Date:      "2020-11-29",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    6,
		PlayHour:  124,
		Date:      "2020-12-06",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    6,
		PlayHour:  147,
		Date:      "2020-12-13",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    6,
		PlayHour:  187,
		Date:      "2020-12-20",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    6,
		PlayHour:  254,
		Date:      "2020-12-27",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    6,
		PlayHour:  174,
		Date:      "2021-01-03",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	//
	db.Create(&GamePlaytime{
		GameID:    7,
		PlayHour:  124,
		Date:      "2020-11-01",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    7,
		PlayHour:  154,
		Date:      "2020-11-08",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    7,
		PlayHour:  199,
		Date:      "2020-11-15",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    7,
		PlayHour:  185,
		Date:      "2020-11-22",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    7,
		PlayHour:  254,
		Date:      "2020-11-29",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    7,
		PlayHour:  28,
		Date:      "2020-12-06",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    7,
		PlayHour:  145,
		Date:      "2020-12-13",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    7,
		PlayHour:  88,
		Date:      "2020-12-20",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    7,
		PlayHour:  95,
		Date:      "2020-12-27",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    7,
		PlayHour:  310,
		Date:      "2021-01-03",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	//
	db.Create(&GamePlaytime{
		GameID:    8,
		PlayHour:  214,
		Date:      "2020-11-01",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    8,
		PlayHour:  254,
		Date:      "2020-11-08",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    8,
		PlayHour:  267,
		Date:      "2020-11-15",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    8,
		PlayHour:  154,
		Date:      "2020-11-22",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    8,
		PlayHour:  254,
		Date:      "2020-11-29",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    8,
		PlayHour:  128,
		Date:      "2020-12-06",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    8,
		PlayHour:  177,
		Date:      "2020-12-13",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    8,
		PlayHour:  159,
		Date:      "2020-12-20",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    8,
		PlayHour:  195,
		Date:      "2020-12-27",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    8,
		PlayHour:  124,
		Date:      "2021-01-03",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	//
	db.Create(&GamePlaytime{
		GameID:    9,
		PlayHour:  259,
		Date:      "2020-11-01",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    9,
		PlayHour:  311,
		Date:      "2020-11-08",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    9,
		PlayHour:  277,
		Date:      "2020-11-15",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    9,
		PlayHour:  259,
		Date:      "2020-11-22",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    9,
		PlayHour:  277,
		Date:      "2020-11-29",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    9,
		PlayHour:  236,
		Date:      "2020-12-06",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    9,
		PlayHour:  133,
		Date:      "2020-12-13",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    9,
		PlayHour:  224,
		Date:      "2020-12-20",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    9,
		PlayHour:  257,
		Date:      "2020-12-27",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    9,
		PlayHour:  333,
		Date:      "2021-01-03",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	//
	db.Create(&GamePlaytime{
		GameID:    10,
		PlayHour:  147,
		Date:      "2020-11-01",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    10,
		PlayHour:  311,
		Date:      "2020-11-08",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    10,
		PlayHour:  258,
		Date:      "2020-11-15",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    10,
		PlayHour:  133,
		Date:      "2020-11-22",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    10,
		PlayHour:  77,
		Date:      "2020-11-29",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    10,
		PlayHour:  89,
		Date:      "2020-12-06",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    10,
		PlayHour:  123,
		Date:      "2020-12-13",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    10,
		PlayHour:  77,
		Date:      "2020-12-20",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    10,
		PlayHour:  257,
		Date:      "2020-12-27",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&GamePlaytime{
		GameID:    10,
		PlayHour:  335,
		Date:      "2021-01-03",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}