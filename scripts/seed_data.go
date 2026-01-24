// Package main provides database seeding for development.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	ctx := context.Background()

	// Get database URL from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Connect to database
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	fmt.Println("Seeding database...")

	// Create admin user
	adminPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	_, err = pool.Exec(ctx, `
		INSERT INTO users (email, password_hash, first_name, last_name, role)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (email) DO NOTHING
	`, "admin@emedic.com", string(adminPassword), "Admin", "User", "admin")
	if err != nil {
		log.Printf("Warning: Failed to create admin user: %v", err)
	}

	// Create sample student
	studentPassword, _ := bcrypt.GenerateFromPassword([]byte("student123"), bcrypt.DefaultCost)
	_, err = pool.Exec(ctx, `
		INSERT INTO users (email, password_hash, first_name, last_name, role)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (email) DO NOTHING
	`, "student@emedic.com", string(studentPassword), "Sample", "Student", "student")
	if err != nil {
		log.Printf("Warning: Failed to create student user: %v", err)
	}

	// Get admin user ID
	var adminID string
	err = pool.QueryRow(ctx, "SELECT id FROM users WHERE email = $1", "admin@emedic.com").Scan(&adminID)
	if err != nil {
		log.Fatalf("Failed to get admin user: %v", err)
	}

	// Create sample course
	_, err = pool.Exec(ctx, `
		INSERT INTO courses (title, description, slug, author_id, is_published, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $6)
		ON CONFLICT (slug) DO NOTHING
	`, "Introduction to Medical Science", "A comprehensive introduction to medical science fundamentals.", "intro-to-medical-science", adminID, true, time.Now())
	if err != nil {
		log.Printf("Warning: Failed to create sample course: %v", err)
	}

	fmt.Println("Database seeding completed!")
	fmt.Println("")
	fmt.Println("Sample users:")
	fmt.Println("  Admin:   admin@emedic.com / admin123")
	fmt.Println("  Student: student@emedic.com / student123")
}
