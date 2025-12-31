package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gym-api/internal/modules/auth/models"
	"gym-api/internal/utils/errors"
)

// Register godoc
// @Summary Register user
// @Description Register with email and password, returns access and refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login request"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} errors.AppError
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var input models.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		errors.Respond(c, err)
		return
	}

	result, err := h.service.Register(input)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusCreated, models.AuthResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	})
}
