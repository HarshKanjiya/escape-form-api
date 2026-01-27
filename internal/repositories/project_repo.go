package repositories

import (
	"log"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	q *query.Query
}

func NewProjectRepo(db *gorm.DB) *ProjectRepo {
	return &ProjectRepo{
		q: query.Use(db),
	}
}

func (r *ProjectRepo) Get(ctx *fiber.Ctx, pagination *types.PaginationQuery, valid bool, teamId string) ([]*types.ProjectResponse, int, error) {

	userId := ctx.Locals("user_id").(string)

	if teamId != "" {
		// Authorization Check
		_, err := r.q.WithContext(ctx.Context()).
			Team.Where(r.q.Team.ID.Eq(teamId), r.q.Team.OwnerID.Eq(userId), r.q.Team.Valid.Is(true)).
			First()
		if err != nil {
			log.Printf("Team not found or not owned by user: %v", err)
			return []*types.ProjectResponse{}, 0, nil
		}

		// Query Projects
		p := r.q.Project
		baseQuery := r.q.WithContext(ctx.Context()).
			Project.Where(p.TeamID.Eq(teamId), p.Valid.Is(valid))

		if pagination.Search != "" {
			baseQuery = baseQuery.Where(p.Name.Like("%" + pagination.Search + "%"))
		}

		// Get total count without pagination
		totalCount, err := baseQuery.Count()
		if err != nil {
			log.Printf("Error counting projects: %v", err)
			return nil, 0, err
		}

		// Apply pagination for fetching projects
		query := baseQuery.Limit(pagination.Limit).Offset((pagination.Page - 1) * pagination.Limit)

		projects, err := query.Find()
		if err != nil {
			log.Printf("Error fetching projects: %v", err)
			return nil, 0, err
		}

		if len(projects) == 0 {
			log.Printf("No projects found for team %s", teamId)
			return []*types.ProjectResponse{}, int(totalCount), nil
		}

		projectIDs := make([]string, len(projects))
		for i, project := range projects {
			projectIDs[i] = project.ID
		}

		formCounts := make(map[string]int)
		var results []struct {
			ProjectID string
			Count     int
		}
		err = r.q.WithContext(ctx.Context()).
			Form.Select(r.q.Form.ProjectID, r.q.Form.ID.Count().As("count")).
			Where(r.q.Form.ProjectID.In(projectIDs...)).
			Group(r.q.Form.ProjectID).
			Scan(&results)
		if err != nil {
			log.Printf("Error fetching form counts: %v", err)
		} else {
			for _, res := range results {
				formCounts[res.ProjectID] = res.Count
			}
		}

		var projectResponses []*types.ProjectResponse
		for _, project := range projects {
			description := ""
			if project.Description != nil {
				description = *project.Description
			}

			projectResponses = append(projectResponses, &types.ProjectResponse{
				ID:          project.ID,
				Name:        project.Name,
				Description: description,
				TeamID:      project.TeamID,
				Valid:       project.Valid,
				CreatedAt:   utils.GetIsoDateTime(project.CreatedAt),
				UpdatedAt:   utils.GetIsoDateTime(project.UpdatedAt),
				FormCount:   formCounts[project.ID],
			})
		}

		return projectResponses, int(totalCount), nil
	}

	// Original logic: fetch for all user's teams
	// First, get user's team IDs
	teamQuery := r.q.WithContext(ctx.Context()).
		Team.Where(r.q.Team.OwnerID.Eq(userId), r.q.Team.Valid.Is(true))

	userTeams, err := teamQuery.Find()
	if err != nil {
		log.Printf("Error fetching user teams: %v", err)
		return nil, 0, err
	}

	if len(userTeams) == 0 {
		log.Printf("No teams found for user %s", userId)
		return []*types.ProjectResponse{}, 0, nil
	}

	teamIDs := make([]string, len(userTeams))
	for i, team := range userTeams {
		teamIDs[i] = team.ID
	}

	// Now, query projects
	p := r.q.Project
	baseQuery := r.q.WithContext(ctx.Context()).
		Project.Where(p.TeamID.In(teamIDs...), p.Valid.Is(valid))

	if pagination.Search != "" {
		baseQuery = baseQuery.Where(p.Name.Like("%" + pagination.Search + "%"))
	}

	// Get total count without pagination
	totalCount, err := baseQuery.Count()
	if err != nil {
		log.Printf("Error counting projects: %v", err)
		return nil, 0, err
	}

	// Apply pagination for fetching projects
	query := baseQuery.Limit(pagination.Limit).Offset((pagination.Page - 1) * pagination.Limit)

	projects, err := query.Find()
	if err != nil {
		log.Printf("Error fetching projects: %v", err)
		return nil, 0, err
	}

	if len(projects) == 0 {
		log.Printf("No projects found for user %s", userId)
		return []*types.ProjectResponse{}, int(totalCount), nil
	}

	projectIDs := make([]string, len(projects))
	for i, project := range projects {
		projectIDs[i] = project.ID
	}

	// Optimized: Single query to count forms per project
	formCounts := make(map[string]int)
	var results []struct {
		ProjectID string
		Count     int
	}
	err = r.q.WithContext(ctx.Context()).
		Form.Select(r.q.Form.ProjectID, r.q.Form.ID.Count().As("count")).
		Where(r.q.Form.ProjectID.In(projectIDs...)).
		Group(r.q.Form.ProjectID).
		Scan(&results)
	if err != nil {
		log.Printf("Error fetching form counts: %v", err)
		// Continue without counts (set to 0)
	} else {
		for _, res := range results {
			formCounts[res.ProjectID] = res.Count
		}
	}

	var projectResponses []*types.ProjectResponse
	for _, project := range projects {
		description := ""
		if project.Description != nil {
			description = *project.Description
		}
		createdAt := ""
		if project.CreatedAt != nil {
			createdAt = project.CreatedAt.UTC().Format("2006-01-02T15:04:05Z07:00")
		}
		updatedAt := ""
		if project.UpdatedAt != nil {
			updatedAt = project.UpdatedAt.UTC().Format("2006-01-02T15:04:05Z07:00")
		}
		projectResponses = append(projectResponses, &types.ProjectResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: description,
			TeamID:      project.TeamID,
			Valid:       project.Valid,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			FormCount:   formCounts[project.ID],
		})
	}

	return projectResponses, int(totalCount), nil
}

