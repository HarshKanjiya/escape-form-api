package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type QuestionController struct {
}

func NewQuestionController(*services.QuestionService, *config.Config) *QuestionController {
	return &QuestionController{}
}

func (pc *QuestionController) Get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "QuestionController Get method called",
	})
}

func (pc *QuestionController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "QuestionController Create method called",
	})
}

func (pc *QuestionController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "QuestionController Update method called",
	})
}

func (pc *QuestionController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "QuestionController Delete method called",
	})
}
