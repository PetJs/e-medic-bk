// Package entity contains the core domain entities.
package entity

import "time"

// DailyMetric is one bucket of a day-by-day trend (revenue, signups, new
// subscriptions — whatever the caller aggregated by day).
type DailyMetric struct {
	Date  time.Time
	Value int64
}

// ModuleCompletionRate is the share of enrolled students who have completed
// every lesson in a module.
type ModuleCompletionRate struct {
	ModuleID      string
	ModuleTitle   string
	CompletionPct float64
}
