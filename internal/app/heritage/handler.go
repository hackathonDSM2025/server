package heritage

import (
	"net/http"
	"strconv"

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
		heritage.GET("", h.SearchHeritage)
		heritage.POST("/:heritageId/visits", middleware.AuthMiddleware(), h.CreateVisit)
		heritage.POST("/:heritageId/reviews", middleware.AuthMiddleware(), h.CreateReview)
		heritage.PUT("/:heritageId/reviews", middleware.AuthMiddleware(), h.UpdateReview)
		heritage.GET("/:heritageId/reviews/me", middleware.AuthMiddleware(), h.GetMyReview)
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

func (h *Handler) CreateVisit(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, errors.Unauthorized("User not authenticated"))
		return
	}

	heritageIDStr := c.Param("heritageId")
	heritageID, err := strconv.Atoi(heritageIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Invalid heritage ID"))
		return
	}

	var req CreateVisitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("QR code is required"))
		return
	}

	data, err := h.service.CreateVisit(c.Request.Context(), userID.(int), heritageID, req.QRCode)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusCreated, CreateVisitResponse{
		Success: true,
		Data:    *data,
	})
}

func (h *Handler) CreateReview(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, errors.Unauthorized("User not authenticated"))
		return
	}

	heritageIDStr := c.Param("heritageId")
	heritageID, err := strconv.Atoi(heritageIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Invalid heritage ID"))
		return
	}

	var req ReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Invalid request body"))
		return
	}

	err = h.service.CreateOnlyReview(c.Request.Context(), userID.(int), heritageID, req.Rating, req.ReviewText)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusCreated, CreateReviewResponse{
		Success: true,
		Message: "Review created successfully",
	})
}

func (h *Handler) UpdateReview(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, errors.Unauthorized("User not authenticated"))
		return
	}

	heritageIDStr := c.Param("heritageId")
	heritageID, err := strconv.Atoi(heritageIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Invalid heritage ID"))
		return
	}

	var req ReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Invalid request body"))
		return
	}

	err = h.service.UpdateOnlyReview(c.Request.Context(), userID.(int), heritageID, req.Rating, req.ReviewText)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, UpdateReviewResponse{
		Success: true,
		Message: "Review updated successfully",
	})
}

func (h *Handler) GetMyReview(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, errors.Unauthorized("User not authenticated"))
		return
	}

	heritageIDStr := c.Param("heritageId")
	heritageID, err := strconv.Atoi(heritageIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Invalid heritage ID"))
		return
	}

	data, err := h.service.GetMyReview(c.Request.Context(), userID.(int), heritageID)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, GetMyReviewResponse{
		Success: true,
		Data:    data,
	})
}

