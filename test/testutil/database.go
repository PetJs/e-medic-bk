// Package testutil provides test utilities.
package testutil

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupTestDB creates a test database connection.
func SetupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	// TODO: Create test database connection
	// TODO: Run migrations
	return nil
}

// TeardownTestDB cleans up the test database.
func TeardownTestDB(t *testing.T, pool *pgxpool.Pool) {
	t.Helper()

	if pool != nil {
		pool.Close()
	}
}

// TruncateTables truncates all tables for test isolation.
func TruncateTables(ctx context.Context, pool *pgxpool.Pool) error {
	tables := []string{
		"progress", "answers", "questions", "payments",
		"subscriptions", "enrollments", "contents",
		"lessons", "modules", "courses", "users",
	}

	for _, table := range tables {
		_, err := pool.Exec(ctx, "TRUNCATE "+table+" CASCADE")
		if err != nil {
			return err
		}
	}

	return nil
}
