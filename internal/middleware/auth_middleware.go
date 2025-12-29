package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gym-api/internal/ent"
	"gym-api/internal/modules/sessions"
	"gym-api/internal/modules/shared/utils"
)

const sessionKey = "session"

func AuthMiddleware(sessions sessions.SessionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		session, err := sessions.FindByID(claims.SessionID)
		if err != nil || session.ExpiresAt.Before(time.Now()) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(sessionKey, session)
		c.Next()
	}
}

func SetupAuthMiddleware(client *ent.Client) gin.HandlerFunc {
	sessionRepository := sessions.NewSessionEntRepository(client)
	return AuthMiddleware(sessionRepository)
}
