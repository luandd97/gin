package dto

import "time"

type FacebookPageDTO struct {
	UserID      uint64
	PageID      string
	AccessToken string
	Name        string
	Picture     string
}

type FacebookPageConversationDTO struct {
	ID             uint64
	UserID         uint64
	PageID         string
	ConversationID string
	InboxUrl       string
	LastMessage    string
	Sender         string
	IsSync         bool
	UnreadCount    uint
	MessageCount   uint
	UpdatedTime    time.Time
}
