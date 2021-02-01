package image_video_post

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"time"
)

func GetAll() []models.ImageVideoPost {
	db := database.GetInstance()
	var posts []models.ImageVideoPost
	db.Find(&posts)
	return posts
}

func GetById(id int) models.ImageVideoPost {
	db := database.GetInstance()
	var post models.ImageVideoPost
	db.Find(&post, id)
	return post
}

func GetComments(postId int) []models.ImageVideoComment {
	db := database.GetInstance()
	var comments []models.ImageVideoComment
	db.Where("post_id = ?", postId).Find(&comments)
	return comments
}

func AddComments(postId, userId int, comment string) bool {
	db := database.GetInstance()
	isSuccess := db.Create(&models.ImageVideoComment{
		PostID:    uint(postId),
		UserID:    uint(userId),
		Comments:  comment,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	return isSuccess.Error == nil
}

func AddLike(postId int) bool {{
	db := database.GetInstance()
	var post models.ImageVideoPost
	db.Find(&post, postId)

	post.LikeCount = post.LikeCount + 1
	res := db.Save(&post)

	return res.Error == nil
}}

func AddDislike(postId int) bool {{
	db := database.GetInstance()
	var post models.ImageVideoPost
	db.Find(&post, postId)

	post.LikeCount = post.DislikeCount + 1
	res := db.Save(&post)

	return res.Error == nil
}}