package models

import "time"

type Role struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"uniqueIndex;size:50;not null"`
	Status    int8    `gorm:"type:tinyint"`
	Remark    string  `gorm:"type:varchar(255)"`
	Users     []*User `gorm:"many2many:user_roles"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoleRequest struct {
	Role
	Permissions []uint
}
