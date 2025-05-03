# Use the official Golang image as the base image
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /payment-service ./cmd/payment

# Create a new stage with a minimal image
FROM alpine:latest

# Copy the binary from builder
COPY --from=0 /payment-service .

# Command to run the executable
CMD ["./payment-service"] 