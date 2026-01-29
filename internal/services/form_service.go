package services

import (
	"context"
	"log"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FormService struct {
	formRepo    *repositories.FormRepo
	projectRepo *repositories.ProjectRepo
}

func NewFormService(formRepo *repositories.FormRepo, projectRepo *repositories.ProjectRepo) *FormService {
	return &FormService{
		formRepo:    formRepo,
		projectRepo: projectRepo,
	}
}

func (fs *FormService) Get(ctx context.Context, userId string, pagination *types.PaginationQuery, valid bool, projectId string) ([]*types.FormResponse, int, error) {

	// Check if user owns the project
	_, err := fs.projectRepo.GetWithTeam(ctx, userId, projectId)
	if err != nil {
		return nil, 0, utils.HandleDatabaseError(err, "Project")
	}

	forms, total, err := fs.formRepo.Get(ctx, pagination, valid, projectId)
	if err != nil {
		return nil, 0, utils.NewAppError("Failed to fetch forms", fiber.StatusInternalServerError, err)
	}

	return forms, total, nil
}

func (fs *FormService) GetById(ctx context.Context, userId string, formId string) (*types.FormResponse, error) {

	form, err := fs.formRepo.GetById(ctx, formId)
	if err != nil {
		return nil, utils.HandleDatabaseError(err, "Form")
	}

	return form, nil
}

func (fs *FormService) Create(ctx context.Context, userId string, formDto *types.CreateFormDto) (*types.FormResponse, error) {
	// Check if user owns the project
	project, err := fs.projectRepo.GetWithTeam(ctx, userId, formDto.ProjectID)
	if err != nil {
		return nil, utils.HandleDatabaseError(err, "Project")
	}
	// Build the form model
	status := models.FormStatusDraft
	form := &models.Form{
		ID:           uuid.New().String(),
		Name:         formDto.Name,
		Description:  formDto.Description,
		ProjectID:    formDto.ProjectID,
		TeamID:       project.TeamID,
		Valid:        true,
		CreatedBy:    userId,
		FormPageType: models.FormPageTypeSingle,
		Status:       &status,
	}

	formResponse, err := fs.formRepo.Create(ctx, form)
	if err != nil {
		return nil, utils.NewAppError("Failed to create form", fiber.StatusInternalServerError, err)
	}
	return formResponse, nil
}

func (fs *FormService) Update(ctx *fiber.Ctx, formDto *types.CreateFormDto) (*types.FormResponse, error) {
	return nil, nil

}

func (fs *FormService) Delete(ctx *fiber.Ctx, formDto *types.CreateFormDto) (*types.FormResponse, error) {
	return nil, nil

}

func (fs *FormService) UpdateStatus(
	ctx context.Context,
	userId string,
	formId string,
	status models.FormStatus,
) (*types.FormResponse, error) {

	_, err := fs.formRepo.GetWithTeam(ctx, userId, formId)
	if err != nil {
		log.Printf("Error fetching form: %v", err)
		return nil, utils.NewAppError("Form not found", fiber.StatusNotFound, err)
	}

	if err := fs.formRepo.UpdateStatus(ctx, formId, status); err != nil {
		log.Printf("Error updating form status: %v", err)
		return nil, utils.NewAppError("Failed to update form status", fiber.StatusInternalServerError, err)
	}
	return fs.formRepo.GetById(ctx, formId)
}
