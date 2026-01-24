# eMedic Backend

A production-ready Go backend for a tertiary-level educational/tutorial platform.

## Features

- **Authentication & Authorization**: JWT-based authentication with role-based access control (student, admin)
- **Course Management**: Create, update, and manage courses with modules and lessons
- **Content Delivery**: Support for PDFs and videos stored in S3-compatible storage
- **Subscription System**: Premium module access with Stripe/Paystack payment integration
- **Q&A System**: Questions and answers for each lesson
- **Progress Tracking**: Track user learning progress per lesson and course

## Architecture

This project follows **Clean Architecture / Hexagonal Architecture** principles:

```
├── cmd/                    # Application entry points
├── internal/
│   ├── domain/            # Core business logic (entities, interfaces)
│   ├── application/       # Use cases and DTOs
│   ├── infrastructure/    # External services (DB, S3, payments)
│   └── delivery/          # HTTP handlers, middleware, routing
├── pkg/                   # Public reusable packages
├── config/                # Configuration management
├── migrations/            # Database migrations
└── test/                  # Integration tests
```

## Tech Stack

- **Language**: Go 1.22+
- **HTTP Framework**: Gin
- **Database**: PostgreSQL with pgx/sqlc
- **Cache/Queue**: Redis
- **Object Storage**: S3-compatible (AWS S3, MinIO)
- **Payments**: Stripe, Paystack

## Getting Started

### Prerequisites

- Go 1.22+
- PostgreSQL 16+
- Redis 7+
- Docker & Docker Compose (optional)

### Quick Start with Docker

```bash
# Start all services
make docker-up

# Run migrations
make migrate-up

# Start the API (with hot reload)
make dev
```

### Manual Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and configure
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run migrations:
   ```bash
   make migrate-up
   ```
5. Start the server:
   ```bash
   make run
   ```

## API Documentation

API documentation is available at `/api/openapi.yaml` (OpenAPI 3.0 spec).

## Development

### Available Commands

```bash
make build          # Build the application
make run            # Run the application
make dev            # Run with hot reload
make test           # Run tests
make test-coverage  # Run tests with coverage
make lint           # Run linter
make sqlc           # Generate SQLC code
make migrate-up     # Run migrations
make migrate-down   # Rollback migrations
```

### Project Structure

See `FOLDER_STRUCTURE.md` for detailed folder structure documentation.

## License

MIT
