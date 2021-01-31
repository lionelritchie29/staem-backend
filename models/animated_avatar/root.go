package animated_avatar

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
	"time"
)

func GetAll() []models.AnimatedAvatar {
	db := database.GetInstance()
	var avatars []models.AnimatedAvatar
	db.Find(&avatars)
	return avatars
}

func GetByUserId(userId int) []models.AnimatedAvatar {
	db := database.GetInstance()
	var avatars []models.AnimatedAvatar
	db.Raw("SELECT * FROM animated_avatars WHERE id IN ( " +
		"SELECT animated_avatar_id FROM user_animated_avatars " +
		"WHERE user_id = " + strconv.FormatInt(int64(userId), 10) +
		")").Scan(&avatars)

	return avatars
}

func AddTransaction(userId, itemId int) bool {
	db := database.GetInstance()
	res := db.Create(&models.UserAnimatedAvatar{
		UserID:        uint(userId),
		AnimatedAvatarID: uint(itemId),
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	var profile models.UserProfile
	db.Find(&profile, userId)

	var item models.AnimatedAvatar
	db.Find(&item, itemId)

	profile.Point = profile.Point - item.Price

	db.Save(&profile)

	if res.Error != nil {
		return false
	}

	return true
}