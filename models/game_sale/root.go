package game_sale

import (
	"fmt"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"time"
)

func GetAll() []models.GameSale{
	db := database.GetInstance()
	var sales []models.GameSale
	db.Find(&sales)
	return sales
}

func GetByGameId(id int) models.GameSale {
	db := database.GetInstance()
	var sale models.GameSale
	db.Where("game_id = ?", id).Find(&sale)
	return sale
}

func GetMoreThanFiftyPercentDiscount() []models.GameSale {
	db := database.GetInstance()
	var sales []models.GameSale
	db.Find(&sales, "discount >= 50")
	return sales
}

func Create(gameId, discount int, validTo string) models.GameSale {
	db := database.GetInstance()

	parsed, _ := time.Parse("2006-01-02", validTo)
	validToDateTime := time.Date(parsed.Year(), parsed.Month(), parsed.Day(), 12, 0, 0, 0, time.UTC)
	sale := models.GameSale{
		GameID:    uint(gameId),
		Discount:  discount,
		ValidTo:   validToDateTime,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}
	db.Create(&sale)

	return sale
}

func Update(gameId, discount int, validTo string) models.GameSale {
	db := database.GetInstance()

	parsed, _ := time.Parse("2006-01-02", validTo)
	validToDateTime := time.Date(parsed.Year(), parsed.Month(), parsed.Day(), 12, 0, 0, 0, time.UTC)

	var sale models.GameSale
	db.Where("game_id = ?", gameId).First(&sale)
	sale.Discount = discount
	sale.ValidTo = validToDateTime
	sale.UpdatedAt = time.Now()
	db.Save(&sale)

	return sale
}

func Delete(gameId int) bool {
	db := database.GetInstance()
	var sale models.GameSale
	fmt.Println(sale)
	res := db.Where("game_id = ?" , gameId).Unscoped().Delete(&sale)
	fmt.Println(sale)

	if res.Error != nil {
		fmt.Println(res.Error)
		return false
	}

	return true
}
