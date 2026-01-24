# eMedic Backend - Folder Structure

A production-ready Go backend for a tertiary-level educational/tutorial platform following Clean Architecture principles.

---

## Complete Folder Tree

```
emedic-bk/
├── cmd/
│   ├── api/
│   │   └── main.go                          # Application entry point, bootstraps the HTTP server
│   ├── migrate/
│   │   └── main.go                          # CLI tool for running database migrations
│   └── worker/
│       └── main.go                          # Entry point for background job workers
│
├── internal/
│   ├── domain/                              # Core business entities and interfaces (innermost layer)
│   │   ├── entity/
│   │   │   ├── user.go                      # User entity with roles (student, admin)
│   │   │   ├── course.go                    # Course entity
│   │   │   ├── module.go                    # Module entity (can be free or premium)
│   │   │   ├── lesson.go                    # Lesson entity (contains content references)
│   │   │   ├── content.go                   # Content entity (PDFs, videos metadata)
│   │   │   ├── enrollment.go                # User-course enrollment entity
│   │   │   ├── subscription.go              # User subscription entity
│   │   │   ├── payment.go                   # Payment transaction entity
│   │   │   ├── question.go                  # Q&A question entity
│   │   │   ├── answer.go                    # Q&A answer entity
│   │   │   └── progress.go                  # User progress tracking entity
│   │   │
│   │   ├── repository/                      # Repository interfaces (ports)
│   │   │   ├── user_repository.go           # User data access interface
│   │   │   ├── course_repository.go         # Course data access interface
│   │   │   ├── module_repository.go         # Module data access interface
│   │   │   ├── lesson_repository.go         # Lesson data access interface
│   │   │   ├── content_repository.go        # Content metadata access interface
│   │   │   ├── enrollment_repository.go     # Enrollment data access interface
│   │   │   ├── subscription_repository.go   # Subscription data access interface
│   │   │   ├── payment_repository.go        # Payment data access interface
│   │   │   ├── question_repository.go       # Question data access interface
│   │   │   ├── answer_repository.go         # Answer data access interface
│   │   │   └── progress_repository.go       # Progress data access interface
│   │   │
│   │   ├── service/                         # Domain service interfaces
│   │   │   ├── auth_service.go              # Authentication service interface
│   │   │   ├── storage_service.go           # Object storage service interface
│   │   │   └── payment_service.go           # Payment gateway service interface
│   │   │
│   │   └── valueobject/                     # Value objects (immutable domain primitives)
│   │       ├── email.go                     # Email value object with validation
│   │       ├── money.go                     # Money value object for payments
│   │       ├── role.go                      # User role enum (student, admin)
│   │       ├── content_type.go              # Content type enum (PDF, video)
│   │       └── subscription_status.go       # Subscription status enum
│   │
│   ├── application/                         # Application layer (use cases / business logic)
│   │   ├── usecase/
│   │   │   ├── auth/
│   │   │   │   ├── register.go              # User registration use case
│   │   │   │   ├── login.go                 # User login use case
│   │   │   │   ├── logout.go                # User logout use case
│   │   │   │   ├── refresh_token.go         # JWT token refresh use case
│   │   │   │   └── reset_password.go        # Password reset use case
│   │   │   │
│   │   │   ├── user/
│   │   │   │   ├── get_profile.go           # Get user profile use case
│   │   │   │   ├── update_profile.go        # Update user profile use case
│   │   │   │   └── list_users.go            # List users (admin) use case
│   │   │   │
│   │   │   ├── course/
│   │   │   │   ├── create_course.go         # Create course use case
│   │   │   │   ├── update_course.go         # Update course use case
│   │   │   │   ├── delete_course.go         # Delete course use case
│   │   │   │   ├── get_course.go            # Get course details use case
│   │   │   │   └── list_courses.go          # List all courses use case
│   │   │   │
│   │   │   ├── module/
│   │   │   │   ├── create_module.go         # Create module use case
│   │   │   │   ├── update_module.go         # Update module use case
│   │   │   │   ├── delete_module.go         # Delete module use case
│   │   │   │   └── list_modules.go          # List modules by course use case
│   │   │   │
│   │   │   ├── lesson/
│   │   │   │   ├── create_lesson.go         # Create lesson use case
│   │   │   │   ├── update_lesson.go         # Update lesson use case
│   │   │   │   ├── delete_lesson.go         # Delete lesson use case
│   │   │   │   ├── get_lesson.go            # Get lesson with content use case
│   │   │   │   └── list_lessons.go          # List lessons by module use case
│   │   │   │
│   │   │   ├── content/
│   │   │   │   ├── upload_content.go        # Upload PDF/video content use case
│   │   │   │   ├── get_content_url.go       # Get signed URL for content access use case
│   │   │   │   └── delete_content.go        # Delete content use case
│   │   │   │
│   │   │   ├── enrollment/
│   │   │   │   ├── enroll_course.go         # Enroll user in course use case
│   │   │   │   ├── unenroll_course.go       # Unenroll user from course use case
│   │   │   │   └── list_enrollments.go      # List user enrollments use case
│   │   │   │
│   │   │   ├── subscription/
│   │   │   │   ├── create_subscription.go   # Create subscription use case
│   │   │   │   ├── cancel_subscription.go   # Cancel subscription use case
│   │   │   │   ├── renew_subscription.go    # Renew subscription use case
│   │   │   │   └── check_access.go          # Check premium access use case
│   │   │   │
│   │   │   ├── payment/
│   │   │   │   ├── initiate_payment.go      # Initiate payment use case
│   │   │   │   ├── verify_payment.go        # Verify payment callback use case
│   │   │   │   └── list_payments.go         # List payment history use case
│   │   │   │
│   │   │   ├── qna/
│   │   │   │   ├── create_question.go       # Create question use case
│   │   │   │   ├── update_question.go       # Update question use case
│   │   │   │   ├── delete_question.go       # Delete question use case
│   │   │   │   ├── list_questions.go        # List questions by lesson use case
│   │   │   │   ├── create_answer.go         # Create answer use case
│   │   │   │   ├── update_answer.go         # Update answer use case
│   │   │   │   ├── delete_answer.go         # Delete answer use case
│   │   │   │   └── mark_best_answer.go      # Mark best answer use case
│   │   │   │
│   │   │   └── progress/
│   │   │       ├── update_progress.go       # Update lesson progress use case
│   │   │       ├── get_progress.go          # Get user progress use case
│   │   │       └── get_course_progress.go   # Get course completion stats use case
│   │   │
│   │   ├── dto/                             # Data Transfer Objects for use cases
│   │   │   ├── auth_dto.go                  # Auth request/response DTOs
│   │   │   ├── user_dto.go                  # User request/response DTOs
│   │   │   ├── course_dto.go                # Course request/response DTOs
│   │   │   ├── module_dto.go                # Module request/response DTOs
│   │   │   ├── lesson_dto.go                # Lesson request/response DTOs
│   │   │   ├── content_dto.go               # Content request/response DTOs
│   │   │   ├── enrollment_dto.go            # Enrollment request/response DTOs
│   │   │   ├── subscription_dto.go          # Subscription request/response DTOs
│   │   │   ├── payment_dto.go               # Payment request/response DTOs
│   │   │   ├── qna_dto.go                   # Q&A request/response DTOs
│   │   │   └── progress_dto.go              # Progress request/response DTOs
│   │   │
│   │   └── port/                            # Application ports (secondary interfaces)
│   │       ├── hasher.go                    # Password hashing interface
│   │       ├── token_generator.go           # JWT token generator interface
│   │       ├── mailer.go                    # Email sending interface
│   │       └── id_generator.go              # Unique ID generator interface
│   │
│   ├── infrastructure/                      # Infrastructure layer (adapters)
│   │   ├── persistence/
│   │   │   ├── postgres/
│   │   │   │   ├── connection.go            # PostgreSQL connection pool setup
│   │   │   │   ├── user_repository.go       # User repository implementation
│   │   │   │   ├── course_repository.go     # Course repository implementation
│   │   │   │   ├── module_repository.go     # Module repository implementation
│   │   │   │   ├── lesson_repository.go     # Lesson repository implementation
│   │   │   │   ├── content_repository.go    # Content repository implementation
│   │   │   │   ├── enrollment_repository.go # Enrollment repository implementation
│   │   │   │   ├── subscription_repository.go # Subscription repository implementation
│   │   │   │   ├── payment_repository.go    # Payment repository implementation
│   │   │   │   ├── question_repository.go   # Question repository implementation
│   │   │   │   ├── answer_repository.go     # Answer repository implementation
│   │   │   │   └── progress_repository.go   # Progress repository implementation
│   │   │   │
│   │   │   └── sqlc/                        # SQLC generated code (if using sqlc)
│   │   │       ├── db.go                    # SQLC database interface
│   │   │       ├── models.go                # SQLC generated models
│   │   │       └── queries.sql.go           # SQLC generated query functions
│   │   │
│   │   ├── storage/
│   │   │   └── s3/
│   │   │       ├── client.go                # S3 client initialization
│   │   │       └── storage_service.go       # S3 storage service implementation
│   │   │
│   │   ├── payment/
│   │   │   ├── stripe/
│   │   │   │   └── stripe_service.go        # Stripe payment service implementation
│   │   │   └── paystack/
│   │   │       └── paystack_service.go      # Paystack payment service implementation
│   │   │
│   │   ├── auth/
│   │   │   ├── jwt.go                       # JWT token generator implementation
│   │   │   └── bcrypt.go                    # Bcrypt password hasher implementation
│   │   │
│   │   ├── mail/
│   │   │   └── smtp.go                      # SMTP mailer implementation
│   │   │
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── client.go                # Redis client initialization
│   │   │       └── cache.go                 # Redis cache implementation
│   │   │
│   │   └── queue/
│   │       └── redis/
│   │           ├── publisher.go             # Redis job queue publisher
│   │           └── consumer.go              # Redis job queue consumer
│   │
│   ├── delivery/                            # Delivery layer (primary adapters)
│   │   └── http/
│   │       ├── router/
│   │       │   └── router.go                # Main router setup and route registration
│   │       │
│   │       ├── handler/
│   │       │   ├── auth_handler.go          # Authentication HTTP handlers
│   │       │   ├── user_handler.go          # User HTTP handlers
│   │       │   ├── course_handler.go        # Course HTTP handlers
│   │       │   ├── module_handler.go        # Module HTTP handlers
│   │       │   ├── lesson_handler.go        # Lesson HTTP handlers
│   │       │   ├── content_handler.go       # Content HTTP handlers
│   │       │   ├── enrollment_handler.go    # Enrollment HTTP handlers
│   │       │   ├── subscription_handler.go  # Subscription HTTP handlers
│   │       │   ├── payment_handler.go       # Payment HTTP handlers (including webhooks)
│   │       │   ├── qna_handler.go           # Q&A HTTP handlers
│   │       │   ├── progress_handler.go      # Progress HTTP handlers
│   │       │   └── health_handler.go        # Health check HTTP handler
│   │       │
│   │       ├── middleware/
│   │       │   ├── auth.go                  # JWT authentication middleware
│   │       │   ├── role.go                  # Role-based authorization middleware
│   │       │   ├── subscription.go          # Premium content access middleware
│   │       │   ├── cors.go                  # CORS middleware
│   │       │   ├── logger.go                # Request logging middleware
│   │       │   ├── recovery.go              # Panic recovery middleware
│   │       │   ├── ratelimit.go             # Rate limiting middleware
│   │       │   └── requestid.go             # Request ID middleware
│   │       │
│   │       ├── response/
│   │       │   ├── response.go              # Standardized HTTP response helpers
│   │       │   └── error.go                 # HTTP error response helpers
│   │       │
│   │       └── validator/
│   │           └── validator.go             # Request validation helpers
│   │
│   ├── worker/                              # Background job workers
│   │   ├── processor.go                     # Job processor initialization
│   │   ├── subscription_expiry.go           # Subscription expiry check job
│   │   ├── email_notification.go            # Email notification job
│   │   └── content_cleanup.go               # Orphaned content cleanup job
│   │
│   └── shared/                              # Shared utilities within internal
│       ├── errors/
│       │   ├── errors.go                    # Custom application errors
│       │   └── codes.go                     # Error codes constants
│       │
│       ├── pagination/
│       │   └── pagination.go                # Pagination utilities
│       │
│       └── context/
│           └── context.go                   # Context key helpers
│
├── pkg/                                     # Public reusable packages
│   ├── logger/
│   │   └── logger.go                        # Structured logging package (zap/zerolog wrapper)
│   │
│   ├── validator/
│   │   └── validator.go                     # Input validation package
│   │
│   └── uid/
│       └── uid.go                           # Unique ID generation package (UUID/ULID)
│
├── config/                                  # Configuration management
│   ├── config.go                            # Config struct and loader
│   └── config.yaml                          # Default configuration file (example)
│
├── migrations/                              # Database migrations
│   ├── 000001_create_users_table.up.sql     # Create users table migration
│   ├── 000001_create_users_table.down.sql   # Rollback users table migration
│   ├── 000002_create_courses_table.up.sql   # Create courses table migration
│   ├── 000002_create_courses_table.down.sql # Rollback courses table migration
│   ├── 000003_create_modules_table.up.sql   # Create modules table migration
│   ├── 000003_create_modules_table.down.sql # Rollback modules table migration
│   ├── 000004_create_lessons_table.up.sql   # Create lessons table migration
│   ├── 000004_create_lessons_table.down.sql # Rollback lessons table migration
│   ├── 000005_create_contents_table.up.sql  # Create contents table migration
│   ├── 000005_create_contents_table.down.sql # Rollback contents table migration
│   ├── 000006_create_enrollments_table.up.sql # Create enrollments table migration
│   ├── 000006_create_enrollments_table.down.sql # Rollback enrollments table migration
│   ├── 000007_create_subscriptions_table.up.sql # Create subscriptions table migration
│   ├── 000007_create_subscriptions_table.down.sql # Rollback subscriptions table migration
│   ├── 000008_create_payments_table.up.sql  # Create payments table migration
│   ├── 000008_create_payments_table.down.sql # Rollback payments table migration
│   ├── 000009_create_questions_table.up.sql # Create questions table migration
│   ├── 000009_create_questions_table.down.sql # Rollback questions table migration
│   ├── 000010_create_answers_table.up.sql   # Create answers table migration
│   ├── 000010_create_answers_table.down.sql # Rollback answers table migration
│   ├── 000011_create_progress_table.up.sql  # Create progress table migration
│   └── 000011_create_progress_table.down.sql # Rollback progress table migration
│
├── scripts/                                 # Utility scripts
│   ├── generate_sqlc.sh                     # Script to generate sqlc code
│   ├── run_migrations.sh                    # Script to run database migrations
│   └── seed_data.go                         # Database seeding script
│
├── sql/                                     # SQL query files (for sqlc)
│   ├── queries/
│   │   ├── users.sql                        # User queries
│   │   ├── courses.sql                      # Course queries
│   │   ├── modules.sql                      # Module queries
│   │   ├── lessons.sql                      # Lesson queries
│   │   ├── contents.sql                     # Content queries
│   │   ├── enrollments.sql                  # Enrollment queries
│   │   ├── subscriptions.sql                # Subscription queries
│   │   ├── payments.sql                     # Payment queries
│   │   ├── questions.sql                    # Question queries
│   │   ├── answers.sql                      # Answer queries
│   │   └── progress.sql                     # Progress queries
│   │
│   └── schema/
│       └── schema.sql                       # Complete database schema
│
├── api/                                     # API documentation
│   └── openapi.yaml                         # OpenAPI/Swagger specification
│
├── test/                                    # Test utilities and integration tests
│   ├── integration/
│   │   ├── auth_test.go                     # Authentication integration tests
│   │   ├── course_test.go                   # Course integration tests
│   │   ├── enrollment_test.go               # Enrollment integration tests
│   │   ├── subscription_test.go             # Subscription integration tests
│   │   └── qna_test.go                      # Q&A integration tests
│   │
│   ├── fixtures/
│   │   └── fixtures.go                      # Test data fixtures
│   │
│   └── testutil/
│       ├── database.go                      # Test database helpers
│       └── http.go                          # Test HTTP helpers
│
├── docker/                                  # Docker configuration
│   ├── Dockerfile                           # Production Dockerfile
│   ├── Dockerfile.dev                       # Development Dockerfile
│   └── docker-compose.yml                   # Docker Compose for local development
│
├── .env.example                             # Example environment variables
├── .gitignore                               # Git ignore file
├── .golangci.yml                            # GolangCI-Lint configuration
├── Makefile                                 # Build and development commands
├── sqlc.yaml                                # SQLC configuration
├── go.mod                                   # Go module file
├── go.sum                                   # Go dependencies checksum
└── README.md                                # Project documentation
```

