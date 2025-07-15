package quiz

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
	quiz := router.Group("/quiz")
	quiz.Use(middleware.AuthMiddleware())
	{
		quiz.GET("/:heritageId", h.GetQuiz)
		quiz.POST("/:heritageId/submit", h.SubmitQuiz)
	}
}

func (h *Handler) GetQuiz(c *gin.Context) {
	heritageIDStr := c.Param("heritageId")
	heritageID, err := strconv.Atoi(heritageIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Invalid heritage ID"))
		return
	}

	data, err := h.service.GetQuiz(c.Request.Context(), heritageID)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, QuizResponse{
		Success: true,
		Data:    *data,
	})
}

func (h *Handler) SubmitQuiz(c *gin.Context) {
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

	var req SubmitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Invalid request body"))
		return
	}

	data, err := h.service.SubmitQuiz(c.Request.Context(), userID.(int), heritageID, req.Answers)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, SubmitResponse{
		Success: true,
		Data:    *data,
	})
}