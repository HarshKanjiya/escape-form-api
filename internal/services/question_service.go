package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"gorm.io/gorm"
)

type QuestionService struct {
	q *query.Query
}

func NewQuestionService(db *gorm.DB) *QuestionService {
	return &QuestionService{
		q: query.Use(db),
	}
}
