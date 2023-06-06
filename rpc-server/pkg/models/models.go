package models

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/pkg/config"
)

var db *gorm.DB

type Message struct {
	gorm.Model
	Chat     string `gorm:""json:"chat"`
	Text     string `json:"text"`
	Sender   string `json:"sender"`
	SendTime uint64 `json:"send_time"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Message{})
}

func (message *Message) CreateMessage() (*Message, error) {
	err := db.Create(message).Error
	return message, err
}

func GetMessages(chat string, cursor string, reverse bool, limit int) ([]Message, *gorm.DB) {
	var messages []Message

	var order string
	if reverse {
		order = "DESC"
	} else {
		order = "ASC"
	}

	order = fmt.Sprintf("send_time %s", order)

	var db *gorm.DB = db.Where("chat = ? AND send_time >= ?", chat, cursor).Order(order).Limit(limit).Find(&messages)
	return messages, db
}
