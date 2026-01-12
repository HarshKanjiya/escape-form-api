package services

import "github.com/HarshKanjiya/escape-form-api/internal/config"

type ProjectService struct {
	cfg *config.Config
	db  *config.DatabaseConfig
}

func NewProjectService(cfg *config.Config, db *config.DatabaseConfig) *ProjectService {
	return &ProjectService{
		cfg: cfg,
		db:  db,
	}
}
