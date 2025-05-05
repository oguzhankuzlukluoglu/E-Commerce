package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oguzhan/e-commerce/internal/auth"
	"github.com/oguzhan/e-commerce/pkg/config"
	"github.com/oguzhan/e-commerce/pkg/database"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize service
	authService := auth.NewService(db)
	authHandler := auth.NewHandler(authService)

	// Initialize router
	router := gin.Default()

	// Register routes
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	router.GET("/me", authHandler.AuthMiddleware(), authHandler.GetUserFromToken)

	// Start server
	log.Printf("Auth service starting on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
