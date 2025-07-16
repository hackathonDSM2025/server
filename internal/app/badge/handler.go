package badge

import (
	"net/http"

	"hackathon-dsm-server/internal/pkg/middleware"
	"hackathon-dsm-server/internal/pkg/utils/errors"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	badge := router.Group("/badges")
	badge.Use(middleware.AuthMiddleware())
	{
		badge.GET("/me", h.GetUserBadges)
	}
}

func (h *Handler) GetUserBadges(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, errors.Unauthorized("User not authenticated"))
		return
	}

	data, err := h.service.GetUserBadges(c.Request.Context(), userID.(int))
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, BadgeListResponse{
		Success: true,
		Data:    *data,
	})
}