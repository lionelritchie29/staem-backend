package mini_profile

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
	"time"
)

func GetAll() []models.MiniProfileBackground {
	db := database.GetInstance()
	var bg []models.MiniProfileBackground
	db.Find(&bg)
	return bg
}

func GetByUserId(userId int) []models.ProfileBackground{
	db := database.GetInstance()
	var bg []models.ProfileBackground
	db.Raw("SELECT * FROM mini_profile_backgrounds WHERE id IN ( " +
		"SELECT mini_profile_background_id FROM user_mini_profile_backgrounds " +
		"WHERE user_id = " + strconv.FormatInt(int64(userId), 10) +
		")").Scan(&bg)

	return bg
}

func AddTransaction(userId, itemId int) bool{
	db := database.GetInstance()
	res := db.Create(&models.UserMiniProfileBackground{
		UserID:              uint(userId),
		MiniProfileBackgroundID: uint(itemId),
		CreatedAt:           time.Time{},
		UpdatedAt:           time.Time{},
		DeletedAt:           nil,
	})

	var profile models.UserProfile
	db.Find(&profile, userId)

	var item models.MiniProfileBackground
	db.Find(&item, itemId)

	profile.Point = profile.Point - item.Price

	db.Save(&profile)

	if res.Error != nil {
		return false
	}

	return true
}