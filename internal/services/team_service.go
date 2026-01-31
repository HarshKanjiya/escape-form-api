package services

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
)

type ITeamService interface {
	Get(ctx context.Context, userId string, pagination *types.PaginationQuery, valid bool) ([]*types.TeamResponse, int64, error)
	Create(ctx context.Context, userId string, team *types.TeamRequest) error
	Update(ctx context.Context, userId string, team *types.TeamRequest) error
	Delete(ctx context.Context, userId string, teamId string) error
}

type TeamService struct {
	teamRepo repositories.ITeamRepo
}

func NewTeamService(teamRepo repositories.ITeamRepo) *TeamService {
	return &TeamService{
		teamRepo: teamRepo,
	}
}

func (ts *TeamService) Get(ctx context.Context, userId string, pagination *types.PaginationQuery, valid bool) ([]*types.TeamResponse, int64, error) {

	teams, totalCount, err := ts.teamRepo.Get(ctx, userId, pagination, valid)
	if err != nil {
		return nil, 0, err
	}
	return teams, totalCount, nil
}

func (ts *TeamService) Create(ctx context.Context, userId string, team *types.TeamRequest) error {

	newTeam := &models.Team{
		ID:        utils.GenerateUUID(),
		Name:      &team.Name,
		OwnerID:   &userId,
		PlanID:    nil,
		Valid:     true,
		CreatedAt: *utils.GetCurrentTime(),
	}

	err := ts.teamRepo.Create(ctx, newTeam)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TeamService) Update(ctx context.Context, userId string, team *types.TeamRequest) error {

	teamDb, err := ts.teamRepo.GetById(ctx, team.ID)
	if err != nil {
		return err
	}

	if teamDb == nil {
		return errors.NotFound("Team")
	}

	if teamDb.OwnerID == nil || *teamDb.OwnerID != userId {
		return errors.Unauthorized("")
	}

	err = ts.teamRepo.Update(ctx, &models.Team{
		ID:        team.ID,
		Name:      &team.Name,
		UpdatedAt: utils.GetCurrentTime(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (ts *TeamService) Delete(ctx context.Context, userId string, teamId string) error {

	teamDb, err := ts.teamRepo.GetById(ctx, teamId)
	if err != nil {
		return err
	}
	if teamDb == nil {
		return errors.NotFound("Team")
	}
	if teamDb.OwnerID == nil || *teamDb.OwnerID != userId {
		return errors.Unauthorized("")
	}

	err = ts.teamRepo.Delete(ctx, teamId)
	if err != nil {
		return err
	}
	return nil
}
