package entities

import "time"

type FacebookPageConversation struct {
	ID             uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserID         uint64    `gorm:"index,not null" json:"user_id"`
	PageID         string    `gorm:"index,not null" json:"page_id"`
	ConversationID string    `json:"conversation_id"`
	InboxUrl       string    `json:"inbox_url"`
	IsSync         bool      `gorm:"default(0)" json:"is_sync"`
	LastMessage    string    `gorm:"type:text" json:"last_message"`
	Sender         string    `json:"sender"`
	UnreadCount    uint      `json:"unread_count"`
	UpdatedTime    time.Time `json:"updated_time"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	MessageCount   uint      `gorm:"default(0)" json:"message_count"`
}
