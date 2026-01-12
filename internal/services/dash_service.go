package services

import "github.com/HarshKanjiya/escape-form-api/internal/config"

type DashService struct {
	cfg *config.Config
	db  *config.DatabaseConfig
}

func NewDashService(cfg *config.Config, db *config.DatabaseConfig) *DashService {
	return &DashService{
		cfg: cfg,
		db:  db,
	}
}
