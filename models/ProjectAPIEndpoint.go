package models

import "time"

type ProjectAPIEndpoint struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	ProjectID     uint
	Name          string `gorm:"size:255;not null"`
	Description   string `gorm:"size:1024"`
	Method        string `gorm:"size:10"`
	Path          string `gorm:"size:1024;not null"`
	Documentation string `gorm:"size:2048"`
	Parameters    string `gorm:"size:2048"`
	Responses     string `gorm:"size:2048"`
	RequiredRoles string `gorm:"size:255"`
	Version       string `gorm:"size:50"`
	Environment   string `gorm:"size:50"`
	Sort          int    `gorm:"default:1"`
	Status        int8   `gorm:"type:tinyint"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
