package dto

type PatchCategoryRequest struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=1,max=100"`
	Description *string `json:"description,omitempty" validate:"omitempty,min=1,max=500"`
	IsActive    *bool   `json:"is_active,omitempty"`
}
