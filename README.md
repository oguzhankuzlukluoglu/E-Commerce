# Microservices-based E-Commerce System

This is a microservices-based e-commerce system built with Go, Chi framework, and PostgreSQL.

## Project Structure

```
.
├── cmd/                    # Main applications for each microservice
│   ├── api/               # Main API gateway
│   ├── auth/              # Authentication service
│   ├── product/           # Product service
│   ├── order/             # Order service
│   ├── payment/           # Payment service
│   └── user/              # User service
├── internal/              # Private application and library code
│   ├── auth/             # Authentication service implementation
│   ├── cart/             # Shopping cart implementation
│   ├── middleware/       # HTTP middleware implementations
│   ├── model/           # Internal data models
│   ├── order/           # Order service implementation
│   ├── payment/         # Payment service implementation
│   ├── product/         # Product service implementation
│   ├── repository/      # Database repository implementations
│   ├── service/         # Business logic implementations
│   └── user/            # User service implementation
├── pkg/                   # Public library code
│   ├── cache/           # Caching utilities
│   ├── config/          # Configuration management
│   ├── database/        # Database connection and utilities
│   ├── docs/            # Documentation utilities
│   ├── errors/          # Error handling utilities
│   ├── logger/          # Logging utilities
│   ├── metrics/         # Metrics and monitoring
│   ├── middleware/      # Shared middleware components
│   ├── models/          # Shared data models
│   └── utils/           # Utility functions
├── docs/                 # Project documentation
├── migrations/           # Database migrations
└── deployments/          # Deployment configurations
```

## Technology Stack

### Backend Technologies
- **Go 1.21+**: Core programming language
- **Chi**: HTTP router and middleware
- **PostgreSQL**: Primary database
- **Redis**: Caching and session management
- **gRPC**: Inter-service communication
- **JWT**: Authentication and authorization
- **Docker**: Containerization
- **Prometheus**: Metrics and monitoring
- **Swagger/OpenAPI**: API documentation

### Development Tools
- **Go Modules**: Dependency management
- **Air**: Hot reload for development
- **Golangci-lint**: Code linting
- **Go Test**: Unit and integration testing
- **Make**: Build automation

## API Documentation

Detailed API documentation is available in [API_ENDPOINTS.md](API_ENDPOINTS.md). The API includes endpoints for:

- Authentication (Register, Login)
- Products (CRUD operations)
- Orders (Create, Update, Cancel)
- Payments (Process, Refund)
- User Management
- Cart Operations

## Default Admin Credentials

For initial setup, the following admin user is created:

```
Email: admin@example.com
Password: admin123
```

**Important**: Please change these credentials immediately after first login.

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/e-commerce.git
cd e-commerce
```

### 2. Environment Setup

1. Copy the example configuration file:
```bash
cp config.example.yaml config.yaml
```

2. Update the `config.yaml` file with your local settings:
   - Set your PostgreSQL credentials
   - Configure Redis connection details
   - Set JWT secret
   - Update service URLs if needed

3. Create a `.env` file in the root directory with the following variables:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=ecommerce
REDIS_HOST=localhost
REDIS_PORT=6379
JWT_SECRET=your_jwt_secret
```

### 3. Database Setup

1. Create the PostgreSQL database:
```bash
createdb ecommerce
```

2. Run database migrations:
```bash
go run migrations/main.go
```

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Running the Services

You can run each service individually:

```bash
# Auth Service
go run cmd/auth/main.go

# Product Service
go run cmd/product/main.go

# Order Service
go run cmd/order/main.go

# Payment Service
go run cmd/payment/main.go

# User Service
go run cmd/user/main.go
```

Or use the Makefile commands (if available):
```bash
make run-auth
make run-product
make run-order
make run-payment
make run-user
```

## Development

### Running Tests
```bash
go test ./...
```

### Code Generation
If you modify any protobuf files:
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/*.proto
```

## Troubleshooting

1. **Database Connection Issues**
   - Verify PostgreSQL is running
   - Check database credentials in config.yaml
   - Ensure database exists

2. **Service Connection Issues**
   - Verify all services are running
   - Check service URLs in config.yaml
   - Check network connectivity between services

3. **Redis Connection Issues**
   - Verify Redis is running
   - Check Redis connection details in config.yaml

## License

MIT 
