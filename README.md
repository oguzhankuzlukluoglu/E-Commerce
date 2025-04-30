# Microservices-based E-Commerce System

This is a microservices-based e-commerce system built with Go, Chi framework, and PostgreSQL.

## Project Structure

```
.
├── cmd/                    # Main applications for each microservice
│   ├── auth/              # Authentication service
│   ├── product/           # Product service
│   ├── order/             # Order service
│   ├── payment/           # Payment service
│   └── user/              # User service
├── internal/              # Private application and library code
│   ├── auth/             # Authentication service implementation
│   ├── product/          # Product service implementation
│   ├── order/            # Order service implementation
│   ├── payment/          # Payment service implementation
│   └── user/             # User service implementation
├── pkg/                   # Public library code
│   ├── config/           # Configuration management
│   ├── models/           # Shared data models
│   └── utils/            # Utility functions
└── deployments/          # Deployment configurations
    └── docker/           # Docker configurations
```

## Services

1. **Auth Service**: Handles user authentication and authorization
2. **Product Service**: Manages product catalog and inventory
3. **Order Service**: Handles order processing and management
4. **Payment Service**: Manages payment processing
5. **User Service**: Handles user profile and management

## Technology Stack

- Go
- Chi (HTTP router)
- PostgreSQL
- Docker
- gRPC (for inter-service communication)

## Getting Started

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Set up PostgreSQL database
4. Configure environment variables
5. Run services using Docker Compose

## Development

To start development:

1. Install Go 1.21 or later
2. Install PostgreSQL
3. Set up your development environment
4. Run `go mod tidy` to install dependencies

## License

MIT 