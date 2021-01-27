package user_profile

import (
	"fmt"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/input_models"
	"github.com/lionelritchie29/staem-backend/models"
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