---

## Folder Responsibility Summary

### Root Level Folders

| Folder | Responsibility |
|--------|----------------|
| `cmd/` | Application entry points (binaries) for API server, migrations, and workers |
| `internal/` | Private application code not importable by other projects |
| `pkg/` | Public reusable packages that can be imported by external projects |
| `config/` | Configuration loading and environment-based settings |
| `migrations/` | Database migration files (up/down) using golang-migrate format |
| `scripts/` | Utility scripts for development and automation |
| `sql/` | Raw SQL files for SQLC code generation |
| `api/` | API documentation (OpenAPI/Swagger specs) |
| `test/` | Integration tests, fixtures, and test utilities |
| `docker/` | Docker and containerization configuration |

---

### Internal Package Structure (Clean Architecture Layers)

| Layer | Folder | Responsibility |
|-------|--------|----------------|
| **Domain** | `internal/domain/entity/` | Core business entities (User, Course, Lesson, etc.) |
| **Domain** | `internal/domain/repository/` | Repository interfaces (data access contracts) |
| **Domain** | `internal/domain/service/` | Domain service interfaces (external services contracts) |
| **Domain** | `internal/domain/valueobject/` | Immutable value objects with validation (Email, Money, Role) |
| **Application** | `internal/application/usecase/` | Business logic orchestration organized by feature |
| **Application** | `internal/application/dto/` | Data Transfer Objects for request/response mapping |
| **Application** | `internal/application/port/` | Secondary port interfaces (hasher, token generator, mailer) |
| **Infrastructure** | `internal/infrastructure/persistence/` | Database implementations (PostgreSQL with pgx/sqlc) |
| **Infrastructure** | `internal/infrastructure/storage/` | Object storage implementations (S3-compatible) |
| **Infrastructure** | `internal/infrastructure/payment/` | Payment gateway implementations (Stripe, Paystack) |
| **Infrastructure** | `internal/infrastructure/auth/` | Authentication implementations (JWT, bcrypt) |
| **Infrastructure** | `internal/infrastructure/mail/` | Email service implementations (SMTP) |
| **Infrastructure** | `internal/infrastructure/cache/` | Caching implementations (Redis) |
| **Infrastructure** | `internal/infrastructure/queue/` | Job queue implementations (Redis) |
| **Delivery** | `internal/delivery/http/handler/` | HTTP request handlers for each feature |
| **Delivery** | `internal/delivery/http/middleware/` | HTTP middleware (auth, CORS, logging, rate limiting) |
| **Delivery** | `internal/delivery/http/router/` | Route registration and grouping |
| **Delivery** | `internal/delivery/http/response/` | Standardized HTTP response helpers |
| **Delivery** | `internal/delivery/http/validator/` | Request validation helpers |
| **Workers** | `internal/worker/` | Background job processors (subscriptions, notifications) |
| **Shared** | `internal/shared/` | Shared utilities (errors, pagination, context helpers) |

