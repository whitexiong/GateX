package models

import "time"

type ChatRoomUser struct {
	ChatRoomID uint
	UserID     uint
	CreatedAt  time.Time
}
