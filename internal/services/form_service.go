package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/mapper"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"gorm.io/datatypes"
)

type IFormService interface {
	Get(ctx context.Context, userId string, pagination *types.PaginationQuery, projectId string, status string) ([]*types.FormResponse, int64, error)
	GetById(ctx context.Context, userId string, formId string) (*types.FormResponse, error)
	Create(ctx context.Context, userId string, projectId string, form *types.CreateFormRequest) (*types.FormResponse, error)
	Update(ctx context.Context, userId string, formId string, updates map[string]interface{}) error
	UpdateStatus(ctx context.Context, userId string, formId string, status models.FormStatus) error
	Delete(ctx context.Context, userId string, formId string) error

	UpdateSequence(ctx context.Context, userId string, formId string, sequences []*types.SequenceItem) error
	Demo(ctx context.Context, userId string) (string, error)
	Publish(ctx context.Context, formId string) (*types.FormResponse, error)
	Unpublish(ctx context.Context, formId string) (*types.FormResponse, error)
}

type FormService struct {
	formRepo        repositories.IFormRepo
	projectRepo     repositories.IProjectRepo
	formVersionRepo repositories.IFormVersionRepo
	questionRepo    repositories.IQuestionRepo
	edgeRepo        repositories.IEdgeRepo
}

func NewFormService(formRepo repositories.IFormRepo, projectRepo repositories.IProjectRepo, formVersionRepo repositories.IFormVersionRepo, questionRepo repositories.IQuestionRepo, edgeRepo repositories.IEdgeRepo) *FormService {
	return &FormService{
		formRepo:        formRepo,
		projectRepo:     projectRepo,
		formVersionRepo: formVersionRepo,
		questionRepo:    questionRepo,
		edgeRepo:        edgeRepo,
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

func (s *FormService) Update(ctx context.Context, userId string, formId string, updates map[string]interface{}) error {

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

	updates["updatedAt"] = utils.GetCurrentTime()

	return s.formRepo.Update(ctx, form.ID, updates)
}

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

func (s *FormService) Demo(ctx context.Context, userId string) (string, error) {
	// Demo function - just return a simple message
	return "Demo function executed for user: " + userId, nil
}

func (s *FormService) Publish(ctx context.Context, formId string) (*types.FormResponse, error) {
	// Get form details
	form, err := s.formRepo.GetById(ctx, formId)
	if err != nil {
		return nil, err
	}
	if form == nil {
		return nil, errors.NotFound("Form")
	}

	// Get all questions with options
	questions, err := s.questionRepo.GetQuestions(ctx, formId)
	if err != nil {
		return nil, err
	}

	// Get all edges
	edges, err := s.edgeRepo.Get(ctx, formId)
	if err != nil {
		return nil, err
	}

	// Create snapshot
	snapshot := s.createPublishSnapshot(form, questions, edges)

	// Convert snapshot to JSONB
	schemaBytes, err := json.Marshal(snapshot)
	if err != nil {
		return nil, errors.Internal(err)
	}

	// Get latest version number
	latestVersion, err := s.formVersionRepo.GetLatestVersion(ctx, formId)
	var versionNumber int
	if err != nil {
		return nil, err
	}
	if latestVersion != nil {
		versionNumber = latestVersion.VersionNumber + 1
	} else {
		versionNumber = 1
	}

	now := time.Now()

	// Create form version
	formVersion := &models.FormVersion{
		ID:            utils.GenerateUUID(),
		FormID:        formId,
		VersionNumber: versionNumber,
		Schema:        datatypes.JSON(schemaBytes),
		CreatedAt:     now,
		PublishedAt:   &now,
	}

	createdVersion, err := s.formVersionRepo.Create(ctx, formVersion)
	if err != nil {
		return nil, err
	}

	// Update form status and revision
	editorRevision := form.EditorRevision
	if editorRevision == nil {
		defaultRevision := 1
		editorRevision = &defaultRevision
	}

	updates := map[string]interface{}{
		"status":              models.FormStatusPublished,
		"publishedVersionId": createdVersion.ID,
		"publishedRevision":  *editorRevision,
		"updatedAt":          time.Now(),
	}

	err = s.formRepo.Update(ctx, formId, updates)
	if err != nil {
		return nil, err
	}

	// Get updated form
	updatedForm, err := s.formRepo.GetById(ctx, formId)
	if err != nil {
		return nil, err
	}

	return mapper.ToFormResponse(updatedForm), nil
}

func (s *FormService) createPublishSnapshot(form *models.Form, questions []*models.Question, edges []*models.Edge) *types.PublishVersionSnapshot {
	// Map questions
	publishedQuestions := make([]types.PublishedQuestion, len(questions))
	for i, q := range questions {
		publishedOptions := make([]types.PublishedQuestionOption, len(q.Options))
		for j, opt := range q.Options {
			publishedOptions[j] = types.PublishedQuestionOption{
				ID:        opt.ID,
				Label:     opt.Label,
				Value:     opt.Value,
				SortOrder: opt.SortOrder,
			}
		}

		publishedQuestions[i] = types.PublishedQuestion{
			ID:          q.ID,
			Title:       q.Title,
			Placeholder: q.Placeholder,
			Description: q.Description,
			Required:    q.Required,
			Type:        string(q.Type),
			Metadata:    q.Metadata,
			PosX:        q.PosX,
			PosY:        q.PosY,
			SortOrder:   q.SortOrder,
			Options:     publishedOptions,
		}
	}

	// Map edges
	publishedEdges := make([]types.PublishedEdge, len(edges))
	for i, e := range edges {
		publishedEdges[i] = types.PublishedEdge{
			ID:           e.ID,
			SourceNodeID: e.SourceNodeID,
			TargetNodeID: e.TargetNodeID,
			Condition:    e.Condition,
		}
	}

	// Create snapshot
	return &types.PublishVersionSnapshot{
		Name:                form.Name,
		Description:         form.Description,
		Theme:               form.Theme,
		LogoURL:             form.LogoURL,
		RequireConsent:      form.RequireConsent,
		AllowAnonymous:      form.AllowAnonymous,
		MultipleSubmissions: form.MultipleSubmissions,
		PasswordProtected:   form.PasswordProtected,
		FormPageType:        string(form.FormPageType),
		Metadata:            form.Metadata,
		Questions:           publishedQuestions,
		Edges:               publishedEdges,
	}
}

func (s *FormService) Unpublish(ctx context.Context, formId string) (*types.FormResponse, error) {
	err := s.formRepo.UpdateStatus(ctx, formId, models.FormStatusDraft)
	if err != nil {
		return nil, err
	}

	form, err := s.formRepo.GetById(ctx, formId)
	if err != nil {
		return nil, err
	}

	return mapper.ToFormResponse(form), nil
}
