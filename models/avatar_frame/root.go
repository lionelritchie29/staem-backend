package avatar_frame

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
	"time"
)

func GetAll() []models.AvatarFrame {
	db := database.GetInstance()
	var frames []models.AvatarFrame
	db.Find(&frames)
	return frames
}

func GetByUserId(userId int) []models.AvatarFrame {
	db := database.GetInstance()
	var frames []models.AvatarFrame
	db.Raw("SELECT * FROM avatar_frames WHERE id IN ( " +
		"SELECT avatar_frame_id FROM user_avatar_frames " +
		"WHERE user_id = " + strconv.FormatInt(int64(userId), 10) +
		")").Scan(&frames)

	return frames
}

func AddTransaction(userId, itemId int) bool {
	db := database.GetInstance()
	res := db.Create(&models.UserAvatarFrame{
		UserID:        uint(userId),
		AvatarFrameID: uint(itemId),
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	var profile models.UserProfile
	db.Find(&profile, userId)

	var item models.AvatarFrame
	db.Find(&item, itemId)

	profile.Point = profile.Point - item.Price

	db.Save(&profile)

	if res.Error != nil {
		return false
	}

	return true
}
