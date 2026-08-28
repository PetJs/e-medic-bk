// Package dto contains Data Transfer Objects for use cases.
package dto

// SearchResultItem is one lightweight search hit — just enough to render a
// result row and navigate to it; not the full resource response.
type SearchResultItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// SearchResponse groups search results by entity type.
type SearchResponse struct {
	Courses []*SearchResultItem `json:"courses"`
	Modules []*SearchResultItem `json:"modules"`
	Lessons []*SearchResultItem `json:"lessons"`
}
