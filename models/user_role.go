package models

import "time"

type UserRole struct {
	UserID    uint
	RoleID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
