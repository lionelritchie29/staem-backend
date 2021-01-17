package game_review

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
)

func GetByGameId(gameId int) []models.GameReview {
	db := database.GetInstance()
	var reviews []models.GameReview
	db.Raw("SELECT * FROM game_reviews WHERE game_id = " + strconv.FormatInt(int64(gameId), 10)).
		Scan(&reviews)
	return reviews
}
