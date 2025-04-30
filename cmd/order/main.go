package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oguzhan/e-commerce/internal/order"
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
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize order service
	orderRepo := order.NewRepository(db)
	orderService := order.NewService(orderRepo)

	// Initialize router
	r := gin.Default()

	// Register routes
	orderHandler := order.NewHandler(orderService)
	orderHandler.RegisterRoutes(r)

	// Start server
	log.Printf("Starting order service on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
