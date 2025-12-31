package impl

import (
	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
