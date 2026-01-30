package repositories

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"gorm.io/gorm"
)

type IProjectRepo interface {
	Get(ctx context.Context, userId string, pagination *types.PaginationQuery, teamId string) ([]*models.Project, int, error)
	GetById(ctx context.Context, projectId string) (*models.Project, error)
	Create(ctx context.Context, project *types.ProjectDto) (*models.Project, error)
	Update(ctx context.Context, project *models.Project) (bool, error)
	Delete(ctx context.Context, projectId string) (bool, error)
	GetWithTeam(ctx context.Context, userId string, projectId string) (*models.Project, error)
}

type ProjectRepo struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectRepo {
	return &ProjectRepo{
		db: db,
	}
}

func (r *ProjectRepo) Get(ctx context.Context, pagination *types.PaginationQuery, teamId string) ([]*types.ProjectResponse, int64, error) {

	var projects []*types.ProjectResponse
	var totalCount int64

	baseQuery := r.db.
		WithContext(ctx).
		Model(&models.Project{}).
		Where("team_id = ? AND valid = ?", teamId, true)

	if pagination.Search != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+pagination.Search+"%")
	}

	// total count
	if err := baseQuery.Count(&totalCount).Error; err != nil {
		return nil, 0, errors.Internal(err)
	}

	if err := baseQuery.
		Select(`
			projects.id,
			projects.name,
			projects.description,
			projects.teamId AS team_id,
			projects.valid,
			projects.createdAt,
			projects.updatedAt,
			COUNT(forms.id) AS form_count
		`).
		Joins(`
			LEFT JOIN forms
			ON forms.project_id = projects.id
		`).
		Group("projects.id").
		Order("projects.created_at DESC").
		Limit(pagination.Limit).
		Offset((pagination.Page - 1) * pagination.Limit).
		Scan(&projects).Error; err != nil {
		return nil, 0, errors.Internal(err)
	}

	return projects, totalCount, nil
}

func (r *ProjectRepo) GetById(ctx context.Context, projectId string) (*types.ProjectResponse, error) {

	var project types.ProjectResponse

	if err := r.db.Model(&models.Project{}).WithContext(ctx).
		Select(`
			projects.id,
			projects.name,
			projects.description,
			projects.teamId AS team_id,
			projects.valid,
			projects.createdAt,
			projects.updatedAt,
			COUNT(forms.id) AS form_count
		`).
		Joins(`
			LEFT JOIN forms
			ON forms.project_id = projects.id
		`).
		Where("projects.id = ? AND projects.valid = ?", projectId, true).
		Group("projects.id").
		Scan(&project).Error; err != nil {
		return nil, errors.Internal(err)
	}
	return &project, nil
}

func (r *ProjectRepo) GetWithTeam(ctx context.Context, userId string, projectId string) (*models.Project, error) {
	var project models.Project

}

func (r *ProjectRepo) Create(ctx context.Context, project *types.ProjectDto) (*models.Project, error) {

	// projectModel := &models.Project{
	// 	ID:          uuid.New().String(),
	// 	Name:        project.Name,
	// 	Description: project.Description,
	// 	TeamID:      project.TeamID,
	// 	Valid:       true,
	// }

	// err := r.db.WithContext(ctx.Context()).Create(projectModel).Error
	// if err != nil {
	// 	return nil, err
	// }
	// return projectModel, nil
}

func (r *ProjectRepo) Update(ctx context.Context, project *models.Project) (bool, error) {
	// err := r.db.WithContext(ctx.Context()).Model(&models.Project{}).Where("id = ?", project.ID).Updates(project).Error
	// if err != nil {
	// 	return false, err
	// }
	// return true, nil
}

func (r *ProjectRepo) Delete(ctx context.Context, projectId string) (bool, error) {
	// err := r.db.WithContext(ctx.Context()).Model(&models.Project{}).Where("id = ?", projectId).Update("valid", false).Error
	// if err != nil {
	// 	return false, err
	// }
	// return true, nil
}
