package game_sale

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
)

func GetAll() []models.GameSale{
	db := database.GetInstance()
	var sales []models.GameSale
	db.Find(&sales)
	return sales
}

func GetMoreThanFiftyPercentDiscount() []models.GameSale {
	db := database.GetInstance()
	var sales []models.GameSale
	db.Find(&sales, "discount >= 50")
	return sales
}
