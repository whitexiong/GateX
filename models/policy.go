package models

import "time"

type Policy struct {
	ID          uint   `gorm:"primaryKey"`
	Sub         string `gorm:"size:100;not null"` // 对应Role的Name
	Obj         string `gorm:"size:100;not null"`
	Act         string `gorm:"size:100;not null"`
	Description string `gorm:"size:255"`
	Type        string `gorm:"type:enum('path', 'operation');not null"`
	Roles       []Role `gorm:"many2many:role_policies"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
