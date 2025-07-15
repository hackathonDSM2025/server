package auth

import (
	"net/http"

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
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Invalid request body"))
		return
	}

	authData, err := h.service.Register(c.Request.Context(), &req)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusCreated, AuthResponse{
		Success: true,
		Data:    *authData,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("Invalid request body"))
		return
	}

	authData, err := h.service.Login(c.Request.Context(), &req)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Success: true,
		Data:    *authData,
	})
}