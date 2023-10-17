package models

import "time"

type ProjectFile struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	ProjectID  uint
	FileName   string `gorm:"size:255;not null"`
	FilePath   string `gorm:"size:1024;not null"`
	FileType   string `gorm:"size:50"`
	Version    string `gorm:"size:50"`
	UploadedBy uint   `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
