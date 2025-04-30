package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func HTTPMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())

		HttpRequestDuration.WithLabelValues(
			c.Request.Method,
			c.Request.URL.Path,
			status,
		).Observe(duration)

		HttpRequestsTotal.WithLabelValues(
			c.Request.Method,
			c.Request.URL.Path,
			status,
		).Inc()
	}
}
