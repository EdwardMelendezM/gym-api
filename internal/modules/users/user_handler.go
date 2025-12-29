package users

import (
	"net/http"

	"gym-api/internal/errors"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewUserHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) List(c *gin.Context) {
	users, err := h.service.ListUsers()
	if err != nil {
		errors.Respond(c, err)
		return
	}

	response := make([]UserResponse, 0, len(users))
	for _, u := range users {
		response = append(response, UserResponse{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Respond(c, err)
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusCreated, UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
}
