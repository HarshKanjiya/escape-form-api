package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
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

	form, err := fs.formRepo.GetById(ctx, formId)
	if err != nil {
		return nil, err
	}

	return form, nil
}

func (fs *FormService) Create(ctx *fiber.Ctx, formDto *types.CreateFormDto) (*types.FormResponse, error) {
	return fs.formRepo.Create(ctx, formDto)
}

func (fs *FormService) Update(ctx *fiber.Ctx, formDto *types.CreateFormDto) (*types.FormResponse, error) {
	return nil, nil

}

func (fs *FormService) Delete(ctx *fiber.Ctx, formDto *types.CreateFormDto) (*types.FormResponse, error) {
	return nil, nil

}

func (fs *FormService) UpdateStatus(ctx *fiber.Ctx, formId string, status *models.FormStatus) (*types.FormResponse, error) {
	return fs.formRepo.UpdateStatus(ctx, formId, status)
}
