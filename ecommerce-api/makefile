.PHONY: help run build test clean migrate-up migrate-down docker-up docker-down

# Default target - show help
help:
	@echo "Available commands:"
	@echo "  make run          - Run the application"
	@echo "  make build        - Build the application binary"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make migrate-up   - Apply database migrations"
	@echo "  make migrate-down - Rollback database migrations"
	@echo "  make docker-up    - Start with Docker Compose"
	@echo "  make docker-down  - Stop Docker containers"

# Run the application
run:
	@echo "Starting application..."
	go run main.go

# Build the application
build:
	@echo "Building application..."
	go build -o ecommerce-api main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f ecommerce-api
	go clean

# Apply migrations (we'll implement this later)
migrate-up:
	@echo "Applying migrations..."
	go run main.go migrate up

# Rollback migrations
migrate-down:
	@echo "Rolling back migrations..."
	go run main.go migrate down

# Start with Docker
docker-up:
	@echo "Starting Docker containers..."
	docker-compose up --build

# Stop Docker
docker-down:
	@echo "Stopping Docker containers..."
	docker-compose down