package heritage

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
	heritage := router.Group("/heritage")
	{
		heritage.GET("/search", h.SearchHeritage)
		heritage.POST("/scan", middleware.AuthMiddleware(), h.ScanQRCode)
	}
}

func (h *Handler) SearchHeritage(c *gin.Context) {
	var req SearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Keyword is required"))
		return
	}

	data, err := h.service.SearchHeritage(c.Request.Context(), req.Keyword)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, SearchResponse{
		Success: true,
		Data:    *data,
	})
}

func (h *Handler) ScanQRCode(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, errors.Unauthorized("User not authenticated"))
		return
	}

	var req ScanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("QR code is required"))
		return
	}

	data, err := h.service.ScanQRCode(c.Request.Context(), userID.(int), req.QRCode)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, ScanResponse{
		Success: true,
		Data:    *data,
	})
}