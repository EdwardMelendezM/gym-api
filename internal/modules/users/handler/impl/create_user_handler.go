package impl

import (
	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create a new user
// @Description Creates a new user in the system
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User creation payload"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} errors.ErrorResponse "Invalid request body"
// @Failure 409 {object} errors.ErrorResponse "User already exists"
// @Failure 500 {object} errors.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /api/v1/users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Respond(c, err)
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusCreated, models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
}
