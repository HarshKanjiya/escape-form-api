package services

import "github.com/HarshKanjiya/escape-form-api/internal/config"

type FormService struct {
	cfg *config.Config
	db  *config.DatabaseConfig
}

func NewFormService(cfg *config.Config, db *config.DatabaseConfig) *FormService {
	return &FormService{
		cfg: cfg,
		db:  db,
	}
}
