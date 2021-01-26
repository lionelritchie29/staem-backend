package user_wishlist

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"time"
)

func GetAll() []models.UserWishlist {
	db := database.GetInstance()
	var wishlist []models.UserWishlist
	db.Find(&wishlist)
	return wishlist
}

func GetById(userId int) []models.UserWishlist {
	db := database.GetInstance()
	var wishlist []models.UserWishlist
	db.Where("user_id = ?", userId).Find(&wishlist)
	return wishlist
}

func Create(userId, gameId int) models.UserWishlist {
	db := database.GetInstance()
	wishlist := models.UserWishlist{
		UserID:    uint(userId),
		GameID:    uint(gameId),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}

	db.Create(&wishlist)
	return wishlist
}

func Delete(userId, gameId int) bool {
	db := database.GetInstance()
	res := db.Where("user_id = ? and game_id = ?", userId, gameId).Unscoped().Delete(&models.UserWishlist{})

	if res.Error != nil {
		return false
	}

	return true
}
