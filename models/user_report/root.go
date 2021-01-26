package user_report

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"github.com/lionelritchie29/staem-backend/models/user"
	"time"
)

func GetByUserId(userId int) []models.UserReport {
	db := database.GetInstance()
	var reports []models.UserReport
	db.Where("user_id = ?", userId).Find(&reports)
	return reports
}

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
	} else {
		var userIds []int
		checkCount := db.Raw("SELECT user_id FROM user_reports " +
								"WHERE created_at >= CURRENT_DATE - INTERVAL '7 day' " +
			"					GROUP BY user_id HAVING COUNT(user_id) >= 5").Scan(&userIds)
		if checkCount.RowsAffected == 0 {
			return true
		} else { //user has been reported more than equals to 5 in last week
			user.Suspend(targetId)
		}
	}

	return true
}

func Delete(targetId int) bool {
	db := database.GetInstance()
	res := db.Where("user_id = ?", targetId).Unscoped().Delete(&models.UserReport{})
	if res != nil {
		return false
	}

	return true
}
