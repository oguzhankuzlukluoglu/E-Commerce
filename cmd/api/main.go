package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oguzhan/e-commerce/internal/auth"
	"github.com/oguzhan/e-commerce/internal/order"
	"github.com/oguzhan/e-commerce/internal/payment"
	"github.com/oguzhan/e-commerce/internal/product"
	"github.com/oguzhan/e-commerce/internal/user"
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
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize services
	authService := auth.NewService(db)
	userService := user.NewService(db)
	productService := product.NewService(db)
	orderService := order.NewService(db)
	paymentService := payment.NewService(db)

	// Initialize handlers
	authHandler := auth.NewHandler(authService)
	userHandler := user.NewHandler(userService)
	productHandler := product.NewHandler(productService)
	orderHandler := order.NewHandler(orderService)
	paymentHandler := payment.NewHandler(paymentService)

	// Initialize router
	router := gin.Default()

	// Auth routes
	router.POST("/auth/register", authHandler.Register)
	router.POST("/auth/login", authHandler.Login)
	router.GET("/me", authHandler.AuthMiddleware(), authHandler.GetUserFromToken)

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

	// Product routes
	productGroup := router.Group("/products")
	{
		productGroup.GET("", productHandler.ListProducts)
		productGroup.GET("/:id", productHandler.GetProduct)
		productGroup.GET("/search", productHandler.SearchProducts)
	}

	// Protected product routes
	protectedProductGroup := router.Group("/products")
	protectedProductGroup.Use(authHandler.AuthMiddleware())
	{
		protectedProductGroup.POST("", productHandler.CreateProduct)
		protectedProductGroup.PUT("/:id", productHandler.UpdateProduct)
		protectedProductGroup.DELETE("/:id", productHandler.DeleteProduct)
	}

	// Order routes
	orderGroup := router.Group("/orders")
	orderGroup.Use(authHandler.AuthMiddleware())
	{
		orderGroup.POST("", orderHandler.CreateOrder)
		orderGroup.GET("/:id", orderHandler.GetOrder)
		orderGroup.GET("", orderHandler.ListOrders)
		orderGroup.PUT("/:id", orderHandler.UpdateOrder)
		orderGroup.DELETE("/:id", orderHandler.CancelOrder)
	}

	// Payment routes
	paymentGroup := router.Group("/payments")
	paymentGroup.Use(authHandler.AuthMiddleware())
	{
		paymentGroup.POST("", paymentHandler.CreatePayment)
		paymentGroup.GET("/:id", paymentHandler.GetPayment)
		paymentGroup.GET("", paymentHandler.ListPayments)
		paymentGroup.POST("/:id/process", paymentHandler.ProcessPayment)
		paymentGroup.POST("/:id/refund", paymentHandler.RefundPayment)
	}

	// Start server
	log.Printf("Starting server on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
