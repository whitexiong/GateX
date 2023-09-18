package models

import "time"

type Message struct {
	ID         uint   `gorm:"primaryKey"`
	ChatRoomID uint   `gorm:"index"`
	SenderID   uint   `gorm:"index"`
	ToUserID   *uint  `gorm:"index"`
	Content    string `gorm:"type:text;not null"`
	CreatedAt  time.Time
}
