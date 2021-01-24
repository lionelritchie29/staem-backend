package user_report

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"time"
)

func Create(targetId, reporterId int, reason string) bool {
	db := database.GetInstance()
	report := models.UserReport{
		UserID:     uint(targetId),
		ReporterID: uint(reporterId),
		Reason: reason,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	}

	res := db.Create(&report)

	if res.Error != nil {
		return false
	}

	return true
}
