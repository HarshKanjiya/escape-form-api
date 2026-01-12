package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"gorm.io/gorm"
)

type EdgeService struct {
	q *query.Query
}

func NewEdgeService(db *gorm.DB) *EdgeService {
	return &EdgeService{
		q: query.Use(db),
	}
}
