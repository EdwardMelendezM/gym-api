package impl

import (
	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserById godoc
// @Summary Get user by ID
// @Description Returns a user by its unique identifier
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} errors.ErrorResponse "Missing or invalid user id"
// @Failure 404 {object} errors.ErrorResponse "User not found"
// @Failure 500 {object} errors.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		errors.Respond(c, errors.New().
			SetStatus(http.StatusBadRequest).
			SetLayer("users.handler").
			SetFunction("GetUserById").
			SetMessage("id parameter is required"))
		return
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
}
