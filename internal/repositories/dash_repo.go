package repositories

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"gorm.io/gorm"
)

type DashRepo struct {
	q *query.Query
}

func NewDashRepo(db *gorm.DB) *DashRepo {
	return &DashRepo{
		q: query.Use(db),
	}
}
