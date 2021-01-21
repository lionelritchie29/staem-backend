package wallet

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
)

func Get(id int) models.WalletVoucher{
	db := database.GetInstance()
	var wallet models.WalletVoucher
	db.First(&wallet, id)
	return wallet
}

func Redeem(userId int, code string) int{
	db := database.GetInstance()
	var wallet models.WalletVoucher
	var user models.UserAccount

	res := db.Where("code = ? AND is_valid = true", code).First(&wallet)

	if res.RowsAffected > 0 {
		wallet.IsValid = false
		db.Save(&wallet)

		// update user wallet amount
		db.First(&user, userId)
		user.WalletAmount = user.WalletAmount + wallet.Amount
		db.Save(&user)

		return int(wallet.ID)
	}

	return -1
}
