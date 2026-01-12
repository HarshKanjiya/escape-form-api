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

// @Summary Get all questions
// @Description Retrieve a list of questions
// @Tags questions
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /questions [get]
func (pc *QuestionController) Get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "QuestionController Get method called",
	})
}

// @Summary Create a new question
// @Description Create a new question
// @Tags questions
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /questions [post]
func (pc *QuestionController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "QuestionController Create method called",
	})
}

// @Summary Update a question
// @Description Update an existing question by ID
// @Tags questions
// @Accept json
// @Produce json
// @Param id path string true "Question ID"
// @Success 200 {object} map[string]interface{}
// @Router /questions/{id} [patch]
func (pc *QuestionController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "QuestionController Update method called",
	})
}

// @Summary Delete a question
// @Description Delete a question by ID
// @Tags questions
// @Accept json
// @Produce json
// @Param id path string true "Question ID"
// @Success 200 {object} map[string]interface{}
// @Router /questions/{id} [delete]
func (pc *QuestionController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "QuestionController Delete method called",
	})
}
