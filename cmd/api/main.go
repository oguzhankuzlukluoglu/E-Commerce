package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/oguzhan/e-commerce/internal/order"
	"github.com/oguzhan/e-commerce/internal/payment"
	"github.com/oguzhan/e-commerce/pkg/metrics"
)

func main() {
	// Initialize database connection
	dsn := "host=localhost user=postgres password=postgres dbname=ecommerce port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Initialize repositories
	orderRepo := order.NewRepository(db)
	paymentRepo := payment.NewRepository(db)

	// Initialize services
	orderService := order.NewService(orderRepo)
	paymentService := payment.NewService(paymentRepo, rdb)

	// Initialize HTTP handlers
	orderHandler := order.NewHandler(orderService)
	paymentHandler := payment.NewHandler(paymentService)

	// Initialize Gin router
	router := gin.Default()

	// Add Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Add middleware
	router.Use(metrics.HTTPMetricsMiddleware())

	// Register routes
	orderHandler.RegisterRoutes(router)
	paymentHandler.RegisterRoutes(router)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
