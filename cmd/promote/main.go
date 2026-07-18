// Package main promotes a user to the admin role.
//
// Usage:
//
//	go run ./cmd/promote -email someone@example.com
//	go run ./cmd/promote -email someone@example.com -role student   # demote back
package main

import (
	"context"
	"flag"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	"emedic-bk/config"
)

func main() {
	email := flag.String("email", "", "email of the user to update")
	role := flag.String("role", "admin", "role to assign (admin or student)")
	flag.Parse()

	if *email == "" {
		log.Fatal("usage: go run ./cmd/promote -email someone@example.com [-role admin|student]")
	}
	if *role != "admin" && *role != "student" {
		log.Fatalf("invalid role %q: must be admin or student", *role)
	}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	cfg := config.Load()

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, cfg.Database.DatabaseURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	tag, err := conn.Exec(ctx, `UPDATE users SET role = $1, updated_at = NOW() WHERE email = $2`, *role, *email)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	if tag.RowsAffected() == 0 {
		log.Fatalf("No user found with email %s", *email)
	}
	log.Printf("%s is now %s", *email, *role)
}
