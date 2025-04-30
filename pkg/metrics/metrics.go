package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HTTP Metrics
	HttpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
	)

	HttpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// Business Metrics
	OrdersCreated = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "orders_created_total",
			Help: "Total number of orders created",
		},
	)

	PaymentsProcessed = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "payments_processed_total",
			Help: "Total number of payments processed",
		},
	)

	PaymentAmount = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "payment_amount",
			Help:    "Distribution of payment amounts",
			Buckets: []float64{10, 50, 100, 500, 1000, 5000},
		},
	)

	// Error Metrics
	ErrorsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "errors_total",
			Help: "Total number of errors by type",
		},
		[]string{"type", "service"},
	)
)

func RecordHttpRequest(method, path, status string, duration float64) {
	HttpRequestDuration.WithLabelValues(method, path, status).Observe(duration)
	HttpRequestsTotal.WithLabelValues(method, path, status).Inc()
}

func RecordOrderCreated() {
	OrdersCreated.Inc()
}

func RecordPaymentProcessed(amount float64) {
	PaymentsProcessed.Inc()
	PaymentAmount.Observe(amount)
}

func RecordError(errorType, service string) {
	ErrorsTotal.WithLabelValues(errorType, service).Inc()
}
