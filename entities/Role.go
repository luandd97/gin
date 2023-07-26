package entities

import "time"

type Role struct {
	Id          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	DisplayName string    `gorm:"type:varchar(255)" json:"display_name"`
	Description string    `gorm:"text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Users       []User    `gorm:"many2many:role_users"`
}
