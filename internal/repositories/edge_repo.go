package repositories

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"gorm.io/gorm"
)

type IEdgeRepo interface {
	Get(ctx context.Context, formId string) ([]*models.Edge, error)
	Create(ctx context.Context, edge *models.Edge) (*models.Edge, error)
	Update(ctx context.Context, edgeId string, edge *models.Edge) error
	Delete(ctx context.Context, edgeId string) error
}

type EdgeRepo struct {
	db *gorm.DB
}

func NewEdgeRepo(db *gorm.DB) *EdgeRepo {
	return &EdgeRepo{
		db: db,
	}
}

func (r *EdgeRepo) Get(ctx context.Context, formId string) ([]*models.Edge, error) {

	var edges []*models.Edge
	err := r.db.WithContext(ctx).
		Where("form_id = ?", formId).
		Find(&edges).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return edges, nil
}

func (r *EdgeRepo) Create(ctx context.Context, edge *models.Edge) (*models.Edge, error) {

	err := r.db.WithContext(ctx).
		Model(&models.Edge{}).
		Create(edge).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return edge, nil
}

func (r *EdgeRepo) Update(ctx context.Context, edgeId string, edge *models.Edge) error {

	err := r.db.WithContext(ctx).
		Model(&models.Edge{}).
		Where("id = ?", edgeId).
		Updates(edge).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}

func (r *EdgeRepo) Delete(ctx context.Context, edgeId string) error {

	err := r.db.WithContext(ctx).
		Model(&models.Edge{}).
		Where("id = ?", edgeId).
		Delete(&models.Edge{}).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}
