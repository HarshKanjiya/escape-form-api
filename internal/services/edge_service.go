package services

import "github.com/HarshKanjiya/escape-form-api/internal/config"

type EdgeService struct {
	cfg *config.Config
}

func NewEdgeService(cfg *config.Config) *EdgeService {
	return &EdgeService{
		cfg: cfg,
	}
}
