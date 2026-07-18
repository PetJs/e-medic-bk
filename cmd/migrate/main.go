// Package main provides a CLI tool for running database migrations.
//
// Usage:
//
//	go run ./cmd/migrate            # apply all pending up migrations
//	go run ./cmd/migrate -down      # roll back the most recent migration
//
// The database is taken from DATABASE_URL (or the DB_* variables).
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	"emedic-bk/config"
)

func main() {
	down := flag.Bool("down", false, "roll back the most recent migration")
	dir := flag.String("dir", "migrations", "migrations directory")
	flag.Parse()

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

	if _, err := conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version BIGINT PRIMARY KEY,
			applied_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
		)
	`); err != nil {
		log.Fatalf("Failed to create schema_migrations table: %v", err)
	}

	if *down {
		if err := rollback(ctx, conn, *dir); err != nil {
			log.Fatalf("Rollback failed: %v", err)
		}
		return
	}

	if err := migrateUp(ctx, conn, *dir); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}

type migration struct {
	version int64
	name    string
	path    string
}

func listMigrations(dir, suffix string) ([]migration, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var migrations []migration
	for _, e := range entries {
		name := e.Name()
		if !strings.HasSuffix(name, suffix) {
			continue
		}
		parts := strings.SplitN(name, "_", 2)
		version, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("bad migration filename %q: %w", name, err)
		}
		migrations = append(migrations, migration{
			version: version,
			name:    strings.TrimSuffix(name, suffix),
			path:    filepath.Join(dir, name),
		})
	}
	sort.Slice(migrations, func(i, j int) bool { return migrations[i].version < migrations[j].version })
	return migrations, nil
}

func appliedVersions(ctx context.Context, conn *pgx.Conn) (map[int64]bool, error) {
	rows, err := conn.Query(ctx, `SELECT version FROM schema_migrations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applied := map[int64]bool{}
	for rows.Next() {
		var v int64
		if err := rows.Scan(&v); err != nil {
			return nil, err
		}
		applied[v] = true
	}
	return applied, rows.Err()
}

func migrateUp(ctx context.Context, conn *pgx.Conn, dir string) error {
	migrations, err := listMigrations(dir, ".up.sql")
	if err != nil {
		return err
	}
	applied, err := appliedVersions(ctx, conn)
	if err != nil {
		return err
	}

	ran := 0
	for _, m := range migrations {
		if applied[m.version] {
			continue
		}
		sql, err := os.ReadFile(m.path)
		if err != nil {
			return err
		}

		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}
		if _, err := tx.Exec(ctx, string(sql)); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("migration %s: %w", m.name, err)
		}
		if _, err := tx.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, m.version); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
		if err := tx.Commit(ctx); err != nil {
			return err
		}
		log.Printf("Applied %s", m.name)
		ran++
	}

	if ran == 0 {
		log.Println("Database is up to date")
	} else {
		log.Printf("Applied %d migration(s)", ran)
	}
	return nil
}

func rollback(ctx context.Context, conn *pgx.Conn, dir string) error {
	applied, err := appliedVersions(ctx, conn)
	if err != nil {
		return err
	}
	if len(applied) == 0 {
		log.Println("Nothing to roll back")
		return nil
	}

	var latest int64
	for v := range applied {
		if v > latest {
			latest = v
		}
	}

	migrations, err := listMigrations(dir, ".down.sql")
	if err != nil {
		return err
	}
	for _, m := range migrations {
		if m.version != latest {
			continue
		}
		sql, err := os.ReadFile(m.path)
		if err != nil {
			return err
		}

		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}
		if _, err := tx.Exec(ctx, string(sql)); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("rollback %s: %w", m.name, err)
		}
		if _, err := tx.Exec(ctx, `DELETE FROM schema_migrations WHERE version = $1`, m.version); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
		if err := tx.Commit(ctx); err != nil {
			return err
		}
		log.Printf("Rolled back %s", m.name)
		return nil
	}
	return fmt.Errorf("no .down.sql found for version %d", latest)
}
