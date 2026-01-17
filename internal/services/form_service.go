package services

import (
	"log"

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

func (fs *FormService) Get(ctx *fiber.Ctx, pagination *types.PaginationQuery, valid bool, projectId string) []*types.FormResponse {

	forms, err := fs.formRepo.Get(ctx, pagination, valid, projectId)
	log.Printf("Fetched forms: %+v", len(forms))
	if err != nil {
		return []*types.FormResponse{}
	}

	return forms
}

func (fs *FormService) Create() types.FormResponse {
	return types.FormResponse{}
}

func (fs *FormService) Update() types.FormResponse {
	return types.FormResponse{}
}

func (fs *FormService) Delete() types.FormResponse {
	return types.FormResponse{}
}
