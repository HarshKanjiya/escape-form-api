package services

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
)

type IProjectService interface {
	Get(ctx context.Context, userId string, pagination *types.PaginationQuery, teamId string) ([]*types.ProjectResponse, int64, error)
	GetById(ctx context.Context, userId string, projectId string) (*types.ProjectResponse, error)
	Create(ctx context.Context, userId string, project *types.ProjectDto) (types.ProjectResponse, error)
	Update(ctx context.Context, userId string, project *types.ProjectDto) (bool, error)
	Delete(ctx context.Context, userId string, projectId string) (bool, error)
}

type ProjectService struct {
	projectRepo repositories.IProjectRepo
	teamRepo    repositories.ITeamRepo
}

func NewProjectService(
	projectRepo repositories.IProjectRepo,
	teamRepo repositories.ITeamRepo,
) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
		teamRepo:    teamRepo,
	}
}

func (s *ProjectService) Get(ctx context.Context, userId string, pagination *types.PaginationQuery, teamId string) ([]*types.ProjectResponse, int64, error) {

	team, err := s.teamRepo.GetById(ctx, teamId)
	if err != nil {
		return []*types.ProjectResponse{}, 0, errors.NotFound("Team")
	}
	if *team.OwnerID != userId {
		return []*types.ProjectResponse{}, 0, errors.Unauthorized("team's projects")
	}

	projects, total, err := s.projectRepo.Get(ctx, pagination, teamId)
	if err != nil {
		return []*types.ProjectResponse{}, 0, err
	}

	return projects, total, nil
}

func (s *ProjectService) GetById(ctx context.Context, userId string, projectId string) (*types.ProjectResponse, error) {

	project, err := s.projectRepo.GetWithTeam(ctx, projectId)
	if err != nil {
		return nil, errors.NotFound("Project")
	}

	if *project.Team.OwnerID != userId {
		return nil, errors.Unauthorized("project")
	}

	return &types.ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: "",
		TeamID:      project.TeamID,
		Valid:       project.Valid,
		CreatedAt:   utils.GetIsoDateTime(project.CreatedAt),
		UpdatedAt:   utils.GetIsoDateTime(project.UpdatedAt),
	}, nil
}

func (s *ProjectService) Create(ctx context.Context, userId string, project *types.ProjectDto) (types.ProjectResponse, error) {
	team, err := s.teamRepo.GetById(ctx, project.TeamID)
	if err != nil {
		return types.ProjectResponse{}, errors.NotFound("Team")
	}

	if *team.OwnerID != userId {
		return types.ProjectResponse{}, errors.Unauthorized("team")
	}

	createdProject, err := s.projectRepo.Create(ctx, &models.Project{
		ID:          utils.GenerateUUID(),
		Name:        project.Name,
		Description: project.Description,
		TeamID:      project.TeamID,
		Valid:       true,
	})
	if err != nil {
		return types.ProjectResponse{}, err
	}

	return types.ProjectResponse{
		ID:          createdProject.ID,
		Name:        createdProject.Name,
		Description: *createdProject.Description,
		TeamID:      createdProject.TeamID,
		Valid:       createdProject.Valid,
		CreatedAt:   utils.GetIsoDateTime(createdProject.CreatedAt),
		UpdatedAt:   utils.GetIsoDateTime(createdProject.UpdatedAt),
	}, nil
}

func (s *ProjectService) Update(ctx context.Context, userId string, project *types.ProjectDto) (bool, error) {

	existingProject, err := s.projectRepo.GetWithTeam(ctx, project.ID)
	if err != nil {
		return false, errors.NotFound("Project")
	}
	if *existingProject.Team.OwnerID != userId {
		return false, errors.Unauthorized("project")
	}

	ok, err := s.projectRepo.Update(ctx, &models.Project{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
	})
	if err != nil || !ok {
		return false, err
	}
	return true, nil
}

func (s *ProjectService) Delete(ctx context.Context, userId string, projectId string) (bool, error) {
	existingProject, err := s.projectRepo.GetWithTeam(ctx, projectId)
	if err != nil {
		return false, errors.NotFound("Project")
	}
	if *existingProject.Team.OwnerID != userId {
		return false, errors.Unauthorized("project")
	}
	err = s.projectRepo.Delete(ctx, projectId)
	if err != nil {
		return false, err
	}
	return true, nil
}
