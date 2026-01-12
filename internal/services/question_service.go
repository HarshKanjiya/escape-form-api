package services

import "github.com/HarshKanjiya/escape-form-api/internal/config"

type QuestionService struct {
	cfg *config.Config
}

func NewQuestionService(cfg *config.Config) *QuestionService {
	return &QuestionService{
		cfg: cfg,
	}
}
