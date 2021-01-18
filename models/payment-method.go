package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type PaymentMethod struct {
	ID uint `gorm:"primary_key"`
	Name string
	HasCard bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&PaymentMethod{})
	db.AutoMigrate(&PaymentMethod{})

	p := &PaymentMethod{}
	p.seed(db)
}

func (p *PaymentMethod) seed(db *gorm.DB) {
	db.Create(&PaymentMethod{
		Name:      "Visa",
		HasCard: true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&PaymentMethod{
		Name:      "Credit Card",
		HasCard: true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	db.Create(&PaymentMethod{
		Name:      "Miumiu",
		HasCard: false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})
}
