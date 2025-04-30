package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oguzhan/e-commerce/pkg/logger"
	"go.uber.org/zap"
)

// CustomError represents a structured error response
type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface
func (e *CustomError) Error() string {
	return e.Message
}

// ErrorHandler is a middleware for handling errors globally
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// Log the error
			logger.Log.Error("request error",
				zap.String("path", c.Request.URL.Path),
				zap.String("error", err.Error()),
			)

			// Check if it's a custom error
			if customErr, ok := err.Err.(*CustomError); ok {
				c.JSON(customErr.Code, gin.H{
					"error": customErr.Message,
				})
				return
			}

			// Default error response
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}
	}
}

// NewCustomError creates a new custom error
func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}
