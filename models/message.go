package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SenderID   uint `json:"sender_id"`
	ReceiverID uint `json:"receiver_id"`
	Content    string
}

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string

	MessagesSent     []Message `gorm:"foreignKey:SenderID"`
	MessagesReceived []Message `gorm:"foreignKey:ReceiverID"`
}
