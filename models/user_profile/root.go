package user_profile

import (
	"fmt"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/helpers"
	"github.com/lionelritchie29/staem-backend/input_models"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
)

func UpdateInfo(newInfo input_models.NewProfileInfo) bool {
	db := database.GetInstance()
	var profile models.UserProfile
	db.Find(&profile, newInfo.UserID)

	profile.DisplayName = newInfo.DisplayName
	profile.RealName = newInfo.RealName
	profile.CustomURL = newInfo.CustomUrl
	profile.Country = newInfo.Country
	profile.Summary = newInfo.Summary
	saveProfile := db.Save(&profile)

	if saveProfile.Error != nil {
		fmt.Println(saveProfile.Error)
		return false
	}

	return true
}

func UpdateAvatar(userId int, encodedBase64Image string) bool {
	db := database.GetInstance()
	var profile models.UserProfile
	db.Find(&profile, userId)

	filename := strconv.FormatInt(int64(userId), 10) + ".jpg"
	profile.ProfilePictureUrl = filename
	saveProfile := db.Save(&profile)
	helpers.SaveImage(encodedBase64Image, "avatars/" + filename)

	if saveProfile.Error != nil {
		return false
	}

	return true
}

func UpdateAvatarFrame(userId int, avatarFrameUrl string) bool {
	db := database.GetInstance()
	var profile models.UserProfile
	db.Find(&profile, userId)
	profile.AvatarFrameUrl = avatarFrameUrl
	saveProfile := db.Save(&profile)

	if saveProfile.Error != nil {
		return false
	}

	return true
}

func UpdateBackground(userId int, profileBackgroundUrl string) bool {
	db := database.GetInstance()
	var profile models.UserProfile
	db.Find(&profile, userId)
	profile.ProfileBackgroundUrl = profileBackgroundUrl
	saveProfile := db.Save(&profile)

	if saveProfile.Error != nil {
		return false
	}

	return true
}

func UpdateMiniBackground(userId int, miniProfileBackgroundUrl string) bool {
	db := database.GetInstance()
	var profile models.UserProfile
	db.Find(&profile, userId)
	profile.MiniProfileBackgroundUrl = miniProfileBackgroundUrl
	saveProfile := db.Save(&profile)

	if saveProfile.Error != nil {
		return false
	}

	return true
}

func UpdateTheme(userId int, hexCode string) bool {
	db := database.GetInstance()
	var profile models.UserProfile
	db.Find(&profile, userId)
	profile.Theme = hexCode
	saveProfile := db.Save(&profile)

	if saveProfile.Error != nil {
		return false
	}

	return true
}