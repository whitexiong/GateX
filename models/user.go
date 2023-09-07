package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey"`
	Username  string     `gorm:"uniqueIndex;size:100;not null"`
	Password  string     `gorm:"size:255;not null"`
	Email     string     `gorm:"size:100"`
	Status    int8       `gorm:"type:tinyint"`
	Roles     []*Role    `gorm:"many2many:user_roles"`
	JwtTokens []JwtToken `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindUserByUsername(username string) (*User, error) {
	var user User
	if err := DB.Preload("Roles").Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

type UserRequest struct {
	User
	Roles []uint `json:"Roles"`
}
