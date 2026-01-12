package services

import "github.com/HarshKanjiya/escape-form-api/internal/config"

type EdgeService struct {
	cfg *config.Config
	db  *config.DatabaseConfig
}

func NewEdgeService(cfg *config.Config, db *config.DatabaseConfig) *EdgeService {
	return &EdgeService{
		cfg: cfg,
		db:  db,
	}
}
