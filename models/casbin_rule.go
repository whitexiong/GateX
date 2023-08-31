package models

import (
	"time"
)

type CasbinRule struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`              // 主键
	PType     string `gorm:"column:p_type;size:100;index;not null"` // 策略类型，例如"p"、"g"
	V0        string `gorm:"column:v0;size:100;index"`              // 通常是主体(subject)
	V1        string `gorm:"column:v1;size:100;index"`              // 通常是对象(object)或角色名
	V2        string `gorm:"column:v2;size:100;index"`              // 通常是操作(action)
	V3        string `gorm:"column:v3;size:100"`                    // 可选
	V4        string `gorm:"column:v4;size:100"`                    // 可选
	V5        string `gorm:"column:v5;size:100"`                    // 可选
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
