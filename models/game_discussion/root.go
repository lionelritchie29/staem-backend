package game_discussion

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
)

func Get(postId int) models.GameDiscussionPost {
	db := database.GetInstance()
	var post models.GameDiscussionPost
	db.Find(&post, postId)
	return post
}

func GetAll() []models.GameDiscussionPost{
	db := database.GetInstance()
	var posts []models.GameDiscussionPost
	db.Find(&posts)
	return posts
}

func GetCommentsByPostId(postId int) []models.GameDiscussionComment {
	db := database.GetInstance()
	var comments []models.GameDiscussionComment
	db.Where("post_id=?", postId).Find(&comments)
	return comments
}

func GetTopThreeByGameId(gameId int) []models.GameDiscussionPost {
	db := database.GetInstance()
	var posts []models.GameDiscussionPost
	db.Where("game_id = ?", gameId).Limit(3).Order("created_at DESC").Find(&posts)
	return posts
}