package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oguzhan/e-commerce/internal/auth"
	"github.com/oguzhan/e-commerce/internal/order"
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

	// Initialize services
	orderService := order.NewService(db)
	orderHandler := order.NewHandler(orderService)

	// Initialize router
	router := gin.Default()

	// Register routes
	orderGroup := router.Group("/orders")
	orderGroup.Use(auth.NewHandler(nil).AuthMiddleware())
	{
		orderGroup.POST("/", orderHandler.CreateOrder)
		orderGroup.GET("/:id", orderHandler.GetOrder)
		orderGroup.GET("/user/:userID", orderHandler.GetUserOrders)
		orderGroup.PUT("/:id/status", orderHandler.UpdateOrderStatus)
		orderGroup.DELETE("/:id", orderHandler.CancelOrder)
	}

	// Start server
	log.Printf("Order service starting on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
