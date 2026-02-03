package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/mapper"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
)

type IDashService interface {
	GetAnalytics(ctx context.Context, userId string, formId string) (*types.FormAnalytics, error)
	GetResponses(ctx context.Context, userId string, formId string) ([]*models.Response, error)
	GetQuestions(ctx context.Context, formId string) ([]*types.QuestionResponse, error)

	GetPasswords(ctx context.Context, userId string, formId string) ([]*types.ActivePasswordResponse, error)
	CreatePassword(ctx context.Context, userId string, formId string, password types.PasswordRequest) (*types.ActivePasswordResponse, error)
	UpdatePassword(ctx context.Context, userId string, formId string, passwordId string, body types.PasswordRequest) error
	DeletePassword(ctx context.Context, userId string, formId string, passwordId string) error

	UpdateSecurity(ctx context.Context, userId string, formId string, body *types.UpdateSecurityRequest) error
	UpdateSettings(ctx context.Context, userId string, formId string, body *types.UpdateSettingsRequest) error
}

type DashService struct {
	dashRepo repositories.IDashRepo
	formRepo repositories.IFormRepo
}

func NewDashService(
	dashRepo repositories.IDashRepo,
	formRepo repositories.IFormRepo,
) *DashService {
	return &DashService{
		dashRepo: dashRepo,
		formRepo: formRepo,
	}
}

func (s *DashService) GetAnalytics(ctx context.Context, userId string, formId string) (*types.FormAnalytics, error) {

	form, err := s.formRepo.GetWithTeam(ctx, formId)
	if err != nil {
		return nil, err
	}

	if form == nil {
		return nil, errors.NotFound("Form")
	}

	if *form.Team.OwnerID != userId {
		return nil, errors.Unauthorized("Overview Analytics")
	}

	analytics, err := s.dashRepo.GetAnalytics(ctx, formId)
	if err != nil {
		return nil, err
	}

	return analytics, nil
}

func (s *DashService) GetResponses(ctx context.Context, userId string, formId string) ([]*models.Response, error) {

	form, err := s.formRepo.GetWithTeam(ctx, formId)
	if err != nil {
		return nil, err
	}

	if form == nil {
		return nil, errors.NotFound("Form")
	}

	if *form.Team.OwnerID != userId {
		return nil, errors.Unauthorized("Overview Analytics")
	}

	responses, err := s.dashRepo.GetResponses(ctx, formId)
	if err != nil {
		return nil, err
	}
	if len(responses) == 0 {
		return nil, errors.NotFound("Responses")
	}

	return responses, nil
}

func (s *DashService) GetQuestions(ctx context.Context, formId string) ([]*types.QuestionResponse, error) {

	questions, err := s.dashRepo.GetQuestions(ctx, formId)
	if err != nil {
		return nil, err
	}
	if len(questions) == 0 {
		return nil, errors.NotFound("Questions")
	}

	questionResponses := make([]*types.QuestionResponse, len(questions))
	for i, question := range questions {
		questionResponses[i] = mapper.ToQuestionResponse(question)
	}

	return questionResponses, nil
}

func (s *DashService) GetPasswords(ctx context.Context, userId string, formId string) ([]*types.ActivePasswordResponse, error) {

	form, err := s.formRepo.GetWithTeam(ctx, formId)
	if err != nil {
		return nil, err
	}

	if form == nil {
		return nil, errors.NotFound("Form")
	}

	if *form.Team.OwnerID != userId {
		return nil, errors.Unauthorized("Overview Analytics")
	}

	passwords, err := s.dashRepo.GetPasswords(ctx, formId)
	if err != nil {
		return nil, err
	}

	passResponse := make([]*types.ActivePasswordResponse, len(passwords))
	for i, pass := range passwords {
		passResponse[i] = mapper.ToActivePasswordResponse(pass)
	}

	return passResponse, nil
}

