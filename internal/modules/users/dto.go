package users

// Response DTO
type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Request DTO (para más adelante)
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
