package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oguzhan/e-commerce/internal/auth"
	"github.com/oguzhan/e-commerce/internal/user"
	"github.com/oguzhan/e-commerce/pkg/config"
	"github.com/oguzhan/e-commerce/pkg/database"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize services
	authService := auth.NewService(db)
	userService := user.NewService(db)

	// Initialize handlers
	authHandler := auth.NewHandler(authService)
	userHandler := user.NewHandler(userService)

	// Initialize router
	router := gin.Default()

	// Auth routes
	router.POST("/auth/register", authHandler.Register)
	router.POST("/auth/login", authHandler.Login)

	// User routes
	userGroup := router.Group("/users")
	userGroup.Use(authHandler.AuthMiddleware())
	{
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
		userGroup.GET("", userHandler.ListUsers)
		userGroup.POST("/:id/change-password", userHandler.ChangePassword)
	}

	// Start server
	log.Printf("Starting user service on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
