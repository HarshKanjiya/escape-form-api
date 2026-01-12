package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"gorm.io/gorm"
)

type ProjectService struct {
	q *query.Query
}

func NewProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{
		q: query.Use(db),
	}
}
