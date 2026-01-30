package repositories

import (
	"context"
	"log"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectRepo {
	return &ProjectRepo{
		db: db,
	}
}

func (r *ProjectRepo) Get(ctx *fiber.Ctx, pagination *types.PaginationQuery, valid bool, teamId string) ([]*types.ProjectResponse, int, error) {

	userId := ctx.Locals("user_id").(string)

	if teamId != "" {
		// Authorization Check: ensure the team is owned by the user
		var team models.Team
		err := r.db.WithContext(ctx.Context()).Model(&models.Team{}).Where("id = ? AND owner_id = ? AND valid = ?", teamId, userId, true).First(&team).Error
		if err != nil {
			log.Printf("Team not found or not owned by user: %v", err)
			return []*types.ProjectResponse{}, 0, nil
		}

		// Query Projects
		baseQuery := r.db.WithContext(ctx.Context()).Model(&models.Project{}).Where("team_id = ? AND valid = ?", teamId, valid)

		if pagination.Search != "" {
			baseQuery = baseQuery.Where("name LIKE ?", "%"+pagination.Search+"%")
		}

		// Get total count without pagination
		var totalCount int64
		err = baseQuery.Count(&totalCount).Error
		if err != nil {
			log.Printf("Error counting projects: %v", err)
			return nil, 0, err
		}

		// Apply pagination for fetching projects
		var projects []models.Project
		err = baseQuery.Limit(pagination.Limit).Offset((pagination.Page - 1) * pagination.Limit).Find(&projects).Error
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
		err = r.db.WithContext(ctx.Context()).Model(&models.Form{}).Select("project_id, count(*) as count").Where("project_id IN ? AND valid = ?", projectIDs, true).Group("project_id").Scan(&results).Error
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
	var userTeams []models.Team
	err := r.db.WithContext(ctx.Context()).Model(&models.Team{}).Where("owner_id = ? AND valid = ?", userId, true).Find(&userTeams).Error
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
	baseQuery := r.db.WithContext(ctx.Context()).Model(&models.Project{}).Where("team_id IN ? AND valid = ?", teamIDs, valid)

	if pagination.Search != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+pagination.Search+"%")
	}

	// Get total count without pagination
	var totalCount int64
	err = baseQuery.Count(&totalCount).Error
	if err != nil {
		log.Printf("Error counting projects: %v", err)
		return nil, 0, err
	}

	// Apply pagination for fetching projects
	var projects []models.Project
	err = baseQuery.Limit(pagination.Limit).Offset((pagination.Page - 1) * pagination.Limit).Find(&projects).Error
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
	err = r.db.WithContext(ctx.Context()).Model(&models.Form{}).Select("project_id, count(*) as count").Where("project_id IN ? AND valid = ?", projectIDs, true).Group("project_id").Scan(&results).Error
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

func (r *ProjectRepo) GetById(ctx *fiber.Ctx, projectId string) (*types.ProjectResponse, error) {

	userId := ctx.Locals("user_id").(string)

	// Validate the project belongs to the user
	var project models.Project
	err := r.db.WithContext(ctx.Context()).Model(&models.Project{}).Joins("JOIN teams ON projects.team_id = teams.id").Where("projects.id = ? AND teams.owner_id = ? AND teams.valid = ? AND projects.valid = ?", projectId, userId, true, true).First(&project).Error
	if err != nil {
		log.Printf("Project not found or not owned by user: %v", err)
		return nil, err
	}

	// Get form count for this project
	var formCount int64
	err = r.db.WithContext(ctx.Context()).Model(&models.Form{}).Where("project_id = ? AND valid = ?", projectId, true).Count(&formCount).Error
	if err != nil {
		log.Printf("Error fetching form count: %v", err)
		formCount = 0
	}

	description := ""
	if project.Description != nil {
		description = *project.Description
	}

	projectResponse := &types.ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: description,
		TeamID:      project.TeamID,
		Valid:       project.Valid,
		CreatedAt:   utils.GetIsoDateTime(project.CreatedAt),
		UpdatedAt:   utils.GetIsoDateTime(project.UpdatedAt),
		FormCount:   int(formCount),
	}

	return projectResponse, nil
}

func (r *ProjectRepo) GetWithTeam(ctx context.Context, userId string, projectId string) (*models.Project, error) {
	var project models.Project
	err := r.db.WithContext(ctx).Model(&models.Project{}).Joins("JOIN teams ON projects.team_id = teams.id").Where("projects.id = ? AND teams.owner_id = ? AND teams.valid = ? AND projects.valid = ?", projectId, userId, true, true).First(&project).Error
	if err != nil {
		log.Printf("Project not found or not owned by user: %v", err)
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepo) Create(ctx *fiber.Ctx, project *types.ProjectDto) (*models.Project, error) {

	projectModel := &models.Project{
		ID:          uuid.New().String(),
		Name:        project.Name,
		Description: project.Description,
		TeamID:      project.TeamID,
		Valid:       true,
	}

	err := r.db.WithContext(ctx.Context()).Create(projectModel).Error
	if err != nil {
		return nil, err
	}
	return projectModel, nil
}

func (r *ProjectRepo) Update(ctx *fiber.Ctx, project *models.Project) (bool, error) {
	err := r.db.WithContext(ctx.Context()).Model(&models.Project{}).Where("id = ?", project.ID).Updates(project).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ProjectRepo) Delete(ctx *fiber.Ctx, projectId string) (bool, error) {
	err := r.db.WithContext(ctx.Context()).Model(&models.Project{}).Where("id = ?", projectId).Update("valid", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
