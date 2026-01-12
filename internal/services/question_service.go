package services

import "github.com/HarshKanjiya/escape-form-api/internal/config"

type QuestionService struct {
	cfg *config.Config
	db  *config.DatabaseConfig
}

func NewQuestionService(cfg *config.Config, db *config.DatabaseConfig) *QuestionService {
	return &QuestionService{
		cfg: cfg,
		db:  db,
	}
}
