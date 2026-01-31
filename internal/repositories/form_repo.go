package repositories

import (
	"context"
	"fmt"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"gorm.io/gorm"
)

type IFormRepo interface {
	Get(ctx context.Context, pagination *types.PaginationQuery, projectId string, status *models.FormStatus) ([]*models.Form, int64, error)
	GetById(ctx context.Context, formId string) (*models.Form, error)
	GetWithTeam(ctx context.Context, formId string) (*models.Form, error)
	Create(ctx context.Context, form *models.Form) (*models.Form, error)
	Update(ctx context.Context, formId string, updates *map[string]interface{}) error
	UpdateStatus(ctx context.Context, formId string, status models.FormStatus) error
	Delete(ctx context.Context, formId string) error

	UpdateQuestionSequence(ctx context.Context, formId string, sequence []*types.SequenceItem) error
}

type FormRepo struct {
	db *gorm.DB
}

func NewFormRepo(db *gorm.DB) *FormRepo {
	return &FormRepo{
		db: db,
	}
}

func (r *FormRepo) Get(ctx context.Context, pagination *types.PaginationQuery, projectId string, status *models.FormStatus) ([]*models.Form, int64, error) {

	var forms []*models.Form
	var totalCount int64

	baseQuery := r.db.
		WithContext(ctx).
		Model(&models.Form{}).
		Where("project_id = ? AND valid = ?", projectId, true)

	if pagination.Search != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+pagination.Search+"%")
	}

	if status != nil {
		baseQuery = baseQuery.Where("status = ?", *status)
	}

	// total count
	if err := baseQuery.Count(&totalCount).Error; err != nil {
		return nil, 0, errors.Internal(err)
	}

	err := baseQuery.
		Limit(pagination.Limit).
		Offset((pagination.Page - 1) * pagination.Limit).
		Find(&forms).Error
	if err != nil {
		return nil, 0, errors.Internal(err)
	}
	return forms, totalCount, nil
}

func (r *FormRepo) GetById(ctx context.Context, formId string) (*models.Form, error) {

	var form *models.Form
	err := r.db.WithContext(ctx).
		Model(&models.Form{}).
		Preload("Questions").
		Preload("Edges").
		Where("id = ? AND valid = ?", formId, true).
		First(&form).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return form, nil
}

func (r *FormRepo) Create(ctx context.Context, form *models.Form) (*models.Form, error) {
	err := r.db.WithContext(ctx).
		Model(&models.Form{}).
		Create(form).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return r.GetById(ctx, form.ID)
}

func (r *FormRepo) Update(ctx context.Context, formId string, updates *map[string]interface{}) error {
	err := r.db.WithContext(ctx).
		Model(&models.Form{}).
		Where("id = ? AND valid = ?", formId, true).
		Updates(*updates).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}

func (r *FormRepo) UpdateStatus(ctx context.Context, formId string, status models.FormStatus) error {

	err := r.db.WithContext(ctx).
		Model(&models.Form{}).
		Where("id = ? AND valid = ?", formId, true).
		Update("status", status).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}

func (r *FormRepo) GetWithTeam(ctx context.Context, formId string) (*models.Form, error) {

	var form *models.Form
	err := r.db.WithContext(ctx).
		Model(&models.Form{}).
		Preload("Team").
		Where("forms.id = ? AND teams.valid = ? AND forms.valid = ?", formId, true, true).
		First(&form).Error

	if err != nil {
		return nil, errors.Internal(err)
	}
	return form, nil
}

func (r *FormRepo) Delete(ctx context.Context, formId string) error {

	err := r.db.WithContext(ctx).
		Model(&models.Form{}).
		Where("id = ?", formId).
		Updates(map[string]interface{}{
			"valid":     false,
			"updatedAt": utils.GetCurrentTime(),
		}).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}

func (r *FormRepo) UpdateQuestionSequence(ctx context.Context, formId string, sequence []*types.SequenceItem) error {
	if len(sequence) == 0 {
		return nil
	}

	ids := make([]string, len(sequence))
	for i, item := range sequence {
		ids[i] = item.ID
	}

	caseStr := "CASE id "
	for _, item := range sequence {
		caseStr += fmt.Sprintf("WHEN '%s' THEN %d ", item.ID, item.NewOrder)
	}
	caseStr += "END"

	err := r.db.WithContext(ctx).
		Model(&models.Question{}).
		Where("id IN ? AND form_id = ?", ids, formId).
		Update("sort_order", gorm.Expr(caseStr)).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}
