package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

type TeamService struct {
	teamRepo *repositories.TeamRepo
}

func NewTeamService(teamRepo *repositories.TeamRepo) *TeamService {
	return &TeamService{
		teamRepo: teamRepo,
	}
}

func (ts *TeamService) Get(ctx *fiber.Ctx, pagination *types.PaginationQuery, valid bool) []*types.TeamResponse {

	teams, err := ts.teamRepo.Get(ctx, pagination, valid)
	if err != nil {
		return []*types.TeamResponse{}
	}

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
