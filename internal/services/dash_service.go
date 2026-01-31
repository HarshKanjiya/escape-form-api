package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

type IDashService interface {
}

type DashService struct {
	dashRepo repositories.IDashRepo
}

func NewDashService(dashRepo repositories.IDashRepo) *DashService {
	return &DashService{
		dashRepo: dashRepo,
	}
}

func (s *DashService) GetAnalytics(ctx *fiber.Ctx, formId string) (*types.FormAnalytics, error) {
	analytics, err := s.dashRepo.FetchAnalytics(ctx, formId)
	if err != nil {
		return nil, err
	}

	return analytics, nil
}

func (s *DashService) GetQuestions(ctx *fiber.Ctx, formId string) (interface{}, error) {
	// TODO: Implement GetQuestions logic
	return nil, nil
}

func (s *DashService) GetResponses(ctx *fiber.Ctx, formId string) (interface{}, error) {
	// TODO: Implement GetResponses logic
	return nil, nil
}

func (s *DashService) GetPasswords(ctx *fiber.Ctx, formId string) (interface{}, error) {
	// TODO: Implement GetPasswords logic
	return nil, nil
}

func (s *DashService) CreatePassword(ctx *fiber.Ctx, formId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement CreatePassword logic
	return nil, nil
}

func (s *DashService) UpdatePassword(ctx *fiber.Ctx, formId string, passwordId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement UpdatePassword logic
	return nil, nil
}

func (s *DashService) DeletePassword(ctx *fiber.Ctx, formId string, passwordId string) error {
	// TODO: Implement DeletePassword logic
	return nil
}

func (s *DashService) UpdateSecurity(ctx *fiber.Ctx, formId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement UpdateSecurity logic
	return nil, nil
}

func (s *DashService) UpdateSettings(ctx *fiber.Ctx, formId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement UpdateSettings logic
	return nil, nil
}
