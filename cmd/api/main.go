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

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"emedic-bk/config"
	"emedic-bk/internal/application/usecase/admin"
	"emedic-bk/internal/application/usecase/auth"
	contentUC "emedic-bk/internal/application/usecase/content"
	"emedic-bk/internal/application/usecase/course"
	"emedic-bk/internal/application/usecase/lesson"
	"emedic-bk/internal/application/usecase/module"
	"emedic-bk/internal/application/usecase/payment"
	progressUC "emedic-bk/internal/application/usecase/progress"
	"emedic-bk/internal/application/usecase/subscription"
	"emedic-bk/internal/application/usecase/user"
	"emedic-bk/internal/delivery/http/handler"
	"emedic-bk/internal/delivery/http/middleware"
	"emedic-bk/internal/delivery/http/router"
	infraAuth "emedic-bk/internal/infrastructure/auth"
	"emedic-bk/internal/infrastructure/payment/paystack"
	"emedic-bk/internal/infrastructure/persistence/postgres"
	s3storage "emedic-bk/internal/infrastructure/storage/s3"
	"emedic-bk/pkg/uid"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	cfg := config.Load()
	if cfg.Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

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
	courseRepo := postgres.NewCourseRepository(db)
	moduleRepo := postgres.NewModuleRepository(db)
	lessonRepo := postgres.NewLessonRepository(db)
	contentRepo := postgres.NewContentRepository(db)
	subscriptionRepo := postgres.NewSubscriptionRepository(db)
	paymentRepo := postgres.NewPaymentRepository(db)
	progressRepo := postgres.NewProgressRepository(db)

	// Initialize infrastructure services
	hasher := infraAuth.NewBcryptHasher(12)
	tokenGen := infraAuth.NewJWTGenerator(
		cfg.JWT.AccessSecret,
		cfg.JWT.RefreshSecret,
		cfg.JWT.AccessTTL,
		cfg.JWT.RefreshTTL,
	)
	idGen := uid.NewGenerator()
	paystackSvc := paystack.NewPaymentService(cfg.Paystack.SecretKey)
	storageSvc := s3storage.NewStorageService(s3storage.NewClient(s3storage.Config{
		Endpoint:        cfg.S3.Endpoint,
		Region:          cfg.S3.Region,
		AccessKeyID:     cfg.S3.AccessKeyID,
		SecretAccessKey: cfg.S3.SecretAccessKey,
		BucketName:      cfg.S3.BucketName,
		UsePathStyle:    cfg.S3.UsePathStyle,
	}))

	// Initialize use cases
	registerUC := auth.NewRegisterUseCase(userRepo, hasher, tokenGen, idGen)
	loginUC := auth.NewLoginUseCase(userRepo, hasher, tokenGen)
	refreshUC := auth.NewRefreshTokenUseCase(userRepo, tokenGen)
	getProfileUC := user.NewGetProfileUseCase(userRepo)
	updateProfileUC := user.NewUpdateProfileUseCase(userRepo)
	listUsersUC := user.NewListUsersUseCase(userRepo)

	// Course use cases
	createCourseUC := course.NewCreateCourseUseCase(courseRepo, idGen)
	updateCourseUC := course.NewUpdateCourseUseCase(courseRepo)
	deleteCourseUC := course.NewDeleteCourseUseCase(courseRepo, moduleRepo, lessonRepo, contentRepo, storageSvc)
	getCourseUC := course.NewGetCourseUseCase(courseRepo, moduleRepo)
	listCoursesUC := course.NewListCoursesUseCase(courseRepo)

	// Module use cases
	createModuleUC := module.NewCreateModuleUseCase(moduleRepo, courseRepo, idGen)
	updateModuleUC := module.NewUpdateModuleUseCase(moduleRepo)
	deleteModuleUC := module.NewDeleteModuleUseCase(moduleRepo, lessonRepo, contentRepo, storageSvc)
	listModulesUC := module.NewListModulesUseCase(moduleRepo)
	getModuleUC := module.NewGetModuleUseCase(moduleRepo)

	// Lesson use cases
	createLessonUC := lesson.NewCreateLessonUseCase(lessonRepo, moduleRepo, idGen)
	updateLessonUC := lesson.NewUpdateLessonUseCase(lessonRepo)
	deleteLessonUC := lesson.NewDeleteLessonUseCase(lessonRepo, contentRepo, storageSvc)
	getLessonUC := lesson.NewGetLessonUseCase(lessonRepo, moduleRepo, contentRepo, subscriptionRepo)
	listLessonsUC := lesson.NewListLessonsUseCase(lessonRepo, moduleRepo)

	// Subscription use cases
	activateSubUC := subscription.NewActivateSubscriptionUseCase(subscriptionRepo, idGen)
	cancelSubUC := subscription.NewCancelSubscriptionUseCase(subscriptionRepo)
	listSubsUC := subscription.NewListSubscriptionsUseCase(subscriptionRepo)

	// Payment use cases
	plan := payment.PlanDetails{Amount: cfg.Plan.Amount, Currency: cfg.Plan.Currency}
	paymentCallbackURL := cfg.Server.FrontendURL + "/payment/callback"
	initiatePaymentUC := payment.NewInitiatePaymentUseCase(paymentRepo, userRepo, paystackSvc, idGen, plan, paymentCallbackURL)
	verifyPaymentUC := payment.NewVerifyPaymentUseCase(paymentRepo, paystackSvc, activateSubUC)
	listPaymentsUC := payment.NewListPaymentsUseCase(paymentRepo)

	// Content use cases
	uploadContentUC := contentUC.NewUploadContentUseCase(contentRepo, lessonRepo, storageSvc, idGen)
	getContentURLUC := contentUC.NewGetContentURLUseCase(contentRepo, lessonRepo, moduleRepo, subscriptionRepo, storageSvc)
	deleteContentUC := contentUC.NewDeleteContentUseCase(contentRepo, storageSvc)

	// Progress use cases
	updateProgressUC := progressUC.NewUpdateProgressUseCase(progressRepo, lessonRepo, idGen)
	getProgressUC := progressUC.NewGetProgressUseCase(progressRepo)
	listProgressUC := progressUC.NewListProgressUseCase(progressRepo)
	courseProgressUC := progressUC.NewGetCourseProgressUseCase(progressRepo)

	// Admin use cases
	statsUC := admin.NewGetStatsUseCase(userRepo, subscriptionRepo, paymentRepo, cfg.Plan.Currency)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(registerUC, loginUC, refreshUC)
	userHandler := handler.NewUserHandler(getProfileUC, updateProfileUC, listUsersUC)
	courseHandler := handler.NewCourseHandler(createCourseUC, updateCourseUC, deleteCourseUC, getCourseUC, listCoursesUC)
	moduleHandler := handler.NewModuleHandler(createModuleUC, updateModuleUC, deleteModuleUC, listModulesUC, getModuleUC)
	lessonHandler := handler.NewLessonHandler(createLessonUC, updateLessonUC, deleteLessonUC, getLessonUC, listLessonsUC)
	subscriptionHandler := handler.NewSubscriptionHandler(listSubsUC, cancelSubUC)
	paymentHandler := handler.NewPaymentHandler(initiatePaymentUC, verifyPaymentUC, listPaymentsUC, paystackSvc, plan)
	adminHandler := handler.NewAdminHandler(statsUC)
	contentHandler := handler.NewContentHandler(uploadContentUC, getContentURLUC, deleteContentUC)

	progressHandler := handler.NewProgressHandler(updateProgressUC, getProgressUC, listProgressUC, courseProgressUC)

	// Create placeholder handlers for routes that require them
	enrollmentHandler := handler.NewEnrollmentHandler()
	qnaHandler := handler.NewQnAHandler()
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
		adminHandler,
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
