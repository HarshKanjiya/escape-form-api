package repositories

import (
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/query"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionRepo struct {
	q *query.Query
}

func NewQuestionRepo(db *gorm.DB) *QuestionRepo {
	return &QuestionRepo{
		q: query.Use(db),
	}
}

func (r *QuestionRepo) GetOptions(ctx *fiber.Ctx, questionId string) ([]*models.QuestionOption, error) {
	options, err := r.q.WithContext(ctx.Context()).QuestionOption.Where(r.q.QuestionOption.QuestionID.Eq(questionId)).Find()
	if err != nil {
		return nil, err
	}
	return options, nil
}

func (r *QuestionRepo) CreateOption(ctx *fiber.Ctx, option *types.QuestionOptionDto) (*models.QuestionOption, error) {
	optionModel := &models.QuestionOption{
		ID:         uuid.New().String(),
		QuestionID: option.QuestionID,
		Label:      option.Label,
		Value:      option.Value,
		SortOrder:  option.SortOrder,
	}
	if err := r.q.QuestionOption.Create(optionModel); err != nil {
		return nil, err
	}
	return optionModel, nil
}

func (r *QuestionRepo) UpdateOption(ctx *fiber.Ctx, option *types.QuestionOptionDto) (*models.QuestionOption, error) {
	optionModel := &models.QuestionOption{
		ID:         option.ID,
		QuestionID: option.QuestionID,
		Label:      option.Label,
		Value:      option.Value,
		SortOrder:  option.SortOrder,
	}
	_, err := r.q.WithContext(ctx.Context()).
		QuestionOption.Where(r.q.QuestionOption.ID.Eq(option.ID)).
		Updates(optionModel)
	if err != nil {
		return nil, err
	}
	return optionModel, nil
}

func (r *QuestionRepo) DeleteOption(ctx *fiber.Ctx, optionId string) error {
	_, err := r.q.WithContext(ctx.Context()).
		QuestionOption.Where(r.q.QuestionOption.ID.Eq(optionId)).
		Delete()
	if err != nil {
		return err
	}
	return nil
}
