package models

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey"`
	Username  string     `gorm:"uniqueIndex;size:100;not null"`
	Password  string     `gorm:"size:255;not null"`
	Email     string     `gorm:"size:100"`
	Roles     []*Role    `gorm:"many2many:user_roles"`
	JwtTokens []JwtToken `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
