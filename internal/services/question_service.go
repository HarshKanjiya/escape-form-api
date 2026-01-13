package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
)

type QuestionService struct {
	questionRepo *repositories.QuestionRepo
}

func NewQuestionService(questionRepo *repositories.QuestionRepo) *QuestionService {
	return &QuestionService{
		questionRepo: questionRepo,
	}
}
