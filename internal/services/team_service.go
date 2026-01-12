package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"gorm.io/gorm"
)

type TeamService struct {
	q *query.Query
}

func NewTeamService(db *gorm.DB) *TeamService {
	return &TeamService{
		q: query.Use(db),
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
