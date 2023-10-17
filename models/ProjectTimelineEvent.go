package models

import "time"

type ProjectTimelineEvent struct {
	ID                    uint `gorm:"primaryKey;autoIncrement"`
	ProjectID             uint
	Description           string `gorm:"size:1024"`
	EventType             string `gorm:"size:255"`
	EventDate             time.Time
	AssociatedMilestoneID uint `gorm:"foreignKey:MilestoneID"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
