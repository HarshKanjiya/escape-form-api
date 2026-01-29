package repositories

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FormRepo struct {
	q *query.Query
}

func NewFormRepo(db *gorm.DB) *FormRepo {
	return &FormRepo{
		q: query.Use(db),
	}
}

func (r *FormRepo) Get(ctx context.Context, pagination *types.PaginationQuery, valid bool, projectId string) ([]*types.FormResponse, int, error) {

	project, err := r.q.WithContext(ctx).
		Project.Where(r.q.Project.ID.Eq(projectId)).
		First()
	if err != nil {
		log.Printf("Project not found: %v", err)
		return []*types.FormResponse{}, 0, err
	}
	_ = project

	f := r.q.Form
	baseQuery := r.q.WithContext(ctx).
		Form.Where(f.ProjectID.Eq(projectId), f.Valid.Is(valid))

	if pagination.Search != "" {
		baseQuery = baseQuery.Where(f.Name.Lower().Like("%" + strings.ToLower(pagination.Search) + "%"))
	}

	// Get total count without pagination
	totalCount, err := baseQuery.Count()
	if err != nil {
		log.Printf("Error counting forms: %v", err)
		return nil, 0, err
	}

	// Apply pagination for fetching forms
	query := baseQuery.Limit(pagination.Limit).Offset((pagination.Page - 1) * pagination.Limit)

	forms, err := query.Find()
	if err != nil {
		log.Printf("Error fetching forms: %v", err)
		return nil, 0, err
	}

	if len(forms) == 0 {
		log.Printf("No forms found for project %s", projectId)
		return []*types.FormResponse{}, int(totalCount), nil
	}

	formIDs := make([]string, len(forms))
	for i, form := range forms {
		formIDs[i] = form.ID
	}

	responseCounts := make(map[string]int)
	var results []struct {
		FormID string
		Count  int
	}
	err = r.q.WithContext(ctx).
		Response.Select(r.q.Response.FormID, r.q.Response.ID.Count().As("count")).
		Where(r.q.Response.FormID.In(formIDs...)).
		Group(r.q.Response.FormID).
		Scan(&results)
	if err != nil {
		log.Printf("Error fetching response counts: %v", err)
	} else {
		for _, res := range results {
			responseCounts[res.FormID] = res.Count
		}
	}

	var formResponses []*types.FormResponse
	for _, form := range forms {
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
		status := ""
		if form.Status != nil {
			status = string(*form.Status)
		}
		uniqueSubdomain := ""
		if form.UniqueSubdomain != nil {
			uniqueSubdomain = *form.UniqueSubdomain
		}
		customDomain := ""
		if form.CustomDomain != nil {
			customDomain = *form.CustomDomain
		}
		metadata := map[string]interface{}{}
		if form.Metadata != nil {
			if m, ok := (*form.Metadata).(map[string]interface{}); ok {
				metadata = m
			}
		}
		metadataBytes, _ := json.Marshal(metadata)
		metadataStr := string(metadataBytes)

		formResponses = append(formResponses, &types.FormResponse{
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
			Valid:               form.Valid,
			Metadata:            metadataStr,
			CreatedBy:           form.CreatedBy,
			CreatedAt:           utils.GetIsoDateTime(form.CreatedAt),
			UpdatedAt:           utils.GetIsoDateTime(form.UpdatedAt),
			FormPageType:        string(form.FormPageType),
			ResponseCount:       responseCounts[form.ID],
		})
	}

	return formResponses, int(totalCount), nil
}

func (r *FormRepo) GetById(ctx context.Context, formId string) (*types.FormResponse, error) {

	// Validate the form belongs to the user
	form, err := r.q.WithContext(ctx).
		Form.Where(r.q.Form.ID.Eq(formId)).
		Preload(r.q.Form.Questions).
		Preload(r.q.Form.Edges).
		First()
	if err != nil {
		log.Printf("Form not found or not owned by user: %v", err)
		return nil, err
	}

	// Get response count for this form
	responseCount, err := r.q.WithContext(ctx).
		Response.Where(r.q.Response.FormID.Eq(formId)).
		Count()
	if err != nil {
		log.Printf("Error fetching response count: %v", err)
		responseCount = 0
	}

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
	status := ""
	if form.Status != nil {
		status = string(*form.Status)
	}
	uniqueSubdomain := ""
	if form.UniqueSubdomain != nil {
		uniqueSubdomain = *form.UniqueSubdomain
	}
	customDomain := ""
	if form.CustomDomain != nil {
		customDomain = *form.CustomDomain
	}
	metadata := map[string]interface{}{}
	if form.Metadata != nil {
		if m, ok := (*form.Metadata).(map[string]interface{}); ok {
			metadata = m
		}
	}
	metadataBytes, _ := json.Marshal(metadata)
	metadataStr := string(metadataBytes)

	formResponse := &types.FormResponse{
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
		Valid:               form.Valid,
		Metadata:            metadataStr,
		CreatedBy:           form.CreatedBy,
		CreatedAt:           utils.GetIsoDateTime(form.CreatedAt),
		UpdatedAt:           utils.GetIsoDateTime(form.UpdatedAt),
		FormPageType:        string(form.FormPageType),
		ResponseCount:       int(responseCount),
	}

	// Convert Questions to []any
	if len(form.Questions) > 0 {
		questions := make([]any, len(form.Questions))
		for i, q := range form.Questions {
			questions[i] = q
		}
		formResponse.Questions = questions
	}

	// Convert Edges to []any
	if len(form.Edges) > 0 {
		edges := make([]any, len(form.Edges))
		for i, e := range form.Edges {
			edges[i] = e
		}
		formResponse.Edges = edges
	}

	return formResponse, nil
}

func (r *FormRepo) Create(ctx context.Context, userId string, formDto *types.CreateFormDto) (*types.FormResponse, error) {

	project, err := r.q.WithContext(ctx).
		Project.Where(r.q.Project.ID.Eq(formDto.ProjectID)).
		First()

	if err != nil {
		log.Printf("Project not found: %v", err)
		return nil, err
	}

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

	err = r.q.WithContext(ctx).Form.Create(form)
	if err != nil {
		log.Printf("Error creating form: %v", err)
		return nil, err
	}
	return r.GetById(ctx, form.ID)
}

func (r *FormRepo) UpdateStatus(
	ctx context.Context,
	formId string,
	status models.FormStatus,
) error {

	_, err := r.q.WithContext(ctx).
		Form.
		Where(r.q.Form.ID.Eq(formId)).
		UpdateColumn(r.q.Form.Status, status)

	return err
}

func (r *FormRepo) GetWithTeam(ctx context.Context, userId string, formId string) (*models.Form, error) {
	form, err := r.q.WithContext(ctx).
		Form.Where(r.q.Form.ID.Eq(formId)).
		Join(r.q.Team, r.q.Form.TeamID.EqCol(r.q.Team.ID)).
		Where(r.q.Team.OwnerID.Eq(userId), r.q.Team.Valid.Is(true), r.q.Form.Valid.Is(true)).
		First()
	if err != nil {
		log.Printf("Form not found or not owned by user: %v", err)
		return nil, err
	}
	return form, nil
}
