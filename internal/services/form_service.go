package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"gorm.io/gorm"
)

type FormService struct {
	q *query.Query
}

func NewFormService(db *gorm.DB) *FormService {
	return &FormService{
		q: query.Use(db),
	}
}
