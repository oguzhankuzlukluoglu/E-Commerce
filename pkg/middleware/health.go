package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthCheck struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Database  string    `json:"database"`
	Redis     string    `json:"redis,omitempty"`
}

// HealthCheckMiddleware creates a health check endpoint
func HealthCheckHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		health := HealthCheck{
			Status:    "up",
			Timestamp: time.Now(),
			Database:  "up",
		}

		// Check database connection
		sqlDB, err := db.DB()
		if err != nil {
			health.Status = "degraded"
			health.Database = "down"
		} else if err := sqlDB.Ping(); err != nil {
			health.Status = "degraded"
			health.Database = "down"
		}

		if health.Status == "up" {
			c.JSON(http.StatusOK, health)
		} else {
			c.JSON(http.StatusServiceUnavailable, health)
		}
	}
}
