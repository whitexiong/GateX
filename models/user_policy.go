package models

type UserPolicy struct {
	UserID   uint `gorm:"primaryKey"`
	PolicyID uint `gorm:"primaryKey"`
}
