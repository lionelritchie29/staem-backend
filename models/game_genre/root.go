package game_genre

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
)

func GetAll() []models.GameGenre {
	db := database.GetInstance()
	var genres []models.GameGenre
	db.Find(&genres)
	return genres
}
