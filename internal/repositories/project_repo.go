package repositories

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IProjectRepo interface {
	Get(ctx context.Context, pagination *types.PaginationQuery, teamId string) ([]*types.ProjectResponse, int64, error)
	GetById(ctx context.Context, projectId string) (*models.Project, error)
	Create(ctx context.Context, project *models.Project) (*models.Project, error)
	Update(ctx context.Context, project *models.Project) (bool, error)
	Delete(ctx context.Context, projectId string) error
	GetWithTeam(ctx context.Context, projectId string) (*models.Project, error)
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

	baseQuery := r.db.WithContext(ctx).
		Model(&models.Project{}).
		Where(&models.Project{
			TeamID: teamId,
			Valid:  true,
		})

	if pagination.Search != "" {
		baseQuery = baseQuery.Where(
			clause.Like{
				Column: clause.Column{Name: "name"},
				Value:  "%" + pagination.Search + "%",
			},
		)
	}
	// total count
	if err := baseQuery.Count(&totalCount).Error; err != nil {
		return nil, 0, errors.Internal(err)
	}

	err := baseQuery.Select(`
        projects.id,
        projects.name,
        projects.description,
        projects."teamId",
        projects.valid,
        projects."createdAt",
        projects."updatedAt",
        (
            SELECT COUNT(*)
            FROM forms
            WHERE forms."projectId" = projects.id
        ) AS "formCount"
    `).
		Order(`projects."createdAt" DESC`).
		Limit(pagination.Limit).
		Offset((pagination.Page - 1) * pagination.Limit).
		Scan(&projects).Error

	if err != nil {
		return nil, 0, errors.Internal(err)
	}

	return projects, totalCount, nil
}

func (r *ProjectRepo) GetById(ctx context.Context, projectId string) (*models.Project, error) {

	var project models.Project

	if err := r.db.Model(&models.Project{}).WithContext(ctx).
		Select(`
			projects.id,
        	projects.name,
        	projects.description,
        	projects."teamId",
        	projects.valid,
        	projects."createdAt",
        	projects."updatedAt",
        	(
        	    SELECT COUNT(*)
        	    FROM forms
        	    WHERE forms."projectId" = projects.id
        	) AS "formCount"
		`).
		Where(&models.Project{
			ID:    projectId,
			Valid: true,
		}).
		Group("projects.id").
		Scan(&project).Error; err != nil {
		return nil, errors.Internal(err)
	}
	return &project, nil
}

func (r *ProjectRepo) GetWithTeam(ctx context.Context, projectId string) (*models.Project, error) {
	var project models.Project
	if err := r.db.WithContext(ctx).Preload("Team").
		Where(&models.Project{
			ID:    projectId,
			Valid: true,
		}).
		First(&project).Error; err != nil {
		return nil, errors.Internal(err)
	}
	return &project, nil
}

func (r *ProjectRepo) Create(ctx context.Context, project *models.Project) (*models.Project, error) {

	err := r.db.Model(&models.Project{}).WithContext(ctx).Create(project).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return project, nil
}

func (r *ProjectRepo) Update(ctx context.Context, project *models.Project) (bool, error) {
	if err := r.db.WithContext(ctx).Model(&models.Project{}).
		Where(&models.Project{
			ID:    project.ID,
			Valid: true,
		}).
		Updates(map[string]interface{}{
			"name":        project.Name,
			"description": project.Description,
			"updatedAt":   utils.GetCurrentTime(),
		}).Error; err != nil {
		return false, errors.Internal(err)
	}
	return true, nil
}

func (r *ProjectRepo) Delete(ctx context.Context, projectId string) error {

	if err := r.db.WithContext(ctx).Model(&models.Project{}).
		Where(&models.Project{
			ID:    projectId,
			Valid: true,
		}).
		Updates(map[string]interface{}{
			"valid":     false,
			"updatedAt": utils.GetCurrentTime(),
		}).Error; err != nil {
		return errors.Internal(err)
	}
	return nil
}
