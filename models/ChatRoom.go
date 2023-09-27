package models

import "time"

const (
	NormalChatRoom = 1
	AIChatRoom     = 2
)

type ChatRoom struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255"`
	Description string    `gorm:"size:500"`
	Users       []*User   `gorm:"many2many:chat_room_users"`
	Messages    []Message `gorm:"foreignKey:ChatRoomID"`
	RoomType    int       `gorm:"default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ChatRoomRequest struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	RoomType    int    `json:"RoomType"`
	UserIDs     []uint `json:"UserIDs"`
}
