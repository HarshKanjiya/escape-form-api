package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
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

func (ps *ProjectService) Create() types.ProjectResponse {
	return types.ProjectResponse{}
}

func (ps *ProjectService) Update() types.ProjectResponse {
	return types.ProjectResponse{}
}

func (ps *ProjectService) Delete() types.ProjectResponse {
	return types.ProjectResponse{}
}
