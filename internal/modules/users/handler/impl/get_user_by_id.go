package impl

import (
	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GetById(c *gin.Context) {
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

	c.JSON(http.StatusOK, models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
}
