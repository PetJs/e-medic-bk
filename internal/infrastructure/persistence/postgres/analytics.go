// Package postgres provides PostgreSQL database implementations.
package postgres

import (
	"github.com/jackc/pgx/v5"

	"emedic-bk/internal/domain/entity"
)

// scanDailyMetrics scans rows of (day timestamp, value bigint) — the shape
// every *ByDay trend query (revenue, signups, new subscriptions) returns.
func scanDailyMetrics(rows pgx.Rows) ([]entity.DailyMetric, error) {
	var metrics []entity.DailyMetric
	for rows.Next() {
		var m entity.DailyMetric
		if err := rows.Scan(&m.Date, &m.Value); err != nil {
			return nil, err
		}
		metrics = append(metrics, m)
	}
	return metrics, rows.Err()
}
