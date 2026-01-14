.PHONY: help run build test clean docker-up docker-down

# Default target - show available commands
help:
	@echo "╔════════════════════════════════════════╗"
	@echo "║   E-commerce Backend - Commands        ║"
	@echo "╚════════════════════════════════════════╝"
	@echo ""
	@echo "📦 Development:"
	@echo "  make run              Run application locally"
	@echo "  make build            Build binary"
	@echo "  make test             Run tests"
	@echo "  make clean            Clean build artifacts"
	@echo ""
	@echo "🐳 Docker:"
	@echo "  make docker-up        Start all containers"
	@echo "  make docker-down      Stop all containers"
	@echo "  make docker-logs      View container logs"
	@echo "  make docker-clean     Remove containers and volumes"
	@echo ""
	@echo "📊 Database:"
	@echo "  make db-connect       Connect to database"
	@echo "  make db-migrate       Run migrations"
	@echo ""

# Run locally (requires local Go)
run:
	@echo "🚀 Starting application..."
	go run main.go

# Build binary
build:
	@echo "🔨 Building application..."
	go build -o ecommerce-backend main.go

# Run tests
test:
	@echo "🧪 Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	rm -f ecommerce-backend main
	go clean

# Docker commands (we'll implement these next)
docker-up:
	@echo "🐳 Starting Docker containers..."
	docker-compose up --build

docker-down:
	@echo "🛑 Stopping Docker containers..."
	docker-compose down

docker-logs:
	@echo "📋 Viewing logs..."
	docker-compose logs -f

docker-clean:
	@echo "🗑️  Removing containers and volumes..."
	docker-compose down -v

# Database helpers
db-connect:
	@echo "🔌 Connecting to database..."
	docker-compose exec postgres psql -U postgres -d ecommerce

db-migrate:
	@echo "🔄 Running migrations..."
	go run main.go migrate