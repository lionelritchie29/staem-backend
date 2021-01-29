package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lionelritchie29/staem-backend/database"
	"time"
)

type ChatMessage struct {
	ID 			uint
	SenderID     uint
	RecipientID uint
	Message string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&ChatMessage{})
	db.AutoMigrate(&ChatMessage{})

	p := &ChatMessage{}
	p.seed(db)
}

func (p *ChatMessage) seed(db *gorm.DB) {

}
