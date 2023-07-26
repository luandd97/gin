package entities

import "time"

type FacebookPageConversationMessage struct {
	ID             uint64    `gorm:"primary_key:auto_increment" json:"id"`
	_ID            string    `json:"_id"`
	PageID         string    `gorm:"index,not null" json:"page_id"`
	ConversationID string    `json:"conversation_id"`
	MessageID      string    `json:"message_id"`
	Sender         string    `json:"sender"`
	AttachmentUrl  string    `gorm:"null" json:"attachment_url"`
	AttachmentName string    `gorm:"null" json:"attachment_name"`
	AttachmentID   string    `gorm:"null" json:"attachment_id"`
	AttachmentType string    `gorm:"null" json:"attachment_type"`
	AttachmentSize uint64    `gorm:"null" json:"attachment_size"`
	Message        string    `gorm:"not null" json:"message"`
	From           string    `gorm:"not null" json:"from"`
	To             string    `gorm:"not null" json:"to"`
	CreatedTime    time.Time `json:"created_time"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
