package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Every(time.Second), 100) // 100 requests per second

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			return
		}
		c.Next()
	}
}

// RedisRateLimiter is a more sophisticated rate limiter using Redis
func RedisRateLimiter(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		period := 1 * time.Hour
		maxRequests := 1000 // requests per hour

		count, err := redisClient.Incr(c, key).Result()
		if err != nil {
			c.Next() // In case of Redis error, let the request through
			return
		}

		if count == 1 {
			redisClient.Expire(c, key, period)
		}

		if count > int64(maxRequests) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
				"retry_after": time.Until(
					time.Now().Add(period),
				).Seconds(),
			})
			return
		}

		c.Next()
	}
}
