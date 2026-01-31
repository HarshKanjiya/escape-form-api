package services

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
)

type IQuestionService interface {
	GetQuestions(ctx context.Context, userId string, formId string) ([]*models.Question, error)
	CreateQuestion(ctx context.Context, userId string, formId string, question *types.QuestionDto) (*models.Question, error)
	UpdateQuestion(ctx context.Context, userId string, formId string, question *types.QuestionDto) (*models.Question, error)
	DeleteQuestion(ctx context.Context, userId string, formId string, questionId string) error

	GetOptions(ctx context.Context, userId string, formId string, questionId string) ([]*models.QuestionOption, error)
	CreateOption(ctx context.Context, userId string, formId string, option *types.QuestionOptionDto) (*models.QuestionOption, error)
	UpdateOption(ctx context.Context, userId string, formId string, option *types.QuestionOptionDto) error
	DeleteOption(ctx context.Context, userId string, formId string, optionId string) error
}

type QuestionService struct {
	questionRepo repositories.IQuestionRepo
	formRepo     repositories.IFormRepo
}

func NewQuestionService(questionRepo repositories.IQuestionRepo, formRepo repositories.IFormRepo) *QuestionService {
	return &QuestionService{
		questionRepo: questionRepo,
		formRepo:     formRepo,
	}
}

func (s *QuestionService) GetQuestions(ctx context.Context, userId string, formId string) ([]*models.Question, error) {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return nil, err
	}

	if form == nil {
		return nil, errors.NotFound("Form")
	}

	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return nil, errors.Unauthorized("")
	}

	questions, err := s.questionRepo.GetQuestions(ctx, formId)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (s *QuestionService) CreateQuestion(ctx context.Context, userId string, formId string, question *types.QuestionDto) (*models.Question, error) {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return nil, err
	}

	if form == nil {
		return nil, errors.NotFound("Form")
	}

	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return nil, errors.Unauthorized("")
	}

	queModel := &models.Question{
		ID:          utils.GenerateUUID(),
		FormID:      formId,
		Title:       question.Title,
		Required:    question.Required,
		Type:        question.Type,
		Metadata:    question.Metadata,
		PosX:        question.PosX,
		PosY:        question.PosY,
		Placeholder: question.Placeholder,
		Description: question.Description,
		SortOrder:   question.SortOrder,
	}

	createdQuestion, err := s.questionRepo.CreateQuestion(ctx, queModel)
	if err != nil {
		return nil, err
	}
	return createdQuestion, nil
}

func (s *QuestionService) UpdateQuestion(ctx context.Context, userId string, formId string, question *types.QuestionDto) (*models.Question, error) {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return nil, err
	}

	if form == nil {
		return nil, errors.NotFound("Form")
	}

	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return nil, errors.Unauthorized("")
	}

}

func (s *QuestionService) DeleteQuestion(ctx context.Context, userId string, formId string, questionId string) error {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return err
	}

	if form == nil {
		return errors.NotFound("Form")
	}

	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return errors.Unauthorized("")
	}

}

func (s *QuestionService) GetOptions(ctx context.Context, userId string, formId string, questionId string) ([]*models.QuestionOption, error) {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return nil, err
	}

	if form == nil {
		return nil, errors.NotFound("Form")
	}

	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return nil, errors.Unauthorized("")
	}
}

func (s *QuestionService) CreateOption(ctx context.Context, userId string, formId string, option *types.QuestionOptionDto) (*models.QuestionOption, error) {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return nil, err
	}

	if form == nil {
		return nil, errors.NotFound("Form")
	}

	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return nil, errors.Unauthorized("")
	}
}

func (s *QuestionService) UpdateOption(ctx context.Context, userId string, formId string, option *types.QuestionOptionDto) error {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return err
	}

	if form == nil {
		return errors.NotFound("Form")
	}

	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return errors.Unauthorized("")
	}
}

func (s *QuestionService) DeleteOption(ctx context.Context, userId string, formId string, optionId string) error {

	form, err := s.formRepo.GetByIdWithTeam(ctx, formId)
	if err != nil {
		return err
	}

	if form == nil {
		return errors.NotFound("Form")
	}

	if form.Team.OwnerID == nil || *form.Team.OwnerID != userId {
		return errors.Unauthorized("")
	}
}
