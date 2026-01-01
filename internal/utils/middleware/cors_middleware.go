package middleware

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware(allowOrigins string) gin.HandlerFunc {
	allowed := map[string]struct{}{}

	for _, o := range strings.Split(allowOrigins, ",") {
		origin := strings.TrimSpace(o)
		if origin != "" {
			allowed[origin] = struct{}{}
		}
	}

	return cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			if _, ok := allowed["*"]; ok {
				return true
			}
			_, ok := allowed[origin]
			return ok
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
}
