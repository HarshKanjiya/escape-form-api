package repositories

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"gorm.io/gorm"
)

type QuestionRepo struct {
	q *query.Query
}

func NewQuestionRepo(db *gorm.DB) *QuestionRepo {
	return &QuestionRepo{
		q: query.Use(db),
	}
}