---

### Feature Organization

| Feature | Use Cases Location | Handler Location |
|---------|-------------------|------------------|
| Authentication | `internal/application/usecase/auth/` | `internal/delivery/http/handler/auth_handler.go` |
| User Management | `internal/application/usecase/user/` | `internal/delivery/http/handler/user_handler.go` |
| Courses | `internal/application/usecase/course/` | `internal/delivery/http/handler/course_handler.go` |
| Modules | `internal/application/usecase/module/` | `internal/delivery/http/handler/module_handler.go` |
| Lessons | `internal/application/usecase/lesson/` | `internal/delivery/http/handler/lesson_handler.go` |
| Content (PDFs/Videos) | `internal/application/usecase/content/` | `internal/delivery/http/handler/content_handler.go` |
| Enrollments | `internal/application/usecase/enrollment/` | `internal/delivery/http/handler/enrollment_handler.go` |
| Subscriptions | `internal/application/usecase/subscription/` | `internal/delivery/http/handler/subscription_handler.go` |
| Payments | `internal/application/usecase/payment/` | `internal/delivery/http/handler/payment_handler.go` |
| Q&A System | `internal/application/usecase/qna/` | `internal/delivery/http/handler/qna_handler.go` |
| Progress Tracking | `internal/application/usecase/progress/` | `internal/delivery/http/handler/progress_handler.go` |

