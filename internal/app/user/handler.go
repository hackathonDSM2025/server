package user

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
	users := router.Group("/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("/me", h.GetUserProfile)
		users.GET("/me/visits", h.GetMyVisits)
		users.GET("/me/reviews", h.GetMyReviews)
		users.GET("/me/badges", h.GetMyBadges)
	}
}

func (h *Handler) GetUserProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, errors.Unauthorized("User not authenticated"))
		return
	}

	data, err := h.service.GetUserProfile(c.Request.Context(), userID.(int))
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, UserProfileResponse{
		Success: true,
		Data:    *data,
	})
}

func (h *Handler) GetMyVisits(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, errors.Unauthorized("User not authenticated"))
		return
	}

	countOnly := c.Query("count_only") == "true"

	if countOnly {
		count, err := h.service.GetMyVisitsCount(c.Request.Context(), userID.(int))
		if err != nil {
			if customErr, ok := err.(*errors.CustomError); ok {
				c.JSON(customErr.Status, customErr)
			} else {
				c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, MyVisitsCountResponse{
			Success: true,
			Data: MyVisitsCountData{
				VisitCount: count,
			},
		})
		return
	}

	data, err := h.service.GetMyVisits(c.Request.Context(), userID.(int))
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, MyVisitsResponse{
		Success: true,
		Data:    *data,
	})
}


func (h *Handler) GetMyReviews(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, errors.Unauthorized("User not authenticated"))
		return
	}

	countOnly := c.Query("count_only") == "true"

	if countOnly {
		count, err := h.service.GetMyReviewsCount(c.Request.Context(), userID.(int))
		if err != nil {
			if customErr, ok := err.(*errors.CustomError); ok {
				c.JSON(customErr.Status, customErr)
			} else {
				c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, MyReviewsCountResponse{
			Success: true,
			Data: MyReviewsCountData{
				ReviewCount: count,
			},
		})
		return
	}

	data, err := h.service.GetMyReviews(c.Request.Context(), userID.(int))
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, MyReviewsResponse{
		Success: true,
		Data:    *data,
	})
}


func (h *Handler) GetMyBadges(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, errors.Unauthorized("User not authenticated"))
		return
	}

	countOnly := c.Query("count_only") == "true"

	if countOnly {
		count, err := h.service.GetMyBadgesCount(c.Request.Context(), userID.(int))
		if err != nil {
			if customErr, ok := err.(*errors.CustomError); ok {
				c.JSON(customErr.Status, customErr)
			} else {
				c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
			}
			return
		}

		c.JSON(http.StatusOK, MyBadgesCountResponse{
			Success: true,
			Data: MyBadgesCountData{
				BadgeCount: count,
			},
		})
		return
	}

	data, err := h.service.GetMyBadges(c.Request.Context(), userID.(int))
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok {
			c.JSON(customErr.Status, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("Internal server error"))
		}
		return
	}

	c.JSON(http.StatusOK, MyBadgesResponse{
		Success: true,
		Data:    *data,
	})
}

