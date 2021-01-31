package game_transaction

import (
	"fmt"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/input_models"
	"github.com/lionelritchie29/staem-backend/models"
	"time"
)

func GetAll() []models.GameTransactionHeader {
	db := database.GetInstance()
	var transactions []models.GameTransactionHeader
	db.Find(&transactions)
	return transactions
}

func GetDetails(transactionId int) []models.GameTransactionDetail {
	db := database.GetInstance()
	var details []models.GameTransactionDetail
	db.Where("game_transaction_id = ?", transactionId).Find(&details)
	return details
}

func GetPaymentMethodById(paymentId int) models.PaymentMethod {
	db := database.GetInstance()
	var paymentMethod models.PaymentMethod
	db.Find(&paymentMethod, paymentId)
	return paymentMethod
}

func AddTransaction(newTransaction input_models.NewGameTransaction) bool {
	db := database.GetInstance()

	header := &models.GameTransactionHeader{
		PaymentMethod:   newTransaction.PaymentMethod,
		User:            newTransaction.User,
		CardNo:          newTransaction.CardNo,
		CardExp:         newTransaction.CardExp,
		BillingAddress:  newTransaction.BillingAddress,
		BillingCity:     newTransaction.BillingCity,
		PostalCode:      newTransaction.PostalCode,
		PhoneNumber:     newTransaction.PhoneNumber,
		Country:         newTransaction.Country,
		TransactionDate: time.Now(),
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	}

	resHeader := db.Create(header)

	if resHeader.Error != nil {
		panic(resHeader.Error)
		return false
	}

	var theUser models.UserAccount
	db.Find(&theUser, newTransaction.User)

	for _, detail := range newTransaction.Details {
		resDetail := db.Create(&models.GameTransactionDetail{
			GameTransactionID: header.ID,
			Game:              detail.Game,
			Price: 			   detail.Price,
			Quantity:          detail.Quantity,
			CreatedAt:         time.Time{},
			UpdatedAt:         time.Time{},
			DeletedAt:         nil,
		})

		theUser.WalletAmount = theUser.WalletAmount - detail.Price

		createUserGame := db.Create(&models.UserGame{
			UserID:    uint(newTransaction.User),
			GameID:    uint(detail.Game),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		})

		db.Save(&theUser)


		if resDetail.Error != nil || createUserGame.Error != nil {
			fmt.Println(resDetail.Error)
			return false
		}
	}



	return true
}