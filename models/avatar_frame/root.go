package avatar_frame

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
)

func GetByUserId(userId int) []models.AvatarFrame{
	db := database.GetInstance()
	var frames []models.AvatarFrame
	db.Raw("SELECT * FROM avatar_frames WHERE id IN ( " +
				"SELECT avatar_frame_id FROM user_avatar_frames " +
				"WHERE user_id = " + strconv.FormatInt(int64(userId), 10) +
			")").Scan(&frames)

	return frames
}