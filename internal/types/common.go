package types

type Pagination struct {
	Page     int    `json:"page" validate:"gte=1"`
	Limit    int    `json:"limit" validate:"gte=1,lte=100"`
	OrderBy  string `json:"orderBy"`
	OrderDir string `json:"orderDir" validate:"oneof=asc desc"`
	Search   string `json:"search"`
}
