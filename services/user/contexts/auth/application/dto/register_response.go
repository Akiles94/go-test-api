package dto

type RegisterResponseDto struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
