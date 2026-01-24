// Package logger provides structured logging.
package logger

import (
	"log/slog"
	"os"
)

var defaultLogger *slog.Logger

func init() {
	defaultLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

// SetLogger sets the default logger.
func SetLogger(l *slog.Logger) {
	defaultLogger = l
}

// Info logs an info message.
func Info(msg string, args ...any) {
	defaultLogger.Info(msg, args...)
}

// Debug logs a debug message.
func Debug(msg string, args ...any) {
	defaultLogger.Debug(msg, args...)
}

// Warn logs a warning message.
func Warn(msg string, args ...any) {
	defaultLogger.Warn(msg, args...)
}

// Error logs an error message.
func Error(msg string, args ...any) {
	defaultLogger.Error(msg, args...)
}

// With returns a logger with additional context.
func With(args ...any) *slog.Logger {
	return defaultLogger.With(args...)
}
