package dto

type LoginResponseDto struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
