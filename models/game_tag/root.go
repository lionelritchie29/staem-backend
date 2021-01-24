package game_tag

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
)

func GetAll() []models.GameTag {
	db := database.GetInstance()
	var tags []models.GameTag
	db.Find(&tags)
	return tags
}
