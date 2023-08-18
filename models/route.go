package models

import (
	"time"
)

type Route struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:100;not null"`
	Path       string `gorm:"size:255;not null"`
	Component  string `gorm:"size:255;not null"`
	ParentID   *uint  `gorm:"index:idx_parent_id"`
	Permission string `gorm:"size:100;not null"`
	Redirect   string `gorm:"size:255"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
