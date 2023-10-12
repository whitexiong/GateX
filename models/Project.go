package models

import "time"

type Project struct {
	ID          uint                 `gorm:"primaryKey;autoIncrement"`
	Name        string               `gorm:"size:255;not null;uniqueIndex"`
	Description string               `gorm:"size:1024"`
	CoverImage  string               `gorm:"size:1024"`
	Category    string               `gorm:"size:255"`
	Tags        string               `gorm:"size:255"`
	OwnerID     uint                 `gorm:"foreignKey:UserID"`
	Sort        int                  `gorm:"default:1"`
	Status      int8                 `gorm:"type:tinyint"`
	APIs        []ProjectAPIEndpoint `gorm:"foreignKey:ProjectID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