func (r *ProjectRepo) GetById(ctx *fiber.Ctx, projectId string) (*types.ProjectResponse, error) {

	userId := ctx.Locals("user_id").(string)

	// Validate the project belongs to the user
	project, err := r.q.WithContext(ctx.Context()).
		Project.Where(r.q.Project.ID.Eq(projectId)).
		Join(r.q.Team, r.q.Project.TeamID.EqCol(r.q.Team.ID)).
		Where(r.q.Team.OwnerID.Eq(userId), r.q.Team.Valid.Is(true), r.q.Project.Valid.Is(true)).
		First()
	if err != nil {
		log.Printf("Project not found or not owned by user: %v", err)
		return nil, err
	}

	// Get form count for this project
	var formCount int
	err = r.q.WithContext(ctx.Context()).
		Form.Select(r.q.Form.ID.Count()).
		Where(r.q.Form.ProjectID.Eq(projectId), r.q.Form.Valid.Is(true)).
		Scan(&formCount)
	if err != nil {
		log.Printf("Error fetching form count: %v", err)
		formCount = 0
	}

	description := ""
	if project.Description != nil {
		description = *project.Description
	}
	createdAt := ""
	if project.CreatedAt != nil {
		createdAt = project.CreatedAt.UTC().Format("2006-01-02T15:04:05Z07:00")
	}
	updatedAt := ""
	if project.UpdatedAt != nil {
		updatedAt = project.UpdatedAt.UTC().Format("2006-01-02T15:04:05Z07:00")
	}

	projectResponse := &types.ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: description,
		TeamID:      project.TeamID,
		Valid:       project.Valid,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		FormCount:   formCount,
	}

	return projectResponse, nil
}

func (r *ProjectRepo) Create(ctx *fiber.Ctx, project *types.ProjectDto) (*models.Project, error) {

	projectModel := &models.Project{
		ID:          uuid.New().String(),
		Name:        project.Name,
		Description: project.Description,
		TeamID:      project.TeamID,
		Valid:       true,
	}

	err := r.q.WithContext(ctx.Context()).Project.Create(projectModel)
	if err != nil {
		return nil, err
	}
	return projectModel, nil
}

func (r *ProjectRepo) Update(ctx *fiber.Ctx, project *models.Project) (bool, error) {
	p := r.q.Project
	_, err := r.q.WithContext(ctx.Context()).
		Project.Where(p.ID.Eq(project.ID)).
		Updates(project)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ProjectRepo) Delete(ctx *fiber.Ctx, projectId string) (bool, error) {
	_, err := r.q.WithContext(ctx.Context()).
		Project.Where(r.q.Project.ID.Eq(projectId)).
		UpdateColumn(r.q.Project.Valid, false)
	if err != nil {
		return false, err
	}
	return true, nil
}
