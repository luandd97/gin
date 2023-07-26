package entities

import "time"

type User struct {
	ID                uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name              string    `gorm:"index:idx_name,type:vachar(255)" json:"name"`
	Email             string    `gorm:"index:idx_email,type:vachar(255),unique" json:"email"`
	Phone             string    `gorm:"type:varchar(15)" json:"phone"`
	Password          string    `gorm:"->;<-;not null" json:"-"`
	Status            bool      `gorm:"default(0)" json:"status"`
	Token             string    `gorm:"type:varchar(255)" json:"token"`
	AccessToken       string    `json:"access_token"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Roles             []Role    `gorm:"many2many:role_users"`
	OauthAccessTokens []OauthAccessToken
	FacebookPages     []FacebookPage
	SocialDriver      string `gorm:"type:varchar(30);default:'web'" json:"social_driver"`
}
