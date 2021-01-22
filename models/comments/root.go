package comments

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"time"
)

func Create(srcUserId, destUserId int, content string) models.UserComment {
	db := database.GetInstance()
	var comment models.UserComment
	comment = models.UserComment{
		SrcUserId:  uint(srcUserId),
		DestUserId: uint(destUserId),
		Content:    content,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	}
	db.Create(&comment)
	return comment
}
