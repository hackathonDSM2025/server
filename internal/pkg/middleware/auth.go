package middleware

import (
	"net/http"
	"strings"

	"hackathon-dsm-server/internal/pkg/utils/errors"
	"hackathon-dsm-server/internal/pkg/utils/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, errors.Unauthorized("Authorization header required"))
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, errors.Unauthorized("Invalid authorization header format"))
			c.Abort()
			return
		}

		claims, err := jwt.ValidateAccessToken(tokenParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}