package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
"github.com/lionelritchie29/staem-backend/database"
	"math/rand"
	"time"
)

type WalletVoucher struct {
	ID uint `gorm:"primary_key"`
	Code string
	Amount int
	IsValid bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&WalletVoucher{})
	db.AutoMigrate(&WalletVoucher{})

	p := &WalletVoucher{}
	p.seed(db)
}

func (p *WalletVoucher) seed(db *gorm.DB) {

	for i:=0; i<10; i++ {
		boolInt := rand.Intn(2) + 1

		var isValid bool
		if boolInt == 1{
			isValid = true
		} else {
			isValid = false
		}

		db.Create(&WalletVoucher{
			Code: uuid.Must(uuid.NewRandom()).String(),
			Amount: rand.Intn(2000000) + 10000,
			IsValid: isValid,
			UpdatedAt:  time.Time{},
			DeletedAt:  nil,
		})
	}
}
