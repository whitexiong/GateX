package models

import "time"

type ProjectSetting struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	ProjectID   uint   `gorm:"foreignKey:ProjectID"`
	IP          string `gorm:"size:255;not null"`
	Port        uint16
	Environment string `gorm:"size:255"`  // 例如: Development, Test, Production
	Description string `gorm:"size:1024"` // 其他描述性信息
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
