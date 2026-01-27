package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ProjectService struct {
	projectRepo *repositories.ProjectRepo
}

func NewProjectService(projectRepo *repositories.ProjectRepo) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
	}
}

func (ps *ProjectService) Get(ctx *fiber.Ctx, pagination *types.PaginationQuery, valid bool, teamId string) ([]*types.ProjectResponse, int, error) {

	projects, total, err := ps.projectRepo.Get(ctx, pagination, valid, teamId)
	if err != nil {
		return []*types.ProjectResponse{}, 0, err
	}

	return projects, total, nil
}

func (ps *ProjectService) GetById(ctx *fiber.Ctx, projectId string) (*types.ProjectResponse, error) {

	project, err := ps.projectRepo.GetById(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (ps *ProjectService) Create(ctx *fiber.Ctx, project *types.ProjectDto) (types.ProjectResponse, error) {
	createdProject, err := ps.projectRepo.Create(ctx, project)
	if err != nil {
		return types.ProjectResponse{}, err
	}

	description := ""
	if createdProject.Description != nil {
		description = *createdProject.Description
	}

	return types.ProjectResponse{
		ID:          createdProject.ID,
		Name:        createdProject.Name,
		Description: description,
		TeamID:      createdProject.TeamID,
		Valid:       createdProject.Valid,
		CreatedAt:   utils.GetIsoDateTime(createdProject.CreatedAt),
		UpdatedAt:   utils.GetIsoDateTime(createdProject.UpdatedAt),
	}, nil
}

func (ps *ProjectService) Update(ctx *fiber.Ctx, project *types.ProjectDto) (bool, error) {
	ok, err := ps.projectRepo.Update(ctx, &models.Project{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
	})
	if err != nil || !ok {
		return false, err
	}
	return true, nil
}

func (ps *ProjectService) Delete(ctx *fiber.Ctx, projectId string) (bool, error) {
	ok, err := ps.projectRepo.Delete(ctx, projectId)
	if err != nil || !ok {
		return false, err
	}
	return true, nil
}
