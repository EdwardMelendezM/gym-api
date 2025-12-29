package users

import (
	"gym-api/internal/modules/shared/utils"
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

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		errors.Respond(c, errors.New().
			SetStatus(http.StatusBadRequest).
			SetLayer("users.handler").
			SetFunction("GetById").
			SetMessage("id parameter is required"))
		return
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
}

func (h *Handler) GetMe(c *gin.Context) {
	session, ok := utils.GetSession(c)
	if !ok {
		errors.Respond(c, errors.New().
			SetStatus(http.StatusUnauthorized).
			SetMessage("user not authenticated"))
		return
	}

	user, err := h.service.GetUserById(session.UserID)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
}
