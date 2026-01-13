package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
)

type ProjectService struct {
	projectRepo *repositories.ProjectRepo
}

func NewProjectService(projectRepo *repositories.ProjectRepo) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
	}
}
