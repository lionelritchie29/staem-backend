package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/helpers"
	"syreclabs.com/go/faker"
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
	Status		 string
	Code		 string
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
		Status: "offline",
		Code: faker.Number().Number(9),
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
		Code: faker.Number().Number(9),
		Role:         1, //user-default, basic user
		SuspendedAt:  time.Time{},
		Status: "offline",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass3, _ := helpers.HashPassword("arthur123")
	db.Create(&UserAccount{
		AccountName:  "arthurmorgan",
		Email:        "arthur.morgan@yahoo.com",
		Password:     pass3,
		WalletAmount: 321000,
		Role:         1, //user-default, basic user
		SuspendedAt:  time.Time{},
		Status: "offline",
		Code: faker.Number().Number(9),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass4, _ := helpers.HashPassword("abigail123")
	db.Create(&UserAccount{
		AccountName:  "abigailroberts",
		Email:        "abigail.roberts@yahoo.com",
		Password:     pass4,
		WalletAmount: 555000,
		Role:         1, //user-default, basic user
		SuspendedAt:  time.Time{},
		Status: "offline",
		Code: faker.Number().Number(9),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass5, _ := helpers.HashPassword("aloy123")
	db.Create(&UserAccount{
		AccountName:  "aloy",
		Email:        "aloy@yahoo.com",
		Password:     pass5,
		WalletAmount: 777777,
		Role:         1, //user-default, basic user
		SuspendedAt:  time.Now(),
		Status: "offline",
		Code: faker.Number().Number(9),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass6, _ := helpers.HashPassword("peter123")
	db.Create(&UserAccount{
		AccountName:  "peterparker",
		Email:        "peter.parker@yahoo.com",
		Password:     pass6,
		WalletAmount: 888888,
		Role:         1, //user-default, basic user
		SuspendedAt:  time.Time{},
		Status: "offline",
		Code: faker.Number().Number(9),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass7, _ := helpers.HashPassword("hosea123")
	db.Create(&UserAccount{
		AccountName:  "hoseamathhews",
		Email:        "hosea.matthews@yahoo.com",
		Password:     pass7,
		WalletAmount: 888888,
		Role:         1, //user-default, basic user
		SuspendedAt:  time.Time{},
		Status: "offline",
		Code: faker.Number().Number(9),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass8, _ := helpers.HashPassword("patrick123")
	db.Create(&UserAccount{
		AccountName:  "patrickstar",
		Email:        "patrick.star@yahoo.com",
		Password:     pass8,
		WalletAmount: 888888,
		Role:         1, //user-default, basic user
		SuspendedAt:  time.Time{},
		Status: "offline",
		Code: faker.Number().Number(9),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass9, _ := helpers.HashPassword("spongebob123")
	db.Create(&UserAccount{
		AccountName:  "spongebobsquarepants",
		Email:        "spongebob.squarepants@yahoo.com",
		Password:     pass9,
		WalletAmount: 134679,
		Role:         1, //user-default, basic user
		SuspendedAt:  time.Time{},
		Status: "offline",
		Code: faker.Number().Number(9),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass10, _ := helpers.HashPassword("sadie123")
	db.Create(&UserAccount{
		AccountName:  "sadieadler",
		Email:        "sadie.adler@yahoo.com",
		Password:     pass10,
		WalletAmount: 159789,
		Role:         1, //user-default, basic user
		SuspendedAt:  time.Time{},
		Status: "offline",
		Code: faker.Number().Number(9),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass11, _ := helpers.HashPassword("admin-user")
	db.Create(&UserAccount{
		AccountName:  "admin-user",
		Email:        "admin-user@yahoo.com",
		Password:     pass11,
		WalletAmount: -1,
		Role:         4, //admin-user
		SuspendedAt:  time.Time{},
		Status: "offline",
		Code: faker.Number().Number(9),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass12, _ := helpers.HashPassword("admin-promo")
	db.Create(&UserAccount{
		AccountName:  "admin-promo",
		Email:        "admin-promo@yahoo.com",
		Password:     pass12,
		WalletAmount: -1,
		Role:         3, //admin-promo
		SuspendedAt:  time.Time{},
		Status: "offline",
		Code: faker.Number().Number(9),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})

	pass13, _ := helpers.HashPassword("admin-game")
	db.Create(&UserAccount{
		AccountName:  "admin-game",
		Email:        "admin-game@yahoo.com",
		Password:     pass13,
		WalletAmount: -1,
		Role:         2, //admin-game
		SuspendedAt:  time.Time{},
		Status: "offline",
		Code: faker.Number().Number(9),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	})
}
