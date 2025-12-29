package auth

import (
	"gym-api/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(c *gin.Context) {
	var input RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		errors.Respond(c, errors.BadRequest("invalid request body"))
		return
	}

	token, err := h.service.Register(input)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusCreated, AuthResponse{AccessToken: token})
}

func (h *Handler) Login(c *gin.Context) {
	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		errors.Respond(c, errors.BadRequest("invalid request body"))
		return
	}

	token, err := h.service.Login(input)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, AuthResponse{AccessToken: token})
}
