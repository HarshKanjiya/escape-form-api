package repositories

import (
	"log"

	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
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

func (r *FormRepo) Get(ctx *fiber.Ctx, pagination *types.PaginationQuery, valid bool, projectId string) ([]*types.FormResponse, error) {

	userId := ctx.Locals("user_id").(string)

	// First, validate the project belongs to the user
	project, err := r.q.WithContext(ctx.Context()).
		Project.Where(r.q.Project.ID.Eq(projectId)).
		Join(r.q.Team, r.q.Project.TeamID.EqCol(r.q.Team.ID)).
		Where(r.q.Team.OwnerID.Eq(userId), r.q.Team.Valid.Is(true), r.q.Project.Valid.Is(true)).
		First()
	if err != nil {
		log.Printf("Project not found or not owned by user: %v", err)
		return []*types.FormResponse{}, nil
	}
	_ = project // ensure exists

	// Query forms for this project
	f := r.q.Form
	query := r.q.WithContext(ctx.Context()).
		Form.Where(f.ProjectID.Eq(projectId), f.Valid.Is(valid))

	if pagination.Search != "" {
		query = query.Where(f.Name.Like("%" + pagination.Search + "%"))
	}

	query.Limit(pagination.Limit)
	query.Offset((pagination.Page - 1) * pagination.Limit)

	forms, err := query.Find()
	if err != nil {
		log.Printf("Error fetching forms: %v", err)
		return nil, err
	}

	if len(forms) == 0 {
		log.Printf("No forms found for project %s", projectId)
		return []*types.FormResponse{}, nil
	}

	formIDs := make([]string, len(forms))
	for i, form := range forms {
		formIDs[i] = form.ID
	}

	// Optimized: Single query to count responses per form
	responseCounts := make(map[string]int)
	var results []struct {
		FormID string
		Count  int
	}
	err = r.q.WithContext(ctx.Context()).
		Response.Select(r.q.Response.FormID, r.q.Response.ID.Count().As("count")).
		Where(r.q.Response.FormID.In(formIDs...)).
		Group(r.q.Response.FormID).
		Scan(&results)
	if err != nil {
		log.Printf("Error fetching response counts: %v", err)
		// Continue without counts (set to 0)
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
		openAt := ""
		if form.OpenAt != nil {
			openAt = form.OpenAt.UTC().Format("2006-01-02T15:04:05Z07:00")
		}
		closeAt := ""
		if form.CloseAt != nil {
			closeAt = form.CloseAt.UTC().Format("2006-01-02T15:04:05Z07:00")
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
		metadata := ""
		if form.Metadata != nil {
			metadata = *form.Metadata
		}
		createdAt := ""
		if form.CreatedAt != nil {
			createdAt = form.CreatedAt.UTC().Format("2006-01-02T15:04:05Z07:00")
		}
		updatedAt := ""
		if form.UpdatedAt != nil {
			updatedAt = form.UpdatedAt.UTC().Format("2006-01-02T15:04:05Z07:00")
		}
		formResponses = append(formResponses, &types.FormResponse{
			ID:                  form.ID,
			Name:                form.Name,
			Description:         description,
			TeamID:              form.TeamID,
			ProjectID:           form.ProjectID,
			Theme:               theme,
			LogoURL:             logoURL,
			MaxResponses:        form.MaxResponses,
			OpenAt:              openAt,
			CloseAt:             closeAt,
			Status:              status,
			UniqueSubdomain:     uniqueSubdomain,
			CustomDomain:        customDomain,
			RequireConsent:      form.RequireConsent,
			AllowAnonymous:      form.AllowAnonymous,
			MultipleSubmissions: form.MultipleSubmissions,
			PasswordProtected:   form.PasswordProtected,
			AnalyticsEnabled:    form.AnalyticsEnabled,
			Valid:               form.Valid,
			Metadata:            metadata,
			CreatedBy:           form.CreatedBy,
			CreatedAt:           createdAt,
			UpdatedAt:           updatedAt,
			FormPageType:        string(form.FormPageType),
			ResponseCount:       responseCounts[form.ID],
		})
	}

	return formResponses, nil
}
