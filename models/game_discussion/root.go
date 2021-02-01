package game_discussion

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"time"
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

func CreateComment(postId, userId int, content string) models.GameDiscussionComment{
	db := database.GetInstance()
	comment := models.GameDiscussionComment{
		PostID:    uint(postId),
		UserID:    uint(userId),
		Comments:  content,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}

	res := db.Create(&comment)

	if res.Error != nil {
		return models.GameDiscussionComment{}
	}

	return comment
}

func CreateDiscussion(userId, gameId int, title, content string) int{
	db := database.GetInstance()
	post := models.GameDiscussionPost{
		Title:     title,
		Content:   content,
		GameID:    uint(gameId),
		UserID:    uint(userId),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}

	res := db.Create(&post)

	if res.Error != nil {
		return -1
	}

	return int(post.ID)
}