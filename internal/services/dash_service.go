package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
)

type DashService struct {
	dashRepo *repositories.DashRepo
}

func NewDashService(dashRepo *repositories.DashRepo) *DashService {
	return &DashService{
		dashRepo: dashRepo,
	}
}

func (s *DashService) GetAnalytics(formId string) (*types.FormAnalytics, error) {
	analytics, err := s.dashRepo.FetchAnalytics(formId)
	if err != nil {
		return nil, err
	}

	return analytics, nil
}

func (s *DashService) GetQuestions(formId string) (interface{}, error) {
	// TODO: Implement GetQuestions logic
	return nil, nil
}

func (s *DashService) GetResponses(formId string) (interface{}, error) {
	// TODO: Implement GetResponses logic
	return nil, nil
}

func (s *DashService) GetPasswords(formId string) (interface{}, error) {
	// TODO: Implement GetPasswords logic
	return nil, nil
}

func (s *DashService) CreatePassword(formId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement CreatePassword logic
	return nil, nil
}

func (s *DashService) UpdatePassword(formId string, passwordId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement UpdatePassword logic
	return nil, nil
}

func (s *DashService) DeletePassword(formId string, passwordId string) error {
	// TODO: Implement DeletePassword logic
	return nil
}

func (s *DashService) UpdateSecurity(formId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement UpdateSecurity logic
	return nil, nil
}

func (s *DashService) UpdateSettings(formId string, body map[string]interface{}) (interface{}, error) {
	// TODO: Implement UpdateSettings logic
	return nil, nil
}
