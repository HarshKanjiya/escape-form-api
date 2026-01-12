package services

import "github.com/HarshKanjiya/escape-form-api/internal/config"

type FormService struct {
	cfg *config.Config
}

func NewFormService(cfg *config.Config) *FormService {
	return &FormService{
		cfg: cfg,
	}
}
