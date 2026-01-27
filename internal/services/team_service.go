package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
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

func (ts *TeamService) Get(ctx *fiber.Ctx, pagination *types.PaginationQuery, valid bool) ([]*types.TeamResponse, int) {

	teams, total, err := ts.teamRepo.Get(ctx, pagination, valid)
	if err != nil {
		return []*types.TeamResponse{}, 0
	}

	return teams, total
}

func (ts *TeamService) Create(ctx *fiber.Ctx, team *types.TeamDto) (types.TeamResponse, error) {
	createdTeam, err := ts.teamRepo.Create(ctx, team)
	if err != nil {
		return types.TeamResponse{}, err
	}

	ownerId := ""
	if createdTeam.OwnerID != nil {
		ownerId = *createdTeam.OwnerID
	}
	planId := ""
	if createdTeam.PlanID != nil {
		planId = *createdTeam.PlanID
	}
	name := ""
	if createdTeam.Name != nil {
		name = *createdTeam.Name
	}

	return types.TeamResponse{
		ID:        createdTeam.ID,
		Name:      name,
		OwnerId:   ownerId,
		PlanId:    planId,
		Valid:     createdTeam.Valid,
		CreatedAt: utils.GetIsoDateTime(&createdTeam.CreatedAt),
		UpdatedAt: utils.GetIsoDateTime(createdTeam.UpdatedAt),
	}, nil
}

func (ts *TeamService) Update(ctx *fiber.Ctx, team *types.TeamDto) (bool, error) {
	ok, err := ts.teamRepo.Update(ctx, &models.Team{
		ID:   team.ID,
		Name: &team.Name,
	})
	if err != nil || !ok {
		return false, err
	}
	return true, nil
}

func (ts *TeamService) Delete(ctx *fiber.Ctx, teamId string) (bool, error) {
	ok, err := ts.teamRepo.Delete(ctx, teamId)
	if err != nil || !ok {
		return false, err
	}
	return true, nil
}
