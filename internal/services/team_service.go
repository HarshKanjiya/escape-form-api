package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
)

type TeamService struct {
	teamRepo *repositories.TeamRepo
}

func NewTeamService(teamRepo *repositories.TeamRepo) *TeamService {
	return &TeamService{
		teamRepo: teamRepo,
	}
}

func (ts *TeamService) Get() []types.TeamResponse {

	teams := []types.TeamResponse{}
	// err := ts.q.Team.Where().Scan(&teams).Error
	// if err != nil {
	// 	return []types.TeamResponse{}
	// }

	return teams
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
