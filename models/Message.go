package models

import "time"

type MessageType int8

const (
	HumanToOne   MessageType = iota // 人对人一对一聊天
	HumanToGroup                    // 人对群组聊天
	AIToOne                         // AI对人一对一聊天
)

type AIProviderType int8

const (
	None   AIProviderType = iota // 由人类发送，不是AI
	XunFei                       // 星火大模型
	Baidu                        // 文心一言
)

type Message struct {
	ID         uint   `gorm:"primaryKey"`
	ChatRoomID uint   `gorm:"index"`
	SenderID   uint   `gorm:"index"`
	Content    string `gorm:"type:text;not null"`
	CreatedAt  time.Time
	Type       MessageType    `gorm:"type:tinyint;not null"`
	AIProvider AIProviderType `gorm:"type:tinyint;not null;default:0"` // 默认为None，即由人类发送
}
