package models

import (
	"time"
)

type Route struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null;comment:'路由名称'"`
	Path      string    `gorm:"size:255;not null;comment:'路由路径'"`
	ParentID  *uint     `gorm:"index:idx_parent_id;comment:'父路由ID'"`
	Status    int8      `gorm:"type:tinyint"`
	CreatedAt time.Time `gorm:"comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"comment:'更新时间'"`
}
