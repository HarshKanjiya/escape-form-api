package repositories

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"gorm.io/gorm"
)

type IQuestionRepo interface {
	GetQuestions(ctx context.Context, formId string) ([]*models.Question, error)
	CreateQuestion(ctx context.Context, question *models.Question) (*models.Question, error)
	UpdateQuestion(ctx context.Context, questionId string, question *map[string]interface{}) error
	DeleteQuestion(ctx context.Context, questionId string) error

	GetOptions(ctx context.Context, questionId string) ([]*models.QuestionOption, error)
	CreateOption(ctx context.Context, option *models.QuestionOption) (*models.QuestionOption, error)
	UpdateOption(ctx context.Context, optionId string, option *map[string]interface{}) error
	DeleteOption(ctx context.Context, optionId string) error
}

type QuestionRepo struct {
	db *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) *QuestionRepo {
	return &QuestionRepo{
		db: db,
	}
}

func (r *QuestionRepo) GetQuestions(ctx context.Context, formId string) ([]*models.Question, error) {

	var questions []*models.Question
	err := r.db.WithContext(ctx).
		Model(&models.Question{}).
		Where(`"formId" = ?`, formId).
		Preload("Options").
		Find(&questions).Error
	if err != nil {
		return nil, errors.Internal(err)
	}

	return questions, err
}

func (r *QuestionRepo) CreateQuestion(ctx context.Context, question *models.Question) (*models.Question, error) {

	err := r.db.WithContext(ctx).
		Model(&models.Question{}).
		Create(question).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return question, nil
}

func (r *QuestionRepo) UpdateQuestion(ctx context.Context, questionId string, question *map[string]interface{}) error {

	err := r.db.WithContext(ctx).
		Model(&models.Question{}).
		Where("id = ?", questionId).
		Updates(*question).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}

func (r *QuestionRepo) DeleteQuestion(ctx context.Context, questionId string) error {

	err := r.db.WithContext(ctx).
		Model(&models.Question{}).
		Where("id = ?", questionId).
		Delete(&models.Question{}).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}

func (r *QuestionRepo) GetOptions(ctx context.Context, questionId string) ([]*models.QuestionOption, error) {

	var options []*models.QuestionOption
	err := r.db.WithContext(ctx).
		Model(&models.QuestionOption{}).
		Where(`"questionId" = ?`, questionId).
		Find(&options).Error
	if err != nil {
		return nil, errors.Internal(err)
	}

	return options, err
}

func (r *QuestionRepo) CreateOption(ctx context.Context, option *models.QuestionOption) (*models.QuestionOption, error) {

	err := r.db.WithContext(ctx).
		Model(&models.QuestionOption{}).
		Create(option).Error
	if err != nil {
		return nil, errors.Internal(err)
	}
	return option, nil
}

func (r *QuestionRepo) UpdateOption(ctx context.Context, optionId string, option *map[string]interface{}) error {

	err := r.db.WithContext(ctx).
		Model(&models.QuestionOption{}).
		Where("id = ?", optionId).
		Updates(*option).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}

func (r *QuestionRepo) DeleteOption(ctx context.Context, optionId string) error {

	err := r.db.WithContext(ctx).
		Model(&models.QuestionOption{}).
		Where("id = ?", optionId).
		Delete(&models.QuestionOption{}).Error
	if err != nil {
		return errors.Internal(err)
	}
	return nil
}
