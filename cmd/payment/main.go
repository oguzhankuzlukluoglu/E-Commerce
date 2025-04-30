package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oguzhan/e-commerce/internal/payment"
	"github.com/oguzhan/e-commerce/pkg/database"
	"github.com/oguzhan/e-commerce/pkg/models"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Auto migrate the payment model
	if err := db.AutoMigrate(&models.Payment{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize payment service and handler
	paymentRepo := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepo, rdb)
	paymentHandler := payment.NewHandler(paymentService)

	// Set up Gin router
	router := gin.Default()

	// Payment routes
	router.POST("/payments", paymentHandler.CreatePayment)
	router.POST("/payments/:id/process", paymentHandler.ProcessPayment)
	router.GET("/payments/:id", paymentHandler.GetPayment)
	router.GET("/payments/user/:user_id", paymentHandler.GetUserPayments)
	router.GET("/payments/order/:order_id", paymentHandler.GetOrderPayments)
	router.POST("/payments/:id/refund", paymentHandler.RefundPayment)
	router.GET("/payments", paymentHandler.ListPayments)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084" // Default port for payment service
	}

	// Start the server
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
