package entities

type RoleUser struct {
	RoleId   uint64 `gorm:"index,not null" json:"role_id"`
	UserId   uint64 `gorm:"index,not null" json:"user_id"`
	UserType string ` json:"user_type"`
}
