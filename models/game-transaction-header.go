package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"syreclabs.com/go/faker"
	"time"
)

type GameTransactionHeader struct {
	ID uint `gorm:"primary_key"`
	PaymentMethod int
	User int
	CardNo string
	CardExp string
	BillingAddress string
	BillingCity string
	PostalCode string
	PhoneNumber string
	Country string
	TransactionDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameTransactionHeader{})
	db.AutoMigrate(&GameTransactionHeader{})

	p := &GameTransactionHeader{}
	p.seed(db)
}

func (p *GameTransactionHeader) seed(db *gorm.DB) {

	for i := 0; i < 100; i++ {
		startDate := time.Date(2020, time.August, 17, 12, 0, 0, 0, time.UTC)
		db.Create(&GameTransactionHeader{
			PaymentMethod:   rand.Intn(3) + 1,
			User:            rand.Intn(10) + 1,
			CardNo:          faker.Business().CreditCardNumber(),
			CardExp:         faker.Business().CreditCardExpiryDate(),
			BillingAddress:  faker.Address().String(),
			BillingCity:     faker.Address().City(),
			PostalCode:      faker.Address().Postcode(),
			PhoneNumber:     faker.PhoneNumber().PhoneNumber(),
			Country:         faker.Address().Country(),
			TransactionDate: faker.Date().Between(startDate, time.Now()),
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
			DeletedAt:       nil,
		})
	}
}
