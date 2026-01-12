package services

import "github.com/HarshKanjiya/escape-form-api/internal/config"

type ProjectService struct {
	cfg *config.Config
}

func NewProjectService(cfg *config.Config) *ProjectService {
	return &ProjectService{
		cfg: cfg,
	}
}
