package chat_sticker

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
	"time"
)

func GetAll() []models.ChatSticker {
	db := database.GetInstance()
	var stickers []models.ChatSticker
	db.Find(&stickers)
	return stickers
}

func GetByUserId(userId int) []models.ChatSticker {
	db := database.GetInstance()
	var stickers []models.ChatSticker
	db.Raw("SELECT * FROM chat_stickers WHERE id IN ( " +
		"SELECT chat_sticker_id FROM user_chat_stickers " +
		"WHERE user_id = " + strconv.FormatInt(int64(userId), 10) +
		")").Scan(&stickers)

	return stickers
}

func AddTransaction(userId, itemId int) bool {
	db := database.GetInstance()
	res := db.Create(&models.UserChatSticker{
		UserID:        uint(userId),
		ChatStickerID: uint(itemId),
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
	})

	var profile models.UserProfile
	db.Find(&profile, userId)

	var item models.ChatSticker
	db.Find(&item, itemId)

	profile.Point = profile.Point - item.Price

	db.Save(&profile)

	if res.Error != nil {
		return false
	}

	return true
}