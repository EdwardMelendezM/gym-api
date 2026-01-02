package impl

import (
	"gym-api/internal/utils/errors"
	"gym-api/internal/utils/pagination"
	"net/http"

	_ "gym-api/internal/modules/users/models"

	"github.com/gin-gonic/gin"
)

// GetUsersPaginated godoc
// @Summary List users
// @Description Returns a paginated list of users
// @Tags Users
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Items per page" default(10)
// @Param search query string false "Search text"
// @Param orderBy query string false "Field to order by"
// @Param order query string false "Sort direction (asc|desc)"
// @Success 200 {object} models.UserListResponse
// @Failure 400 {object} errors.ErrorResponse "Invalid query parameters"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 500 {object} errors.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /api/v1/users [get]
func (h *UserHandler) GetUsersPaginated(c *gin.Context) {
	ctx := c.Request.Context()
	p := c.MustGet(pagination.Key).(pagination.Params)

	result, err := h.service.GetUsersPaginated(ctx, p)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
