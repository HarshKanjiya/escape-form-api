package services

import (
	"context"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
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
}

func NewQuestionService(questionRepo repositories.IQuestionRepo) *QuestionService {
	return &QuestionService{
		questionRepo: questionRepo,
	}
}

func (s *QuestionService) GetQuestions(ctx context.Context, userId string, formId string) ([]*models.Question, error) {

}

func (s *QuestionService) CreateQuestion(ctx context.Context, userId string, formId string, question *types.QuestionDto) (*models.Question, error) {

}

func (s *QuestionService) UpdateQuestion(ctx context.Context, userId string, formId string, question *types.QuestionDto) (*models.Question, error) {

}

func (s *QuestionService) DeleteQuestion(ctx context.Context, userId string, formId string, questionId string) error {

}

func (s *QuestionService) GetOptions(ctx context.Context, userId string, formId string, questionId string) ([]*models.QuestionOption, error) {

}

func (s *QuestionService) CreateOption(ctx context.Context, userId string, formId string, option *types.QuestionOptionDto) (*models.QuestionOption, error) {

}

func (s *QuestionService) UpdateOption(ctx context.Context, userId string, formId string, option *types.QuestionOptionDto) error {

}

func (s *QuestionService) DeleteOption(ctx context.Context, userId string, formId string, optionId string) error {

}
