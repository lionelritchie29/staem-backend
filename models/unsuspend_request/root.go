package unsuspend_request

import (
	"fmt"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"time"
)

func GetAll() []models.UnsuspendRequest {
	db := database.GetInstance()
	var requests []models.UnsuspendRequest
	db.Find(&requests)
	return requests
}

func Create(userId int, reason string) bool {
	db := database.GetInstance()
	req := models.UnsuspendRequest{
		UserID:    uint(userId),
		Reason:    reason,
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}

	res := db.Create(&req)

	if res.Error != nil {
		return false
	}

	return true
}

func Delete(userId int) bool {
	db := database.GetInstance()
	res := db.Where("user_id = ?", userId).Unscoped().Delete(&models.UnsuspendRequest{})

	if res.Error != nil {
		fmt.Println(res)
		return false
	}

	return true
}

