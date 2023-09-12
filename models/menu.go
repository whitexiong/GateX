package models

import (
	"time"
)

type Menu struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Path      string `gorm:"size:255;not null"`
	ParentID  *uint  `gorm:"index:idx_parent_id"`
	Icon      string `gorm:"size:50"`
	Order     int    `gorm:"default:999"`
	Status    int8   `gorm:"type:tinyint"`
	RouteID   uint   `gorm:"index:idx_route_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m Menu) GetID() uint {
	return m.ID
}

func (m Menu) GetParentID() *uint {
	return m.ParentID
}
