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
	CreateQuestion(ctx context.Context, userId string, formId string, question *types.QuestionRequest) (*models.Question, error)
	UpdateQuestion(ctx context.Context, userId string, formId string, questionId string, question *types.QuestionRequest) error
	DeleteQuestion(ctx context.Context, userId string, formId string, questionId string) error

	GetOptions(ctx context.Context, userId string, formId string, questionId string) ([]*models.QuestionOption, error)
	CreateOption(ctx context.Context, userId string, formId string, questionId string, option *types.QuestionOptionRequest) (*models.QuestionOption, error)
	UpdateOption(ctx context.Context, userId string, formId string, questionId string, optionId string, option *types.QuestionOptionRequest) error
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

func (s *QuestionService) CreateQuestion(ctx context.Context, userId string, formId string, question *types.QuestionRequest) (*models.Question, error) {

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

func (s *QuestionService) UpdateQuestion(ctx context.Context, userId string, formId string, questionId string, question *types.QuestionRequest) error {

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

	updates := make(map[string]interface{})
	if &question.Title != nil {
		updates["title"] = question.Title
	}
	if &question.Required != nil {
		updates["required"] = question.Required
	}
	if &question.Type != nil {
		updates["type"] = question.Type
	}
	if &question.Metadata != nil {
		updates["metadata"] = question.Metadata
	}
	if &question.PosX != nil {
		updates["pos_x"] = question.PosX
	}
	if &question.PosY != nil {
		updates["pos_y"] = question.PosY
	}
	if &question.Placeholder != nil {
		updates["placeholder"] = question.Placeholder
	}
	if &question.Description != nil {
		updates["description"] = question.Description
	}
	if &question.SortOrder != nil {
		updates["sort_order"] = question.SortOrder
	}

	err = s.questionRepo.UpdateQuestion(ctx, questionId, &updates)
	if err != nil {
		return err
	}
	return nil
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

	err = s.questionRepo.DeleteQuestion(ctx, questionId)
	if err != nil {
		return err
	}
	return nil
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
	options, err := s.questionRepo.GetOptions(ctx, questionId)
	if err != nil {
		return nil, err
	}
	return options, nil
}

func (s *QuestionService) CreateOption(ctx context.Context, userId string, formId string, questionId string, option *types.QuestionOptionRequest) (*models.QuestionOption, error) {

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
	optionModel := &models.QuestionOption{
		ID:         utils.GenerateUUID(),
		QuestionID: questionId,
		Label:      option.Label,
		Value:      option.Value,
		SortOrder:  option.SortOrder,
	}
	createdOption, err := s.questionRepo.CreateOption(ctx, optionModel)
	if err != nil {
		return nil, err
	}
	return createdOption, nil
}

func (s *QuestionService) UpdateOption(ctx context.Context, userId string, formId string, questionId string, optionId string, option *types.QuestionOptionRequest) error {

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

	updates := make(map[string]interface{})
	if &option.Label != nil {
		updates["label"] = option.Label
	}
	if &option.Value != nil {
		updates["value"] = option.Value
	}
	if &option.SortOrder != nil {
		updates["sort_order"] = option.SortOrder
	}

	err = s.questionRepo.UpdateOption(ctx, optionId, &updates)
	if err != nil {
		return err
	}
	return nil
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
	err = s.questionRepo.DeleteOption(ctx, optionId)
	if err != nil {
		return err
	}
	return nil
}
