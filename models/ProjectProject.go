package models

import "time"

type ProjectProject struct {
	ID          uint                 `gorm:"primaryKey;autoIncrement"`
	Name        string               `gorm:"size:255;not null;uniqueIndex"`
	Description string               `gorm:"size:1024"`
	CoverImage  string               `gorm:"size:1024"`
	Category    string               `gorm:"size:255"`
	Tags        string               `gorm:"size:255"`
	OwnerID     uint                 `gorm:"foreignKey:UserID"`
	sort        int                  `gorm:"default:1"`
	APIs        []ProjectAPIEndpoint `gorm:"foreignKey:ProjectID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