func (s *DashService) CreatePassword(ctx context.Context, userId string, formId string, password types.PasswordRequest) (*types.ActivePasswordResponse, error) {

	form, err := s.formRepo.GetWithTeam(ctx, formId)
	if err != nil {
		return nil, err
	}

	if form == nil {
		return nil, errors.NotFound("Form")
	}

	if *form.Team.OwnerID != userId {
		return nil, errors.Unauthorized("Overview Analytics")
	}

	var expireAt *time.Time
	if password.ExpireAt != "" {
		parsedTime, err := time.Parse("2006-01-02", password.ExpireAt)
		if err != nil {
			return nil, err
		}
		expireAt = &parsedTime
	}

	newPassword := &models.ActivePassword{
		ID:         utils.GenerateUUID(),
		FormID:     formId,
		Password:   password.Password,
		Name:       password.Name,
		UsableUpto: password.UsableUpto,
		IsValid:    password.IsValid,
		ExpireAt:   expireAt,
	}

	createdPassword, err := s.dashRepo.CreatePassword(ctx, formId, newPassword)
	if err != nil {
		return nil, err
	}

	return mapper.ToActivePasswordResponse(createdPassword), nil
}

func (s *DashService) UpdatePassword(ctx context.Context, userId string, formId string, passwordId string, password types.PasswordRequest) error {

	form, err := s.formRepo.GetWithTeam(ctx, formId)
	if err != nil {
		return err
	}

	if form == nil {
		return errors.NotFound("Form")
	}

	if *form.Team.OwnerID != userId {
		return errors.Unauthorized("Overview Analytics")
	}

	var expireAt *time.Time
	if password.ExpireAt != "" {
		parsedTime, err := time.Parse("2006-01-02", password.ExpireAt)
		if err != nil {
			return err
		}
		expireAt = &parsedTime
	}

	pass := &models.ActivePassword{
		ID:         passwordId,
		FormID:     formId,
		Password:   password.Password,
		Name:       password.Name,
		UsableUpto: password.UsableUpto,
		IsValid:    password.IsValid,
		ExpireAt:   expireAt,
	}

	err = s.dashRepo.UpdatePassword(ctx, formId, pass)
	if err != nil {
		return err
	}

	return nil
}

func (s *DashService) DeletePassword(ctx context.Context, userId string, formId string, passwordId string) error {

	form, err := s.formRepo.GetWithTeam(ctx, formId)
	if err != nil {
		return err
	}

	if form == nil {
		return errors.NotFound("Form")
	}

	if *form.Team.OwnerID != userId {
		return errors.Unauthorized("Overview Analytics")
	}

	err = s.dashRepo.DeletePassword(ctx, passwordId)
	if err != nil {
		return err
	}
	return nil
}

func (s *DashService) UpdateSecurity(ctx context.Context, userId string, formId string, body *types.UpdateSecurityRequest) error {

	form, err := s.formRepo.GetById(ctx, formId)
	if err != nil {
		return err
	}

	if form == nil {
		return errors.NotFound("Form")
	}

	if form.CreatedBy != userId {
		return errors.Unauthorized("")
	}

	updates := make(map[string]interface{})
	data, _ := json.Marshal(body)
	json.Unmarshal(data, &updates)

	err = s.formRepo.Update(ctx, formId, updates)
	if err != nil {
		return err
	}

	return nil
}

func (s *DashService) UpdateSettings(ctx context.Context, userId string, formId string, body *types.UpdateSettingsRequest) error {

	form, err := s.formRepo.GetById(ctx, formId)
	if err != nil {
		return err
	}

	if form == nil {
		return errors.NotFound("Form")
	}

	if form.CreatedBy != userId {
		return errors.Unauthorized("")
	}

	updates := make(map[string]interface{})
	data, _ := json.Marshal(body)
	json.Unmarshal(data, &updates)

	err = s.formRepo.Update(ctx, formId, updates)
	if err != nil {
		return err
	}

	return nil
}
