package dto

type RoleDTO struct {
	Name        string `json:"name" form:"name" binding:"required"`
	DisplayName string `json:"display_name" form:"display_name" binding:"required"`
	Description string `json:"description" form:"description" binding:"nullable"`
}
