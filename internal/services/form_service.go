package services

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type FormService struct {
	formRepo *repositories.FormRepo
}

func NewFormService(formRepo *repositories.FormRepo) *FormService {
	return &FormService{
		formRepo: formRepo,
	}
}

func (fs *FormService) Get(ctx *fiber.Ctx, pagination *types.PaginationQuery, valid bool, projectId string) ([]*types.FormResponse, int) {

	forms, total, err := fs.formRepo.Get(ctx, pagination, valid, projectId)
	if err != nil {
		return []*types.FormResponse{}, 0
	}

	return forms, total
}

func (fs *FormService) GetById(ctx *fiber.Ctx, formId string) (*types.FormResponse, error) {

	form, err := fs.formRepo.GetById(ctx.Context(), formId)
	if err != nil {
		return nil, utils.HandleDatabaseError(err, "Form")
	}

	return form, nil
}

func (fs *FormService) Create(ctx *fiber.Ctx, formDto *types.CreateFormDto) (*types.FormResponse, error) {
	form, err := fs.formRepo.Create(ctx.Context(), formDto)
	if err != nil {
		return nil, utils.NewAppError("Failed to create form", fiber.StatusInternalServerError, err)
	}
	return form, nil
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

	form, err := fs.formRepo.GetWithTeam(ctx, formId)
	if err != nil {
		return nil, utils.NewAppError("Form not found", fiber.StatusNotFound, err)
	}

	if *form.Team.OwnerID != userId {
		return nil, utils.NewAppError("Forbidden: You do not own this form", fiber.StatusForbidden, nil)
	}

	if err := fs.formRepo.UpdateStatus(ctx, formId, status); err != nil {
		return nil, utils.NewAppError("Failed to update form status", fiber.StatusInternalServerError, err)
	}

	return fs.formRepo.GetById(ctx, formId)
}
