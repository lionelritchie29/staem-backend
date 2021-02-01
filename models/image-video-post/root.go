package image_video_post

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
)

func GetAll() []models.ImageVideoPost {
	db := database.GetInstance()
	var posts []models.ImageVideoPost
	db.Find(&posts)
	return posts
}
