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

func GetHelpfulReviewByGameId(gameId int) []models.GameReview {
	db := database.GetInstance()
	var reviews []models.GameReview
	db.Raw(
		"SELECT * FROM game_reviews WHERE game_id = " +
		strconv.FormatInt(int64(gameId), 10) + " " +
		"AND review_date_time::date >= CURRENT_DATE - INTERVAL '30 day' " +
		"ORDER BY upvote_count LIMIT 5").
		Scan(&reviews)
	return reviews
}

func GetRecentlyPostedbyGameId(gameId int) []models.GameReview {
	db := database.GetInstance()
	var reviews []models.GameReview
	db.Raw("SELECT * FROM game_reviews WHERE game_id = " +
		strconv.FormatInt(int64(gameId), 10) +
		" ORDER BY id DESC LIMIT 10").
		Scan(&reviews)
	return reviews
}