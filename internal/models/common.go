package models

// PaginationQuery represents pagination parameters
type PaginationQuery struct {
	Page     int    `query:"page"`
	PageSize int    `query:"page_size"`
	SortBy   string `query:"sort_by"`
	Order    string `query:"order"` // "asc" or "desc"
}

// PaginationResponse represents paginated response metadata
type PaginationResponse struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}

// GetDefaultPagination returns default pagination values
func GetDefaultPagination() PaginationQuery {
	return PaginationQuery{
		Page:     1,
		PageSize: 10,
		SortBy:   "id",
		Order:    "desc",
	}
}

// Normalize ensures pagination values are within acceptable ranges
func (p *PaginationQuery) Normalize() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
	if p.Order != "asc" && p.Order != "desc" {
		p.Order = "desc"
	}
	if p.SortBy == "" {
		p.SortBy = "id"
	}
}

// GetOffset calculates the offset for database queries
func (p *PaginationQuery) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}
