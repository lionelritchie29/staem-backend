package mini_profile

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
)

func GetByUserId(userId int) []models.ProfileBackground{
	db := database.GetInstance()
	var bg []models.ProfileBackground
	db.Raw("SELECT * FROM mini_profile_backgrounds WHERE id IN ( " +
		"SELECT mini_profile_background_id FROM user_mini_profile_backgrounds " +
		"WHERE user_id = " + strconv.FormatInt(int64(userId), 10) +
		")").Scan(&bg)

	return bg
}
