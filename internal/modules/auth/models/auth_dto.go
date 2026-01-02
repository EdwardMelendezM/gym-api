package models

type RegisterRequest struct {
	FirstName string `json:"firstName" binding:"required,min=2"`
	LastName  string `json:"lastName" binding:"required,min=2"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
