package middleware

import (
	"net/http"

	"hackathon-dsm-server/internal/pkg/utils/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if customErr, ok := err.(*errors.CustomError); ok {
				c.JSON(customErr.Status, customErr)
			} else {
				c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
			}
		}
	}
}