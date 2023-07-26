package dto

import "time"

type FacebookTokenDTO struct {
	Name           string
	AccessToken    string
	RefreshToken   string
	ExpiresAt      time.Time
	UserID         uint64
	FacebookUserId string
}
