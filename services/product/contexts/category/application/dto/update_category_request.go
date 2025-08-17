package dto

type UpdateCategoryRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"required,min=1,max=500"`
	IsActive    bool   `json:"is_active"`
}
