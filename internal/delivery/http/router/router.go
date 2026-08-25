// Package router provides HTTP routing setup.
package router

import (
	"github.com/gin-gonic/gin"

	"emedic-bk/internal/delivery/http/handler"
	"emedic-bk/internal/delivery/http/middleware"
)

// Router holds all route handlers.
type Router struct {
	engine *gin.Engine
}

// NewRouter creates a new router with all routes configured.
func NewRouter(
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	courseHandler *handler.CourseHandler,
	moduleHandler *handler.ModuleHandler,
	lessonHandler *handler.LessonHandler,
	contentHandler *handler.ContentHandler,
	enrollmentHandler *handler.EnrollmentHandler,
	subscriptionHandler *handler.SubscriptionHandler,
	paymentHandler *handler.PaymentHandler,
	qnaHandler *handler.QnAHandler,
	discussionHandler *handler.DiscussionHandler,
	progressHandler *handler.ProgressHandler,
	healthHandler *handler.HealthHandler,
	adminHandler *handler.AdminHandler,
	authMiddleware *middleware.AuthMiddleware,
	roleMiddleware *middleware.RoleMiddleware,
) *Router {
	engine := gin.Default()

	// Global middleware
	engine.Use(middleware.CORS())
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recovery())
	engine.Use(middleware.RequestID())

	// Health check
	engine.GET("/health", healthHandler.Health)

	// API v1
	v1 := engine.Group("/api/v1")
	{
		// Auth routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
			auth.POST("/password/reset", authHandler.RequestPasswordReset)
			auth.POST("/password/confirm", authHandler.ConfirmPasswordReset)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(authMiddleware.Authenticate())
		{
			// Auth
			protected.POST("/auth/logout", authHandler.Logout)

			// Users
			protected.GET("/users/me", userHandler.GetProfile)
			protected.PUT("/users/me", userHandler.UpdateProfile)

			// Courses
			protected.GET("/courses", courseHandler.ListCourses)
			protected.GET("/courses/:id", courseHandler.GetCourse)

			// Modules & Lessons (use nested paths to avoid conflicts)
			protected.GET("/modules", moduleHandler.ListModules) // ?course_id=xxx
			protected.GET("/modules/:id", moduleHandler.GetModule)
			protected.GET("/modules/:id/lessons", lessonHandler.ListLessons)
			protected.GET("/lessons/:id", lessonHandler.GetLesson)

			// Discussion boards (per module; posts + threaded comments)
			protected.GET("/modules/:id/posts", discussionHandler.ListPosts) // ?page=&limit=
			protected.POST("/modules/:id/posts", discussionHandler.CreatePost)
			protected.GET("/posts/:id", discussionHandler.GetPost)
			protected.DELETE("/posts/:id", discussionHandler.DeletePost) // own or admin
			protected.GET("/posts/:id/comments", discussionHandler.ListComments)
			protected.POST("/posts/:id/comments", discussionHandler.CreateComment)
			protected.DELETE("/comments/:id", discussionHandler.DeleteComment) // own or admin

			// Content
			protected.GET("/content/:id/url", contentHandler.GetContentURL)

			// Enrollments
			protected.POST("/enrollments", enrollmentHandler.Enroll)
			protected.GET("/enrollments", enrollmentHandler.ListEnrollments)
			protected.DELETE("/enrollments/:id", enrollmentHandler.Unenroll)

			// Subscriptions (created via payment verification, not directly)
			protected.GET("/subscriptions", subscriptionHandler.List)
			protected.DELETE("/subscriptions/:id", subscriptionHandler.Cancel)

			// Plans & Payments
			protected.GET("/plans", paymentHandler.ListPlans)
			protected.POST("/payments", paymentHandler.InitiatePayment)
			protected.POST("/payments/verify", paymentHandler.VerifyPayment)
			protected.GET("/payments", paymentHandler.ListPayments)

			// Q&A
			protected.POST("/questions", qnaHandler.CreateQuestion)
			protected.GET("/questions", qnaHandler.ListQuestions) // ?lesson_id=xxx
			protected.PUT("/questions/:id", qnaHandler.UpdateQuestion)
			protected.DELETE("/questions/:id", qnaHandler.DeleteQuestion)
			protected.POST("/answers", qnaHandler.CreateAnswer)
			protected.PUT("/answers/:id", qnaHandler.UpdateAnswer)
			protected.DELETE("/answers/:id", qnaHandler.DeleteAnswer)
			protected.POST("/questions/:id/best-answer/:answerId", qnaHandler.MarkBestAnswer)

			// Progress
			protected.POST("/progress", progressHandler.UpdateProgress)
			protected.GET("/progress", progressHandler.ListProgress)
			protected.GET("/progress/:id", progressHandler.GetProgress)
			protected.GET("/progress/course/:id", progressHandler.GetCourseProgress)
		}

		// Admin routes
		admin := v1.Group("/admin")
		admin.Use(authMiddleware.Authenticate())
		admin.Use(roleMiddleware.RequireRole("admin"))
		{
			// Dashboard stats
			admin.GET("/stats", adminHandler.Stats)

			// Users
			admin.GET("/users", userHandler.ListUsers)

			// Courses
			admin.POST("/courses", courseHandler.CreateCourse)
			admin.PUT("/courses/:id", courseHandler.UpdateCourse)
			admin.DELETE("/courses/:id", courseHandler.DeleteCourse)

			// Modules
			admin.POST("/modules", moduleHandler.CreateModule)
			admin.PUT("/modules/:id", moduleHandler.UpdateModule)
			admin.DELETE("/modules/:id", moduleHandler.DeleteModule)
			admin.POST("/modules/:id/regenerate-cover", moduleHandler.RegenerateCover)

			// Discussion moderation
			admin.POST("/posts/:id/pin", discussionHandler.PinPost)
			admin.DELETE("/posts/:id/pin", discussionHandler.UnpinPost)

			// Lessons
			admin.POST("/lessons", lessonHandler.CreateLesson)
			admin.PUT("/lessons/:id", lessonHandler.UpdateLesson)
			admin.DELETE("/lessons/:id", lessonHandler.DeleteLesson)

			// Content
			admin.POST("/content", contentHandler.UploadContent)
			admin.DELETE("/content/:id", contentHandler.DeleteContent)
		}

		// Webhook routes (no auth, but signature verification)
		webhooks := v1.Group("/webhooks")
		{
			webhooks.POST("/paystack", paymentHandler.PaystackWebhook)
		}
	}

	return &Router{engine: engine}
}

// Engine returns the Gin engine.
func (r *Router) Engine() *gin.Engine {
	return r.engine
}
