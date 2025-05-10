package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HTTP request counter
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// HTTP request duration
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		},
		[]string{"method", "path"},
	)

	// Active requests
	httpRequestsInProgress = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_requests_in_progress",
			Help: "Number of HTTP requests in progress",
		},
		[]string{"method", "path"},
	)

	// Database operations
	dbOperationsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_operations_total",
			Help: "Total number of database operations",
		},
		[]string{"operation", "table"},
	)

	// Database operation duration
	dbOperationDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_operation_duration_seconds",
			Help:    "Duration of database operations in seconds",
			Buckets: []float64{.001, .005, .01, .025, .05, .1, .25, .5, 1},
		},
		[]string{"operation", "table"},
	)
)

// Metrics middleware for Prometheus
func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Increment in-progress requests
		httpRequestsInProgress.WithLabelValues(method, path).Inc()

		// Process request
		c.Next()

		// Decrement in-progress requests
		httpRequestsInProgress.WithLabelValues(method, path).Dec()

		// Record request duration
		duration := time.Since(start).Seconds()
		httpRequestDuration.WithLabelValues(method, path).Observe(duration)

		// Record request count
		status := strconv.Itoa(c.Writer.Status())
		httpRequestsTotal.WithLabelValues(method, path, status).Inc()
	}
}

// RecordDBOperation records database operation metrics
func RecordDBOperation(operation, table string, duration time.Duration) {
	dbOperationsTotal.WithLabelValues(operation, table).Inc()
	dbOperationDuration.WithLabelValues(operation, table).Observe(duration.Seconds())
}
