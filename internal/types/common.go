package types

type PaginationQuery struct {
	Page   int    `query:"page" validate:"gte=1"`
	Limit  int    `query:"limit" validate:"gte=1,lte=100"`
	SortBy string `query:"sort_by"`
	Order  string `query:"order" validate:"oneof=asc desc"`
	Search string `query:"search"`
}

// GetDefaultPagination returns default pagination values
func GetDefaultPagination() PaginationQuery {
	return PaginationQuery{
		Page:   1,
		Limit:  10,
		SortBy: "id",
		Order:  "desc",
	}
}

// Normalize ensures pagination values are within acceptable ranges
func (p *PaginationQuery) Normalize() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = 10
	}
	if p.Limit > 100 {
		p.Limit = 100
	}
	if p.Order != "asc" && p.Order != "desc" {
		p.Order = "desc"
	}
	if p.SortBy == "" {
		p.SortBy = "id"
	}
}

func (p *PaginationQuery) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

type ResponseObj struct {
	Type       string      `json:"type" validate:"oneof=success warning error info"`
	Message    string      `json:"message"`
	TotalCount int64       `json:"totalItems"`
	Data       interface{} `json:"data"`
}
