package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/helpers"
	"time"
)

type UserAccount struct {
	ID           uint `gorm:"primary_key"`
	AccountName  string
	Email        string
	Password     string
	WalletAmount int
	Role         int //foreign key to Role
	SuspendedAt  time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&UserAccount{})
	db.AutoMigrate(&UserAccount{})

	p := &UserAccount{}
	p.seed(db)
}

func (p *UserAccount) seed(db *gorm.DB) {

	pass1, _ := helpers.HashPassword("lionel123")
	db.Create(&UserAccount{
		AccountName:  "lionelritchie",
		Email:        "lionel.ritchie@yahoo.com",
		Password:     pass1,
		WalletAmount: 999999,
		Role:         5, //admin-master, can do anything
		SuspendedAt:  time.Time{},
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass2, _ := helpers.HashPassword("cindy123")
	db.Create(&UserAccount{
		AccountName:  "cindymirgiriti",
		Email:        "cindy.mirgiriti@yahoo.com",
		Password:     pass2,
		WalletAmount: 123456,
		Role:         1, //user-default, basic user
		SuspendedAt:  time.Time{},
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
}
