package services

import (
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/repositories"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

type QuestionService struct {
	questionRepo *repositories.QuestionRepo
}

func NewQuestionService(questionRepo *repositories.QuestionRepo) *QuestionService {
	return &QuestionService{
		questionRepo: questionRepo,
	}
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
