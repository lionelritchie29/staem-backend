package friend_request

import (
	"fmt"
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"time"
)

func GetReceivedByUserId(userId int) []models.ReceivedFriendRequest{
	db := database.GetInstance()
	var req []models.ReceivedFriendRequest
	db.Where("user_id = ?", userId).Find(&req)
	return req
}

func GetSentByUserId(userId int) []models.SentFriendRequest{
	db := database.GetInstance()
	var req []models.SentFriendRequest
	db.Where("user_id = ?", userId).Find(&req)
	return req
}

func Create(fromId, toId int) bool {
	db := database.GetInstance()

	received := db.Create(&models.SentFriendRequest{
		UserID:    uint(fromId),
		FriendID:  uint(toId),
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	sent := db.Create(&models.ReceivedFriendRequest{
		UserID:    uint(toId),
		FriendID:  uint(fromId),
		Status:    "pending",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	if sent.Error != nil || received.Error != nil {
		return false
	}

	return true
}

func Accept(fromId, toId int) bool {
	db := database.GetInstance()
	var received models.ReceivedFriendRequest
	var sent models.SentFriendRequest

	deleteReceived := db.Where("user_id = ? AND friend_id = ?", fromId, toId).Delete(&received)
	deleteSent := db.Where("user_id = ? AND friend_id = ?", toId, fromId).Delete(&sent)

	createFriend := db.Create(&models.Friend{
		UserID:    uint(fromId),
		FriendID:  uint(toId),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	createFriendReverse := db.Create(&models.Friend{
		UserID:    uint(toId),
		FriendID:  uint(fromId),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	if deleteReceived.Error != nil || deleteSent.Error != nil || createFriend.Error != nil || createFriendReverse.Error != nil {
		fmt.Println(deleteReceived)
		fmt.Println(deleteSent)
		fmt.Println(createFriend)
		fmt.Println(createFriendReverse)
		return false
	}

	return true
}

func Reject(fromId, toId int) bool {
	db := database.GetInstance()
	var received models.ReceivedFriendRequest
	var sent models.SentFriendRequest

	deleteReceived := db.Where("user_id = ? AND friend_id = ?", fromId, toId).Unscoped().Delete(&received)
	deleteSent := db.Where("user_id = ? AND friend_id = ?", toId, fromId).Unscoped().Delete(&sent)

	if deleteReceived.Error != nil || deleteSent.Error != nil  {
		return false
	}

	return true
}

func Ignore(fromId, toId int) bool {
	db := database.GetInstance()
	var received models.ReceivedFriendRequest

	deleteReceived := db.Where("user_id = ? AND friend_id = ?", fromId, toId).Delete(&received)

	if deleteReceived.Error != nil {
		return false
	}

	return true
}