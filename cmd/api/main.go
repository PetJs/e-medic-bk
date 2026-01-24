// Package main is the entry point for the API server.
// It bootstraps the HTTP server, loads configuration, and initializes all dependencies.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"emedic-bk/config"
	"emedic-bk/internal/application/usecase/auth"
	"emedic-bk/internal/delivery/http/handler"
	"emedic-bk/internal/delivery/http/middleware"
	"emedic-bk/internal/delivery/http/router"
	infraAuth "emedic-bk/internal/infrastructure/auth"
	"emedic-bk/internal/infrastructure/persistence/postgres"
	"emedic-bk/pkg/uid"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Create context
	ctx := context.Background()

	// Connect to database
	db, err := postgres.NewDB(ctx, cfg.Database.DatabaseURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Connected to database")

	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)

	// Initialize infrastructure services
	hasher := infraAuth.NewBcryptHasher(12)
	tokenGen := infraAuth.NewJWTGenerator(
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		cfg.JWT.AccessTTL,
		cfg.JWT.RefreshTTL,
	)
	idGen := uid.NewGenerator()

	// Initialize use cases
	registerUC := auth.NewRegisterUseCase(userRepo, hasher, tokenGen, idGen)
	loginUC := auth.NewLoginUseCase(userRepo, hasher, tokenGen)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(registerUC, loginUC)

	// Create placeholder handlers for routes that require them
	userHandler := handler.NewUserHandler()
	courseHandler := handler.NewCourseHandler()
	moduleHandler := handler.NewModuleHandler()
	lessonHandler := handler.NewLessonHandler()
	contentHandler := handler.NewContentHandler()
	enrollmentHandler := handler.NewEnrollmentHandler()
	subscriptionHandler := handler.NewSubscriptionHandler()
	paymentHandler := handler.NewPaymentHandler()
	qnaHandler := handler.NewQnAHandler()
	progressHandler := handler.NewProgressHandler()
	healthHandler := handler.NewHealthHandler()

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(tokenGen)
	roleMiddleware := middleware.NewRoleMiddleware()

	// Setup router
	r := router.NewRouter(
		authHandler,
		userHandler,
		courseHandler,
		moduleHandler,
		lessonHandler,
		contentHandler,
		enrollmentHandler,
		subscriptionHandler,
		paymentHandler,
		qnaHandler,
		progressHandler,
		healthHandler,
		authMiddleware,
		roleMiddleware,
	)

	// Create HTTP server
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      r.Engine(),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Server starting on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
