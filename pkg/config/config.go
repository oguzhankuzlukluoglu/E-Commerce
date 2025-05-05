package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	Env        string

	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	RedisHost     string
	RedisPort     int
	RedisPassword string
	RedisDB       int

	JWTSecret     string
	JWTExpiration time.Duration

	PaymentServiceURL string
	PaymentAPIKey     string

	OrderServiceURL   string
	ProductServiceURL string
	UserServiceURL    string

	EnableMetrics bool
	MetricsPort   int

	LogLevel  string
	LogFormat string
}

func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	// Parse JWT expiration duration
	jwtExpiration, err := time.ParseDuration(getEnv("JWT_EXPIRATION", "24h"))
	if err != nil {
		return nil, fmt.Errorf("error parsing JWT expiration: %v", err)
	}

	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8081"),
		Env:        getEnv("ENV", "development"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnvAsInt("DB_PORT", 5432),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "ecommerce"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnvAsInt("REDIS_PORT", 6379),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),

		JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiration: jwtExpiration,

		PaymentServiceURL: getEnv("PAYMENT_SERVICE_URL", "http://localhost:8084"),
		PaymentAPIKey:     getEnv("PAYMENT_API_KEY", ""),

		OrderServiceURL:   getEnv("ORDER_SERVICE_URL", "http://localhost:8082"),
		ProductServiceURL: getEnv("PRODUCT_SERVICE_URL", "http://localhost:8083"),
		UserServiceURL:    getEnv("USER_SERVICE_URL", "http://localhost:8081"),

		EnableMetrics: getEnvAsBool("ENABLE_METRICS", true),
		MetricsPort:   getEnvAsInt("METRICS_PORT", 9090),

		LogLevel:  getEnv("LOG_LEVEL", "debug"),
		LogFormat: getEnv("LOG_FORMAT", "json"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
