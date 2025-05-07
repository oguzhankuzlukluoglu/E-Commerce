package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/oguzhan/e-commerce/internal/user/handler"
	"github.com/oguzhan/e-commerce/internal/user/repository"
	"github.com/oguzhan/e-commerce/internal/user/service"
)

func main() {
	// Database connection
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/ecommerce?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize dependencies
	userRepo := repository.NewUserRepository(db)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key" // In production, use a secure secret
	}
	userService := service.NewUserService(userRepo, jwtSecret)
	userHandler := handler.NewUserHandler(userService)

	// Router setup
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Routes
	r.Post("/register", userHandler.Register)
	r.Post("/login", userHandler.Login)
	r.Get("/users/{id}", userHandler.GetUser)
	r.Put("/users/{id}", userHandler.UpdateUser)
	r.Delete("/users/{id}", userHandler.DeleteUser)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("User service starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
