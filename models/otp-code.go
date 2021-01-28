package models

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/helpers"
	"math/rand"
	"time"
)

type OTPCode struct {
	ID uint `gorm:"primary_key"`
	Code string
	IsValid bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&OTPCode{})
	db.AutoMigrate(&OTPCode{})
}

func CreateOTP(email, accountName string) bool{
	db := database.GetInstance()
	code := randCode()
	res := db.Create(&OTPCode{
		Code:      code,
		IsValid:   true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	helpers.SendOTPMail(email, accountName, code)

	if res.Error != nil {
		return false
	}

	return true
}

func VerifiyOTP(code string) bool {
	db := database.GetInstance()
	var otp OTPCode

	res := db.Where("code = ? and is_valid = true", code).Find(&otp)
	otp.IsValid = false
	db.Save(&otp)

	if res.RowsAffected == 0 {
		return false
	}

	return true
}

func randCode() string {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 5)
	for i := 0; i<5; i++ {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}

