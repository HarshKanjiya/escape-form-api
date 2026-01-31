package repositories

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"gorm.io/gorm"
)

type ITeamRepo interface {
	Get(ctx context.Context, userId string, pagination *types.PaginationQuery, valid bool) ([]*types.TeamResponse, int64, error)
	GetById(ctx context.Context, teamId string) (*models.Team, error)
	Create(ctx context.Context, team *models.Team) error
	Update(ctx context.Context, team *models.Team) error
	Delete(ctx context.Context, teamId string) error
}

type TeamRepo struct {
	db *gorm.DB
}

func NewTeamRepo(db *gorm.DB) *TeamRepo {
	return &TeamRepo{
		db: db,
	}
}

func (r *TeamRepo) Get(ctx context.Context, userId string, pagination *types.PaginationQuery, valid bool) ([]*types.TeamResponse, int64, error) {

	var teams []*types.TeamResponse
	var totalCount int64

	baseQuery := r.db.
		WithContext(ctx).
		Model(&models.Team{}).
		Where("valid = ?", valid)

	if pagination.Search != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+pagination.Search+"%")
	}

	// total count
	if err := baseQuery.Count(&totalCount).Error; err != nil {
		return nil, 0, errors.Internal(err)
	}

	err := baseQuery.Select(`
			teams.id,
			teams.name,
			teams.ownerId,
			teams.planId,
			teams.valid,
			teams.createdAt,
			teams.updatedAt,
			COUNT(projects.id) AS projectCount
			`).
		Joins(`LEFT JOIN projects ON projects.teamId = teams.id AND projects.valid = true`).
		Group("teams.id").
		Limit(pagination.Limit).
		Offset((pagination.Page - 1) * pagination.Limit).
		Scan(&teams).Error

	if err != nil {
		return nil, 0, errors.Internal(err)
	}

	return teams, totalCount, nil
}

func (r *TeamRepo) GetById(ctx context.Context, teamId string) (*models.Team, error) {

	var team models.Team
	err := r.db.WithContext(ctx).
		Where("id = ? AND valid = ?", teamId, true).First(&team).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return &team, nil
}

func (r *TeamRepo) Create(ctx context.Context, team *models.Team) error {

	err := r.db.WithContext(ctx).Create(team).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TeamRepo) Update(ctx context.Context, team *models.Team) error {

	err := r.db.WithContext(ctx).Model(&models.Team{}).
		Where("id = ? AND valid = ?", team.ID, true).
		Updates(team).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TeamRepo) Delete(ctx context.Context, teamId string) error {

	err := r.db.WithContext(ctx).Model(&models.Team{}).
		Where("id = ? AND valid = ?", teamId, true).
		Updates(map[string]interface{}{
			"valid":     false,
			"updatedAt": utils.GetCurrentTime(),
		}).Error

	if err != nil {
		return err
	}

	return nil
}
