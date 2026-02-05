package repositories

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"gorm.io/gorm"
)

type IFormVersionRepo interface {
	Create(ctx context.Context, formVersion *models.FormVersion) (*models.FormVersion, error)
	GetByFormID(ctx context.Context, formId string) ([]*models.FormVersion, error)
	GetByID(ctx context.Context, versionId string) (*models.FormVersion, error)
	GetLatestVersion(ctx context.Context, formId string) (*models.FormVersion, error)
}

type FormVersionRepo struct {
	db *gorm.DB
}

func NewFormVersionRepo(db *gorm.DB) *FormVersionRepo {
	return &FormVersionRepo{
		db: db,
	}
}

func (r *FormVersionRepo) Create(ctx context.Context, formVersion *models.FormVersion) (*models.FormVersion, error) {
	err := r.db.WithContext(ctx).
		Model(&models.FormVersion{}).
		Create(formVersion).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return formVersion, nil
}

func (r *FormVersionRepo) GetByFormID(ctx context.Context, formId string) ([]*models.FormVersion, error) {
	var versions []*models.FormVersion
	err := r.db.WithContext(ctx).
		Where(`"formId" = ?`, formId).
		Order(`"versionNumber" DESC`).
		Find(&versions).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return versions, nil
}

func (r *FormVersionRepo) GetByID(ctx context.Context, versionId string) (*models.FormVersion, error) {
	var version models.FormVersion
	err := r.db.WithContext(ctx).
		Where("id = ?", versionId).
		First(&version).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Internal(err)
	}
	return &version, nil
}

func (r *FormVersionRepo) GetLatestVersion(ctx context.Context, formId string) (*models.FormVersion, error) {
	var version models.FormVersion
	err := r.db.WithContext(ctx).
		Where(`"formId" = ?`, formId).
		Order(`"versionNumber" DESC`).
		First(&version).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Internal(err)
	}
	return &version, nil
}
