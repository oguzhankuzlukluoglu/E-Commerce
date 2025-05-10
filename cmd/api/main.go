package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oguzhan/e-commerce/internal/auth"
	"github.com/oguzhan/e-commerce/internal/cart"
	"github.com/oguzhan/e-commerce/internal/middleware"
	"github.com/oguzhan/e-commerce/internal/order"
	"github.com/oguzhan/e-commerce/internal/payment"
	"github.com/oguzhan/e-commerce/internal/product"
	"github.com/oguzhan/e-commerce/internal/user"
	"github.com/oguzhan/e-commerce/pkg/config"
	"github.com/oguzhan/e-commerce/pkg/database"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title           E-Commerce API
// @version         1.0
// @description     A modern e-commerce API built with Go and Gin.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Initialize logger
	logger, err := middleware.NewLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
	}

	// Initialize database
	db, err := database.InitDB(cfg)
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}

	// Run migrations
	if err := database.AutoMigrate(db); err != nil {
		logger.Fatal("Failed to run migrations", zap.Error(err))
	}

	// Initialize services
	authService := auth.NewService(db)
	userService := user.NewService(db)
	productService := product.NewService(db)
	orderService := order.NewService(db)
	paymentService := payment.NewService(db)
	cartService := cart.NewService(db)

	// Initialize handlers
	authHandler := auth.NewHandler(authService)
	userHandler := user.NewHandler(userService)
	productHandler := product.NewHandler(productService)
	orderHandler := order.NewHandler(orderService)
	paymentHandler := payment.NewHandler(paymentService)
	cartHandler := cart.NewHandler(cartService)

	// Initialize router
	router := gin.Default()

	// Apply middlewares
	rateLimitConfig := middleware.DefaultConfig()
	router.Use(middleware.RateLimit(rateLimitConfig))
	router.Use(middleware.Logger(logger))
	router.Use(middleware.Metrics())

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API routes
	api := router.Group("/api/v1")
	{
		// Auth routes
		api.POST("/auth/register", authHandler.Register)
		api.POST("/auth/login", authHandler.Login)
		api.GET("/me", authHandler.AuthMiddleware(), authHandler.GetUserFromToken)

		// User routes
		userGroup := api.Group("/users")
		userGroup.Use(authHandler.AuthMiddleware())
		{
			userGroup.GET("/:id", userHandler.GetUser)
			userGroup.PUT("/:id", userHandler.UpdateUser)
			userGroup.DELETE("/:id", userHandler.DeleteUser)
			userGroup.GET("", userHandler.ListUsers)
			userGroup.POST("/:id/change-password", userHandler.ChangePassword)

			// Admin only routes
			userGroup.POST("/:id/deactivate", userHandler.DeactivateUser)
			userGroup.POST("/:id/activate", userHandler.ActivateUser)
			userGroup.PUT("/:id/role", userHandler.UpdateUserRole)
			userGroup.POST("/:id/reset-password", userHandler.ResetPassword)

			// Address routes
			userGroup.POST("/addresses", userHandler.CreateAddress)
			userGroup.GET("/addresses", userHandler.GetAddresses)
			userGroup.PUT("/addresses/:id", userHandler.UpdateAddress)
			userGroup.DELETE("/addresses/:id", userHandler.DeleteAddress)

			// Contact routes
			userGroup.POST("/contacts", userHandler.CreateContact)
			userGroup.GET("/contacts", userHandler.GetContacts)
			userGroup.PUT("/contacts/:id", userHandler.UpdateContact)
			userGroup.DELETE("/contacts/:id", userHandler.DeleteContact)
		}

		// Product routes
		productGroup := api.Group("/products")
		{
			productGroup.GET("", productHandler.ListProducts)
			productGroup.GET("/:id", productHandler.GetProduct)
			productGroup.GET("/search", productHandler.SearchProducts)
		}

		// Protected product routes
		protectedProductGroup := api.Group("/products")
		protectedProductGroup.Use(authHandler.AuthMiddleware())
		{
			protectedProductGroup.POST("", productHandler.CreateProduct)
			protectedProductGroup.PUT("/:id", productHandler.UpdateProduct)
			protectedProductGroup.DELETE("/:id", productHandler.DeleteProduct)
		}

		// Order routes
		orderGroup := api.Group("/orders")
		orderGroup.Use(authHandler.AuthMiddleware())
		{
			orderGroup.POST("", orderHandler.CreateOrder)
			orderGroup.GET("/:id", orderHandler.GetOrder)
			orderGroup.GET("", orderHandler.ListOrders)
			orderGroup.PUT("/:id", orderHandler.UpdateOrder)
			orderGroup.DELETE("/:id", orderHandler.CancelOrder)
		}

		// Payment routes
		paymentGroup := api.Group("/payments")
		paymentGroup.Use(authHandler.AuthMiddleware())
		{
			paymentGroup.POST("", paymentHandler.CreatePayment)
			paymentGroup.GET("/:id", paymentHandler.GetPayment)
			paymentGroup.GET("", paymentHandler.ListPayments)
			paymentGroup.POST("/:id/process", paymentHandler.ProcessPayment)
			paymentGroup.POST("/:id/refund", paymentHandler.RefundPayment)
		}

		// Cart routes
		cartGroup := api.Group("/cart")
		cartGroup.Use(authHandler.AuthMiddleware())
		{
			cartGroup.GET("", cartHandler.GetCart)
			cartGroup.GET("/items", cartHandler.GetItems)
			cartGroup.POST("/items", cartHandler.AddItem)
			cartGroup.PUT("/items/:id", cartHandler.UpdateItem)
			cartGroup.DELETE("/items/:id", cartHandler.RemoveItem)
			cartGroup.DELETE("", cartHandler.ClearCart)
		}
	}

	// Start server
	logger.Info("Starting server", zap.String("port", cfg.ServerPort))
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
