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
	AvatarUrl string     `gorm:"size:255"`
	Roles     []*Role    `gorm:"many2many:user_roles"`
	ParentID  *uint      `gorm:"index:idx_parent_id;comment:'上级用户'"`
	JwtTokens []JwtToken `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ChatUserResponse struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Avatar     string `json:"avatar"`
	ChatRoomID uint   `json:"chat_room_id,omitempty"`
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
