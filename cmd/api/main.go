package main

import (
	"log"

	"hackathon-dsm-server/internal/app/auth"
	"hackathon-dsm-server/internal/app/badge"
	"hackathon-dsm-server/internal/app/heritage"
	"hackathon-dsm-server/internal/app/quiz"
	"hackathon-dsm-server/internal/app/user"
	"hackathon-dsm-server/internal/pkg/config"
	"hackathon-dsm-server/internal/pkg/db/mysql"
	"hackathon-dsm-server/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connections
	mysqlDB := mysql.NewConnection()

	// Initialize repositories
	authRepo := auth.NewMySQLRepository(mysqlDB)
	badgeRepo := badge.NewMySQLRepository(mysqlDB)
	heritageRepo := heritage.NewMySQLRepository(mysqlDB)
	quizRepo := quiz.NewMySQLRepository(mysqlDB)
	userRepo := user.NewMySQLRepository(mysqlDB)

	// Initialize services
	authService := auth.NewService(authRepo)
	badgeService := badge.NewService(badgeRepo)
	heritageService := heritage.NewService(heritageRepo)
	quizService := quiz.NewService(quizRepo)
	userService := user.NewService(userRepo)

	// Initialize handlers
	authHandler := auth.NewHandler(authService)
	badgeHandler := badge.NewHandler(badgeService)
	heritageHandler := heritage.NewHandler(heritageService)
	quizHandler := quiz.NewHandler(quizService)
	userHandler := user.NewHandler(userService)

	// Initialize Gin router
	router := gin.Default()

	// Apply middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ErrorHandler())
	
	// Add charset handling middleware
	router.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Next()
	})

	// API routes
	api := router.Group("/api")
	{
		authHandler.RegisterRoutes(api)
		badgeHandler.RegisterRoutes(api)
		heritageHandler.RegisterRoutes(api)
		quizHandler.RegisterRoutes(api)
		userHandler.RegisterRoutes(api)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "heritage-tour-api",
		})
	})

	log.Printf("Server starting on port %s", config.GlobalConfig.App.Port)
	if err := router.Run(":" + config.GlobalConfig.App.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Close database connections on shutdown
	defer mysqlDB.Close()
}
