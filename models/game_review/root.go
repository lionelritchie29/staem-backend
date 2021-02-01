package game_review

import (
	"fmt"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"strconv"
	"time"
)

func GetByGameId(gameId int) []models.GameReview {
	db := database.GetInstance()
	var reviews []models.GameReview
	db.Raw("SELECT * FROM game_reviews WHERE game_id = " + strconv.FormatInt(int64(gameId), 10)).
		Scan(&reviews)
	return reviews
}

func GetById(gameId int) models.GameReview {
	db := database.GetInstance()
	var review models.GameReview
	db.Find(&review, gameId)
	return review
}

func GetHelpfulReviewByGameId(gameId int) []models.GameReview {
	db := database.GetInstance()
	var reviews []models.GameReview
	db.Raw(
		"SELECT * FROM game_reviews WHERE game_id = " +
		strconv.FormatInt(int64(gameId), 10) + " " +
		"AND review_date_time::date >= CURRENT_DATE - INTERVAL '30 day' " +
		"ORDER BY upvote_count DESC LIMIT 5").
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

func GetAll() []models.GameReview{
	db := database.GetInstance()
	var reviews []models.GameReview
	db.Limit(20).Find(&reviews)
	return reviews
}


func Create(userId, gameId int, content string, isRecommended bool) bool {
	db := database.GetInstance()
	res := db.Create(&models.GameReview{
		IsRecommended:  isRecommended,
		GameID:         uint(gameId),
		UserID:         uint(userId),
		Content:        content,
		UpvoteCount:    0,
		DownvoteCount:  0,
		ReviewDateTime: time.Now(),
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      nil,
	})

	if res.Error != nil {
		return false
	}

	return true
}

func IncreaseUpvoteCount(id int) bool{
	db := database.GetInstance()
	var review models.GameReview
	getReview := db.Find(&review, id)
	review.UpvoteCount = review.UpvoteCount + 1
	saveReview := db.Save(&review)

	if getReview.Error != nil || saveReview.Error != nil {
		fmt.Println(getReview.Error)
		fmt.Print(saveReview.Error)
		return false
	}

	return true

}

func IncreaseDownvoteCount(id int) bool{
	db := database.GetInstance()
	var review models.GameReview
	getReview := db.Find(&review, id)
	review.DownvoteCount = review.DownvoteCount + 1
	saveReview := db.Save(&review)

	if getReview.Error != nil || saveReview.Error != nil {
		return false
	}

	return true
}

func GetComments(reviewId int) []models.ReviewComment {
	db := database.GetInstance()
	var comments []models.ReviewComment
	db.Where("post_id = ?", reviewId).Find(&comments, reviewId)
	return comments
}

func CreateComment(reviewId, userId int, comment string) models.ReviewComment{
	db := database.GetInstance()
	reviewComment := models.ReviewComment{
		PostID:    uint(reviewId),
		UserID:    uint(userId),
		Comments:  comment,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}

	res := db.Create(&reviewComment)

	if res.Error != nil {
		return models.ReviewComment{}
	}

	return reviewComment
}