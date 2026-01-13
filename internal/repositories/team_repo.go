package repositories

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"gorm.io/gorm"
)

type TeamRepo struct {
	q *query.Query
}

func NewTeamRepo(db *gorm.DB) *TeamRepo {
	return &TeamRepo{
		q: query.Use(db),
	}
}
