package models

import "time"

type JwtToken struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"not null"`
	Token      string `gorm:"size:255;not null"`
	ExpiryDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
