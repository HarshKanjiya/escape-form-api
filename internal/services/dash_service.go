package services

import "github.com/HarshKanjiya/escape-form-api/internal/config"

type DashService struct {
	cfg *config.Config
}

func NewDashService(cfg *config.Config) *DashService {
	return &DashService{
		cfg: cfg,
	}
}
