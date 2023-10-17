package models

import "time"

type ProjectTask struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	ProjectID   uint
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"size:1024"`
	AssigneeID  uint   `gorm:"foreignKey:UserID"`
	Status      int8   `gorm:"type:tinyint"`
	Priority    int
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
