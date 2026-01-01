package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gym-api/internal/modules/auth/models"
	"gym-api/internal/utils/errors"
)

// RefreshToken godoc
// @Summary Refresh access token
// @Description Generates a new access token using a valid refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param refresh body models.RefreshTokenRequest true "Refresh token payload"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {object} errors.ErrorResponse "Invalid request or token"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 500 {object} errors.ErrorResponse "Internal server error"
// @Router /api/v1/auth/refresh-token [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var input models.RefreshTokenRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		errors.Respond(c, err)
		return
	}

	result, err := h.service.RefreshToken(input)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, models.AuthResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	})
}
