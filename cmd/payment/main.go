package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oguzhan/e-commerce/internal/auth"
	"github.com/oguzhan/e-commerce/internal/payment"
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
	paymentService := payment.NewService(db)
	paymentHandler := payment.NewHandler(paymentService)

	// Initialize router
	router := gin.Default()

	// Register routes
	paymentGroup := router.Group("/payments")
	paymentGroup.Use(auth.NewHandler(nil).AuthMiddleware())
	{
		paymentGroup.POST("/", paymentHandler.CreatePayment)
		paymentGroup.POST("/:id/process", paymentHandler.ProcessPayment)
		paymentGroup.GET("/:id", paymentHandler.GetPayment)
		paymentGroup.GET("/user/:user_id", paymentHandler.GetUserPayments)
		paymentGroup.GET("/order/:order_id", paymentHandler.GetOrderPayments)
		paymentGroup.POST("/:id/refund", paymentHandler.RefundPayment)
		paymentGroup.GET("/", paymentHandler.ListPayments)
	}

	// Start server
	log.Printf("Payment service starting on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
