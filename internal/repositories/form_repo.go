package repositories

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"gorm.io/gorm"
)

type FormRepo struct {
	q *query.Query
}

func NewFormRepo(db *gorm.DB) *FormRepo {
	return &FormRepo{
		q: query.Use(db),
	}
}
