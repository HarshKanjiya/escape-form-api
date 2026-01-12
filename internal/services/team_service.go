package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
)

type TeamService struct {
	cfg *config.Config
	db  *config.DatabaseConfig
}

func NewTeamService(cfg *config.Config, db *config.DatabaseConfig) *TeamService {
	return &TeamService{
		cfg: cfg,
		db:  db,
	}
}

func (ts *TeamService) Get() []types.TeamResponse {
	return []types.TeamResponse{}
}

func (ts *TeamService) Create() types.TeamResponse {
	return types.TeamResponse{}
}

func (ts *TeamService) Update() types.TeamResponse {
	return types.TeamResponse{}
}

func (ts *TeamService) Delete() types.TeamResponse {
	return types.TeamResponse{}
}
