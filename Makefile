# Project Variables
APP_NAME=zneha-backend
CMD_DIR=cmd/server
MAIN_FILE=$(CMD_DIR)/main.go

# Database Config (from .env or fallback values)
DB_USER?=postgres
DB_PASSWORD?=zneha
DB_NAME?=zneha_backend
DB_HOST?=localhost
DB_PORT?=5432
DB_SSLMODE?=disable

# Tools
GO=go
AIR=air

## ----------- COMMANDS -----------

# Run the app with live-reload using Air
run:
	@echo "🚀 Starting $(APP_NAME) with live-reload..."
	@echo "Note: Make sure 'air' is installed (run 'make install-dev-tools')."
	@$(AIR) -c .air.toml

# Build binary
build:
	@echo "🔨 Building $(APP_NAME)..."
	@$(GO) build -o bin/$(APP_NAME) $(MAIN_FILE)

# Install development tools
install-dev-tools:
	@echo "Installing development tools..."
	@$(GO) install github.com/cosmtrek/air@latest

# Run database migrations
migrate:
	@echo "📦 Running migrations..."
	@$(GO) run scripts/migrate/main.go

# Seed database with dummy data
seed:
	@echo "🌱 Seeding database..."
	@$(GO) run scripts/seed/main.go

# Run tests
test:
	@echo "🧪 Running tests..."
	@$(GO) test ./... -v

# Format & tidy dependencies
fmt:
	@echo "✨ Formatting code..."
	@$(GO) fmt ./...
	@$(GO) mod tidy

# Clean build files
clean:
	@echo "🧹 Cleaning up..."
	@rm -rf bin

# Help (lists all commands)
help:
	@echo "Available commands:"
	@echo "  make run       - Run the server"
	@echo "  make build     - Build binary"
	@echo "  make migrate   - Run migrations"
	@echo "  make seed      - Seed database"
	@echo "  make test      - Run tests"
	@echo "  make fmt       - Format & tidy code"
	@echo "  make clean     - Clean build files"
