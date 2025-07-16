package badge

import (
	"net/http"
	"strconv"

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
	badges := router.Group("/badges")
	{
		badges.GET("", h.GetAllBadges)
		badges.GET("/:badgeId", h.GetBadgeByID)
	}
}

func (h *Handler) GetAllBadges(c *gin.Context) {
	data, err := h.service.GetAllBadges(c.Request.Context())
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, AllBadgesResponse{
		Success: true,
		Data:    *data,
	})
}

func (h *Handler) GetBadgeByID(c *gin.Context) {
	badgeIDStr := c.Param("badgeId")
	badgeID, err := strconv.Atoi(badgeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Invalid badge ID"))
		return
	}

	data, err := h.service.GetBadgeByID(c.Request.Context(), badgeID)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, BadgeDetailResponse{
		Success: true,
		Data:    *data,
	})
}