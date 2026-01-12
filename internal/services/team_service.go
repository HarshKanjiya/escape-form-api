package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/config"
)

type TeamService struct {
	cfg *config.Config
}

func NewTeamService(cfg *config.Config) *TeamService {
	return &TeamService{
		cfg: cfg,
	}
}
