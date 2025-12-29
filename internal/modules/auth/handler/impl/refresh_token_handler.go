package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gym-api/internal/modules/auth/models"
	"gym-api/internal/modules/shared/errors"
)

// RefreshToken Refresh token godoc
// @Summary Refresh access token
// @Description Refresh access token using refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param refresh body RefreshTokenRequest true "Refresh token request"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} errors.AppError
// @Router /auth/refresh-token [post]
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