---

## Dependency Flow (Clean Architecture)

```
┌─────────────────────────────────────────────────────────────────┐
│                        Delivery Layer                           │
│  (HTTP Handlers, Middleware, Router)                            │
│  Depends on: Application Layer                                  │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Application Layer                          │
│  (Use Cases, DTOs, Ports)                                       │
│  Depends on: Domain Layer                                       │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                        Domain Layer                             │
│  (Entities, Repository Interfaces, Value Objects)               │
│  Depends on: Nothing (innermost layer)                          │
└─────────────────────────────────────────────────────────────────┘
                              ▲
                              │
┌─────────────────────────────────────────────────────────────────┐
│                    Infrastructure Layer                         │
│  (PostgreSQL, S3, Stripe, Redis, SMTP)                          │
│  Implements: Domain Interfaces                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## Key Design Decisions

1. **Separation by Feature in Use Cases**: Each feature (auth, course, enrollment) has its own subdirectory under `usecase/` for better organization
2. **Single Handler Files**: Each feature has one handler file to keep routes grouped logically
3. **Repository Pattern**: Interfaces in domain, implementations in infrastructure for testability
4. **DTOs in Application Layer**: Clean separation between domain entities and API contracts
5. **Middleware for Access Control**: Separate middleware for auth, roles, and subscription checks
6. **Background Workers**: Separate entry point for job processing (subscription expiry, notifications)
7. **SQLC for Type-Safe Queries**: SQL files in `sql/` generate type-safe Go code in `infrastructure/persistence/sqlc/`
8. **Integration Tests Separate**: Full test suite in `test/` directory with fixtures and utilities
