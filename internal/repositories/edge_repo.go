package repositories

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"gorm.io/gorm"
)

type EdgeRepo struct {
	q *query.Query
}

func NewEdgeRepo(db *gorm.DB) *EdgeRepo {
	return &EdgeRepo{
		q: query.Use(db),
	}
}
