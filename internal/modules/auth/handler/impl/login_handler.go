package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gym-api/internal/modules/auth/models"
	"gym-api/internal/utils/errors"
)

// Login godoc
// @Summary Login user
// @Description Login with email and password, returns access and refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login request"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 403 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var input models.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		errors.Respond(c, err)
		return
	}

	result, err := h.service.Login(input)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusCreated, models.AuthResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	})
}
