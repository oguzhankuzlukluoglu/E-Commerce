package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oguzhan/e-commerce/internal/auth"
	"github.com/oguzhan/e-commerce/internal/product"
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
	productService := product.NewService(db)
	productHandler := product.NewHandler(productService)

	// Initialize router
	router := gin.Default()

	// Register routes
	productGroup := router.Group("/products")
	{
		productGroup.GET("/", productHandler.ListProducts)
		productGroup.GET("/search", productHandler.SearchProducts)
		productGroup.GET("/:id", productHandler.GetProduct)
	}

	// Protected routes
	protectedGroup := router.Group("/products")
	protectedGroup.Use(auth.NewHandler(nil).AuthMiddleware())
	{
		protectedGroup.POST("/", productHandler.CreateProduct)
		protectedGroup.PUT("/:id", productHandler.UpdateProduct)
		protectedGroup.DELETE("/:id", productHandler.DeleteProduct)
		protectedGroup.PUT("/:id/stock", productHandler.UpdateStock)
	}

	// Start server
	log.Printf("Product service starting on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
