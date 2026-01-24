.PHONY: all build run test clean migrate-up migrate-down sqlc lint docker-up docker-down

# Binary name
BINARY_NAME=emedic-api

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

# Build the application
build:
	@echo "Building..."
	@go build -o $(GOBIN)/$(BINARY_NAME) ./cmd/api

# Run the application
run:
	@go run ./cmd/api

# Run with hot reload (requires air)
dev:
	@air -c .air.toml

# Run tests
test:
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

# Clean build files
clean:
	@rm -rf $(GOBIN)
	@rm -f coverage.out coverage.html

# Run database migrations up
migrate-up:
	@migrate -path ./migrations -database "$(DATABASE_URL)" up

# Run database migrations down
migrate-down:
	@migrate -path ./migrations -database "$(DATABASE_URL)" down

# Generate SQLC code
sqlc:
	@sqlc generate

# Run linter
lint:
	@golangci-lint run ./...

# Start docker services
docker-up:
	@docker-compose -f docker/docker-compose.yml up -d

# Stop docker services
docker-down:
	@docker-compose -f docker/docker-compose.yml down

# Build docker image
docker-build:
	@docker build -f docker/Dockerfile -t emedic-api:latest .

# Install development tools
install-tools:
	@go install github.com/cosmtrek/air@latest
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Generate all
generate: sqlc
	@echo "Generated all code"

# Help
help:
	@echo "Available targets:"
	@echo "  build          - Build the application"
	@echo "  run            - Run the application"
	@echo "  dev            - Run with hot reload"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage"
	@echo "  clean          - Clean build files"
	@echo "  migrate-up     - Run migrations up"
	@echo "  migrate-down   - Run migrations down"
	@echo "  sqlc           - Generate SQLC code"
	@echo "  lint           - Run linter"
	@echo "  docker-up      - Start docker services"
	@echo "  docker-down    - Stop docker services"
	@echo "  docker-build   - Build docker image"
	@echo "  install-tools  - Install development tools"
