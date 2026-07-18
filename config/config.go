// Package config provides configuration management.
package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds all application configuration.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
	S3       S3Config
	Stripe   StripeConfig
	Paystack PaystackConfig
	SMTP     SMTPConfig
	Plan     PlanConfig
}

// ServerConfig holds HTTP server configuration.
type ServerConfig struct {
	Port         string
	Environment  string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	// FrontendURL is used for payment callback redirects and CORS.
	FrontendURL string
}

// PlanConfig holds the subscription plan settings.
type PlanConfig struct {
	// Amount in the smallest currency unit (kobo for NGN).
	Amount   int64
	Currency string
}

// DatabaseConfig holds database configuration.
type DatabaseConfig struct {
	// URL, when set (e.g. a Neon connection string), takes precedence
	// over the individual host/port/user fields.
	URL      string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// RedisConfig holds Redis configuration.
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// JWTConfig holds JWT configuration.
type JWTConfig struct {
	AccessSecret  string
	RefreshSecret string
	AccessTTL     time.Duration
	RefreshTTL    time.Duration
}

// S3Config holds S3 configuration.
type S3Config struct {
	Endpoint        string
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	UsePathStyle    bool
}

// StripeConfig holds Stripe configuration.
type StripeConfig struct {
	SecretKey     string
	WebhookSecret string
}

// PaystackConfig holds Paystack configuration.
type PaystackConfig struct {
	SecretKey     string
	WebhookSecret string
}

// SMTPConfig holds SMTP configuration.
type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

// Load loads configuration from environment variables.
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			// PORT is what PaaS hosts (Render, Railway) inject; SERVER_PORT is ours.
			Port:         getEnv("PORT", getEnv("SERVER_PORT", "8080")),
			Environment:  getEnv("ENVIRONMENT", "development"),
			ReadTimeout:  getDurationEnv("SERVER_READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getDurationEnv("SERVER_WRITE_TIMEOUT", 10*time.Second),
			FrontendURL:  getEnv("FRONTEND_URL", "http://localhost:5173"),
		},
		Database: DatabaseConfig{
			URL:      getEnv("DATABASE_URL", ""),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "emedic"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getIntEnv("REDIS_DB", 0),
		},
		JWT: JWTConfig{
			AccessSecret:  getEnv("JWT_ACCESS_SECRET", "change-me-access"),
			RefreshSecret: getEnv("JWT_REFRESH_SECRET", "change-me-refresh"),
			AccessTTL:     getDurationEnv("JWT_ACCESS_TTL", 15*time.Minute),
			RefreshTTL:    getDurationEnv("JWT_REFRESH_TTL", 7*24*time.Hour),
		},
		S3: S3Config{
			Endpoint:        getEnv("S3_ENDPOINT", ""),
			Region:          getEnv("S3_REGION", "us-east-1"),
			AccessKeyID:     getEnv("S3_ACCESS_KEY_ID", ""),
			SecretAccessKey: getEnv("S3_SECRET_ACCESS_KEY", ""),
			BucketName:      getEnv("S3_BUCKET_NAME", "emedic-content"),
			UsePathStyle:    getBoolEnv("S3_USE_PATH_STYLE", false),
		},
		Stripe: StripeConfig{
			SecretKey:     getEnv("STRIPE_SECRET_KEY", ""),
			WebhookSecret: getEnv("STRIPE_WEBHOOK_SECRET", ""),
		},
		Paystack: PaystackConfig{
			SecretKey:     getEnv("PAYSTACK_SECRET_KEY", ""),
			WebhookSecret: getEnv("PAYSTACK_WEBHOOK_SECRET", ""),
		},
		Plan: PlanConfig{
			Amount:   getInt64Env("PLAN_MONTHLY_AMOUNT", 500000), // ₦5,000 in kobo
			Currency: getEnv("PLAN_CURRENCY", "NGN"),
		},
		SMTP: SMTPConfig{
			Host:     getEnv("SMTP_HOST", "localhost"),
			Port:     getIntEnv("SMTP_PORT", 587),
			Username: getEnv("SMTP_USERNAME", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
			From:     getEnv("SMTP_FROM", "noreply@emedic.com"),
		},
	}
}

// DatabaseURL returns the PostgreSQL connection string.
func (c *DatabaseConfig) DatabaseURL() string {
	if c.URL != "" {
		return c.URL
	}
	return "postgres://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.DBName + "?sslmode=" + c.SSLMode
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getInt64Env(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getBoolEnv(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
