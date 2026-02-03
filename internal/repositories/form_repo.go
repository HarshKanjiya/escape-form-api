package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IFormRepo interface {
	Get(ctx context.Context, pagination *types.PaginationQuery, projectId string, status *models.FormStatus) ([]*models.Form, int64, error)
	GetById(ctx context.Context, formId string) (*models.Form, error)
	GetWithTeam(ctx context.Context, formId string) (*models.Form, error)
	Create(ctx context.Context, form *models.Form) (*models.Form, error)
	Update(ctx context.Context, formId string, updates map[string]interface{}) error
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
		Where(&models.Form{
			ProjectID: projectId,
			Valid:     true,
		})

	if pagination.Search != "" {
		baseQuery = baseQuery.Where(
			clause.Like{
				Column: clause.Column{Name: "name"},
				Value:  "%" + pagination.Search + "%",
			},
		)
	}

	if status != nil {
		baseQuery = baseQuery.Where(&models.Form{
			Status: status,
		})
	}

	// total count
	if err := baseQuery.Count(&totalCount).Error; err != nil {
		return nil, 0, errors.Internal(err)
	}

	err := baseQuery.
		Select(`forms.*, (SELECT COUNT(*) FROM responses WHERE responses."formId" = forms.id) as "responseCount"`).
		Order(`forms."createdAt" DESC`).
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
		Preload("Questions.Options").
		Preload("Edges").
		Where(&models.Form{
			ID:    formId,
			Valid: true,
		}).
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

func (r *FormRepo) Update(ctx context.Context, formId string, updates map[string]interface{}) error {
	newUpdates := make(map[string]interface{})
	for k, v := range updates {
		if k == "theme" || k == "metadata" {
			jsonBytes, err := json.Marshal(v)
			if err != nil {
				return errors.Internal(err)
			}
			newUpdates[k] = string(jsonBytes)
		} else {
			newUpdates[k] = v
		}
	}
	err := r.db.WithContext(ctx).
		Model(&models.Form{}).
		Where(&models.Form{
			ID:    formId,
			Valid: true,
		}).
		Updates(newUpdates).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}

func (r *FormRepo) UpdateStatus(ctx context.Context, formId string, status models.FormStatus) error {

	err := r.db.WithContext(ctx).
		Model(&models.Form{}).
		Where(&models.Form{
			ID:    formId,
			Valid: true,
		}).
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
		Joins(`
        JOIN teams 
        ON teams.id = forms."teamId"
        AND teams.valid = true
    `).
		Where(`
        forms.id = ?
        AND forms.valid = true
    `, formId).
		Preload("Team").
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
		Where(`id IN ? AND "formId" = ?`, ids, formId).
		Update(`"sortOrder"`, gorm.Expr(caseStr)).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}
