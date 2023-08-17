package models

import "time"

type Role struct {
	ID        uint     `gorm:"primaryKey"`
	Name      string   `gorm:"uniqueIndex;size:50;not null"`
	Users     []*User  `gorm:"many2many:user_roles"`
	Policies  []Policy `gorm:"many2many:role_policies"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
