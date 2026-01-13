package repositories

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	q *query.Query
}

func NewProjectRepo(db *gorm.DB) *ProjectRepo {
	return &ProjectRepo{
		q: query.Use(db),
	}
}
