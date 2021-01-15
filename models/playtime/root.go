package playtime

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
)

func GetAll() []models.GamePlaytime {
	db := database.GetInstance()
	var playtimes []models.GamePlaytime
	db.Find(&playtimes)
	return playtimes
}

func GetAllByGameId(gameId int) []models.GamePlaytime {
	db := database.GetInstance()
	var playtimes []models.GamePlaytime
	db.Find(&playtimes, "game_id = ?", gameId)
	return playtimes
}