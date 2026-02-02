package services

import (
	"context"
	"log"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
)

type IFormService interface {
	Get(ctx context.Context, userId string, pagination *types.PaginationQuery, projectId string, status string) ([]*types.FormResponse, int64, error)
	GetById(ctx context.Context, userId string, formId string) (*types.FormResponse, error)
	Create(ctx context.Context, userId string, projectId string, form *types.CreateFormRequest) (*types.FormResponse, error)
	UpdateStatus(ctx context.Context, userId string, formId string, status models.FormStatus) error
	Delete(ctx context.Context, userId string, formId string) error

	UpdateSequence(ctx context.Context, userId string, formId string, sequences []*types.SequenceItem) error
}

type FormService struct {
	formRepo    repositories.IFormRepo
	projectRepo repositories.IProjectRepo
}

func NewFormService(formRepo repositories.IFormRepo, projectRepo repositories.IProjectRepo) *FormService {
	return &FormService{
		formRepo:    formRepo,
		projectRepo: projectRepo,
	}
}

func (s *FormService) Get(ctx context.Context, userId string, pagination *types.PaginationQuery, projectId string, status string) ([]*types.FormResponse, int64, error) {

	project, err := s.projectRepo.GetWithTeam(ctx, projectId)
	if err != nil {
		return nil, 0, err
	}
	if project == nil {
		return nil, 0, errors.NotFound("Project")
	}
	if project.Team.OwnerID == nil || *project.Team.OwnerID != userId {
		return nil, 0, errors.Unauthorized("")
	}

	var statusPtr *models.FormStatus
	if status != "" {
		statusVal := models.FormStatus(status)
		statusPtr = &statusVal
	}

	forms, total, err := s.formRepo.Get(ctx, pagination, projectId, statusPtr)
	if err != nil {
		return nil, 0, err
	}

	formResponses := make([]*types.FormResponse, len(forms))
	for i, form := range forms {
		formResponses[i] = s.MapToFormResponse(form)
	}

	return formResponses, total, nil
}

func (s *FormService) GetById(ctx context.Context, userId string, formId string) (*types.FormResponse, error) {

	form, err := s.formRepo.GetById(ctx, formId)
	if err != nil {
		return nil, err
	}

	if form.CreatedBy != userId {
		return nil, errors.Unauthorized("")
	}

	return s.MapToFormResponse(form), nil
}

func (s *FormService) Create(ctx context.Context, userId string, projectId string, form *types.CreateFormRequest) (*types.FormResponse, error) {

	project, err := s.projectRepo.GetWithTeam(ctx, projectId)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, errors.NotFound("Project")
	}
	if project.Team.OwnerID == nil || *project.Team.OwnerID != userId {
		return nil, errors.Unauthorized("")
	}

	status := models.FormStatusDraft
	trueVal := true
	falseVal := false

	formModel := &models.Form{
		ID:                  utils.GenerateUUID(),
		ProjectID:           projectId,
		TeamID:              project.TeamID,
		Name:                form.Name,
		Description:         form.Description,
		Status:              &status,
		LogoURL:             nil,
		Theme:               nil,
		OpenAt:              nil,
		CloseAt:             nil,
		CustomDomain:        nil,
		MaxResponses:        nil,
		Metadata:            nil,
		AllowAnonymous:      &trueVal,
		MultipleSubmissions: &trueVal,
		RequireConsent:      &falseVal,
		PasswordProtected:   &falseVal,
		Valid:               true,
		CreatedBy:           userId,
		UniqueSubdomain:     utils.GenerateRandomString(6),
		CreatedAt:           utils.GetCurrentTime(),
		UpdatedAt:           utils.GetCurrentTime(),
	}

	createdForm, err := s.formRepo.Create(ctx, formModel)
	if err != nil {
		return nil, err
	}
	return s.MapToFormResponse(createdForm), nil
}

// func (s *FormService) Update(ctx context.Context, userId string, formId string, updates *map[string]interface{}) error {

// 	form, err := s.formRepo.GetWithTeam(ctx, formId)
// 	if err != nil {
// 		return err
// 	}

// 	if form == nil {
// 		return errors.NotFound("Form")
// 	}

// 	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
// 		return errors.Unauthorized("")
// 	}

// 	(*updates)["updatedAt"] = utils.GetCurrentTime()

// 	return s.formRepo.Update(ctx, form.ID, updates)
// }

func (s *FormService) UpdateStatus(ctx context.Context, userId string, formId string, status models.FormStatus) error {

	form, err := s.formRepo.GetWithTeam(ctx, formId)
	if err != nil {
		return err
	}

	if form == nil {
		return errors.NotFound("Form")
	}

	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return errors.Unauthorized("")
	}

	return s.formRepo.UpdateStatus(ctx, form.ID, status)
}

