package dto

type RegisterDTO struct {
	Name         string `json:"name" form:"name" binding:"required"`
	Email        string `json:"email" form:"email" binding:"required"`
	Password     string `json:"password" form:"password" binding:"required"`
	Phone        string `json:"phone" form:"phone" binding:"required"`
	SocialDriver string `json:"social_driver" form:"social_driver"`
	Status       bool   `json:"status" form:"status"`
}
