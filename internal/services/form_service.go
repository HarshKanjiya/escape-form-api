package services

import (
	"context"
	"log"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/mapper"
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
		formResponses[i] = mapper.ToFormResponse(form)
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
	return mapper.ToFormResponse(form), nil
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
	return mapper.ToFormResponse(createdForm), nil
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
