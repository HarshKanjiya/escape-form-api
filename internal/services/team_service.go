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

	var teamResponses []*types.TeamResponse
	for _, team := range teams {
		teamResponses = append(teamResponses, &types.TeamResponse{
			ID:        team.ID,
			Name:      *team.Name,
			OwnerId:   *team.OwnerID,
			PlanId:    *team.PlanID,
			Valid:     team.Valid,
			CreatedAt: team.CreatedAt.String(),
			UpdatedAt: team.UpdatedAt.String(),
		})
	}

	return teamResponses
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
