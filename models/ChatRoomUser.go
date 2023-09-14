package models

import "time"

type ChatRoomUser struct {
	ID         uint `gorm:"primaryKey"`
	ChatRoomID uint
	UserID     uint
	CreatedAt  time.Time
}
