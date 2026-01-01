package impl

import (
	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/auth"
	"gym-api/internal/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMe godoc
// @Summary Get current authenticated user
// @Description Returns the authenticated user's profile using the access token
// @Tags Users
// @Produce json
// @Success 200 {object} models.UserResponse
// @Failure 401 {object} errors.ErrorResponse "User not authenticated"
// @Failure 404 {object} errors.ErrorResponse "User not found"
// @Failure 500 {object} errors.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /api/v1/users/me [get]
func (h *UserHandler) GetMe(c *gin.Context) {
	session, ok := auth.GetSession(c)
	if !ok {
		errors.Respond(c, errors.New().
			SetStatus(http.StatusUnauthorized).
			SetLayer("users.handler").
			SetMessage("user not authenticated"))
		return
	}

	user, err := h.service.GetUserById(session.UserID)
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
