package chat_message

import (
	"github.com/lionelritchie29/staem-backend/database"
	"github.com/lionelritchie29/staem-backend/models"
	"time"
)

func GetChats(senderId, recipientId int) []models.ChatMessage{
	db := database.GetInstance()
	var chats []models.ChatMessage
	db.Where("(sender_id = ? AND recipient_id = ?) OR ((sender_id = ? AND recipient_id = ?))", senderId, recipientId, recipientId, senderId).Find(&chats)
	return chats
}

func GetLatestChat(senderId, recipientId int) models.ChatMessage{
	db := database.GetInstance()
	var chat models.ChatMessage
	db.Where("(sender_id = ? AND recipient_id = ?) OR ((sender_id = ? AND recipient_id = ?))", senderId, recipientId, recipientId, senderId).Last(&chat)
	return chat
}

func Add(senderId, recipientId  int, message string) bool{
	db := database.GetInstance()
	res := db.Create(&models.ChatMessage{
		SenderID:    uint(senderId),
		RecipientID:  uint(recipientId),
		Message:   message,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	})

	if res.Error != nil { return false }

	return true
}
