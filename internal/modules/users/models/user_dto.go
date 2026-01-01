package models

import "gym-api/internal/utils/pagination"

// Response DTO
type UserResponse struct {
	ID        string  `json:"id"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     string  `json:"email"`
}

// User List Response for Swagger
type UserListResponse struct {
	Data []UserResponse  `json:"data"`
	Meta pagination.Meta `json:"meta"`
}

// Request DTO (para más adelante)
type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
