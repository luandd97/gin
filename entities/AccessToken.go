package entities

import "time"

type OauthAccessToken struct {
	ID             uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserID         uint64    `gorm:"index,not null" json:"user_id"`
	AccessToken    string    `json:"access_token"`
	FacebookUserId string    `json:"facebook_user_id"`
	Name           string    `json:"name"`
	Scope          string    `json:"scope"`
	Revoked        bool      `gorm:"default(0)" json:"revoked"`
	ExpiresAt      time.Time `gorm:"nullable" json:"expries_at"`
}
