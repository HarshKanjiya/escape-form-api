package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

type IQuestionService interface {
}

type QuestionService struct {
	questionRepo repositories.IQuestionRepo
}

func NewQuestionService(questionRepo repositories.IQuestionRepo) *QuestionService {
	return &QuestionService{
		questionRepo: questionRepo,
	}
}

func (s *QuestionService) Get(ctx *fiber.Ctx, formId string) ([]*models.Question, error) {
	return s.questionRepo.Get(ctx, formId)
}

func (s *QuestionService) Create(ctx *fiber.Ctx, question *types.QuestionDto) (*models.Question, error) {
	return s.questionRepo.Create(ctx, question)
}

func (s *QuestionService) Update(ctx *fiber.Ctx, question *types.QuestionDto) (*models.Question, error) {
	return s.questionRepo.Update(ctx, question)
}

func (s *QuestionService) GetOptions(ctx *fiber.Ctx, questionId string) ([]*models.QuestionOption, error) {
	return s.questionRepo.GetOptions(ctx, questionId)
}

func (s *QuestionService) CreateOption(ctx *fiber.Ctx, option *types.QuestionOptionDto) (*models.QuestionOption, error) {
	return s.questionRepo.CreateOption(ctx, option)
}

func (s *QuestionService) UpdateOption(ctx *fiber.Ctx, option *types.QuestionOptionDto) (*models.QuestionOption, error) {
	return s.questionRepo.UpdateOption(ctx, option)
}

func (s *QuestionService) DeleteOption(ctx *fiber.Ctx, optionId string) error {
	return s.questionRepo.DeleteOption(ctx, optionId)
}
