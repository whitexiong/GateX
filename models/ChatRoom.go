package models

import "time"

type ChatRoom struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255"`
	Description string    `gorm:"size:500"`
	Users       []*User   `gorm:"many2many:chatroom_users"`
	Messages    []Message `gorm:"foreignKey:ChatRoomID"`
	IsPrivate   bool      `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
