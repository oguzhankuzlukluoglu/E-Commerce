package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a middleware that logs the request details
func Logger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		ip := c.ClientIP()
		method := c.Request.Method
		userAgent := c.Request.UserAgent()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		// Create structured log entry
		logger.Info("incoming request",
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ip),
			zap.String("method", method),
			zap.String("user-agent", userAgent),
			zap.Int("status", statusCode),
			zap.Duration("latency", latency),
			zap.String("error", errorMessage),
		)
	}
}

// NewLogger creates a new zap logger instance
func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = "" // Disable stacktrace for production

	return config.Build()
}
