package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimitConfig holds the configuration for rate limiting
type RateLimitConfig struct {
	Rate      int           // Number of requests allowed
	Period    time.Duration // Time period for the rate limit
	StoreType string        // Type of store to use (memory, redis, etc.)
}

// DefaultConfig returns a default rate limit configuration
func DefaultConfig() *RateLimitConfig {
	return &RateLimitConfig{
		Rate:      100,         // 100 requests
		Period:    time.Minute, // per minute
		StoreType: "memory",    // using in-memory store
	}
}

// RateLimit creates a new rate limiting middleware
func RateLimit(config *RateLimitConfig) gin.HandlerFunc {
	// Create a new rate limiter
	store := memory.NewStore()
	rate := limiter.Rate{
		Period: config.Period,
		Limit:  int64(config.Rate),
	}
	limiter := limiter.New(store, rate)

	return func(c *gin.Context) {
		// Get the IP address of the client
		ip := c.ClientIP()

		// Get the current rate limit for the IP
		context, err := limiter.Get(c.Request.Context(), ip)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "rate limit error",
			})
			c.Abort()
			return
		}

		// Set rate limit headers
		c.Header("X-RateLimit-Limit", string(context.Limit))
		c.Header("X-RateLimit-Remaining", string(context.Remaining))
		c.Header("X-RateLimit-Reset", string(context.Reset))

		// Check if the rate limit is exceeded
		if context.Reached {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
