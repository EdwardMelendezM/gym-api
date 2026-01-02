package models

import "gym-api/internal/utils/pagination"

// Response DTO
type UserResponse struct {
	ID        string  `json:"id"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     string  `json:"email"`
}

// User GetUsersPaginated Response for Swagger
type UserListResponse struct {
	Data []UserResponse  `json:"data"`
	Meta pagination.Meta `json:"meta"`
}
