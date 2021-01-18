package game_transaction

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
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