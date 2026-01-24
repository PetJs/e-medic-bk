// Package pagination provides pagination utilities.
package pagination

// Pagination holds pagination parameters.
type Pagination struct {
	Page  int
	Limit int
}

// DefaultPagination returns default pagination settings.
func DefaultPagination() Pagination {
	return Pagination{
		Page:  1,
		Limit: 20,
	}
}

// Offset calculates the database offset.
func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

// Normalize ensures pagination values are within valid ranges.
func (p *Pagination) Normalize() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = 20
	}
	if p.Limit > 100 {
		p.Limit = 100
	}
}
