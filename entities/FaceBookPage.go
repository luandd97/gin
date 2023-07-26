package entities

import "time"

type FacebookPage struct {
	ID                 uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserID             uint64    `gorm:"index,not null" json:"user_id"`
	PageID             string    `gorm:"index,not null" json:"page_id"`
	AccessToken        string    `json:"_"`
	Name               string    `json:"name"`
	Picture            string    `json:"picture"`
	IsSyncConversation bool      `gorm:"default(0)" json:"is_sync_conversation"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
