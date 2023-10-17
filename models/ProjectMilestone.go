package models

import "time"

type ProjectMilestone struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	ProjectID uint
	Title     string `gorm:"size:255;not null"`
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