func (s *FormService) Delete(ctx context.Context, userId string, formId string) error {

	form, err := s.formRepo.GetWithTeam(ctx, formId)
	if err != nil {
		return err
	}

	if form == nil {
		return errors.NotFound("Form")
	}

	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return errors.Unauthorized("")
	}

	return s.formRepo.Delete(ctx, form.ID)
}

func (s *FormService) UpdateSequence(ctx context.Context, userId string, formId string, sequences []*types.SequenceItem) error {

	log.Println("Updating sequence for form:", sequences)

	form, err := s.formRepo.GetWithTeam(ctx, formId)
	if err != nil {
		return err
	}
	if form == nil {
		return errors.NotFound("Form")
	}
	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return errors.Unauthorized("")
	}

	err = s.formRepo.UpdateQuestionSequence(ctx, formId, sequences)
	if err != nil {
		return err
	}

	return nil
}

func (s *FormService) MapToEdgeResponse(edge *models.Edge) *types.EdgeResponse {
	return &types.EdgeResponse{
		ID:           edge.ID,
		FormID:       edge.FormID,
		SourceNodeID: edge.SourceNodeID,
		TargetNodeID: edge.TargetNodeID,
		Condition:    edge.Condition,
	}
}

func (s *FormService) MapToQuestionResponse(question *models.Question) *types.QuestionResponse {
	optionResp := make([]*types.QueOptionResponse, len(question.Options))
	for i := range question.Options {
		optionResp[i] = s.MapToQuestionOptionResponse(&question.Options[i])
	}

	return &types.QuestionResponse{
		ID:          question.ID,
		FormID:      question.FormID,
		Title:       question.Title,
		Placeholder: question.Placeholder,
		Description: question.Description,
		Required:    question.Required,
		Type:        question.Type,
		Metadata:    question.Metadata,
		PosX:        question.PosX,
		PosY:        question.PosY,
		SortOrder:   question.SortOrder,
		Options:     optionResp,
	}
}

func (s *FormService) MapToQuestionOptionResponse(option *models.QuestionOption) *types.QueOptionResponse {
	return &types.QueOptionResponse{
		ID:         option.ID,
		QuestionID: option.QuestionID,
		Label:      option.Label,
		Value:      option.Value,
		SortOrder:  option.SortOrder,
	}
}

func (s *FormService) MapToFormResponse(form *models.Form) *types.FormResponse {
	description := ""
	if form.Description != nil {
		description = *form.Description
	}
	theme := ""
	if form.Theme != nil {
		theme = *form.Theme
	}
	logoURL := ""
	if form.LogoURL != nil {
		logoURL = *form.LogoURL
	}
	status := models.FormStatusDraft
	if form.Status != nil {
		status = *form.Status
	}
	uniqueSubdomain := ""
	if form.UniqueSubdomain != nil {
		uniqueSubdomain = *form.UniqueSubdomain
	}
	customDomain := ""
	if form.CustomDomain != nil {
		customDomain = *form.CustomDomain
	}

	edgeResp := make([]*types.EdgeResponse, len(form.Edges))
	for i := range form.Edges {
		edgeResp[i] = s.MapToEdgeResponse(&form.Edges[i])
	}

	questionResp := make([]*types.QuestionResponse, len(form.Questions))
	for i := range form.Questions {
		questionResp[i] = s.MapToQuestionResponse(&form.Questions[i])
	}

	return &types.FormResponse{
		ID:                  form.ID,
		Name:                form.Name,
		Description:         description,
		TeamID:              form.TeamID,
		ProjectID:           form.ProjectID,
		Theme:               theme,
		LogoURL:             logoURL,
		MaxResponses:        form.MaxResponses,
		OpenAt:              utils.GetIsoDateTime(form.OpenAt),
		CloseAt:             utils.GetIsoDateTime(form.CloseAt),
		Status:              status,
		UniqueSubdomain:     uniqueSubdomain,
		CustomDomain:        customDomain,
		RequireConsent:      form.RequireConsent,
		AllowAnonymous:      form.AllowAnonymous,
		MultipleSubmissions: form.MultipleSubmissions,
		PasswordProtected:   form.PasswordProtected,
		AnalyticsEnabled:    form.AnalyticsEnabled,
		Metadata:            form.Metadata,
		Valid:               form.Valid,
		CreatedAt:           utils.GetIsoDateTime(form.CreatedAt),
		UpdatedAt:           utils.GetIsoDateTime(form.UpdatedAt),
		CreatedBy:           form.CreatedBy,
		FormPageType:        form.FormPageType,
		ResponseCount:       form.ResponseCount,
		Questions:           questionResp,
		Edges:               edgeResp,
	}
}
