package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type QuestionController struct {
	validator       *validator.Validate
	questionService *services.QuestionService
}

func NewQuestionController(service *services.QuestionService) *QuestionController {
	return &QuestionController{
		validator:       validator.New(),
		questionService: service,
	}
}

// @Summary Get all questions
// @Description Retrieve a list of questions
// @Tags questions
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /forms/{formId}/questions [get]
func (pc *QuestionController) Get(c *fiber.Ctx) error {

	formId := c.Params("formId")

	if formId == "" {
		return utils.BadRequest(c, "Form ID is required")
	}
	questions, err := pc.questionService.Get(c, formId)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to fetch questions")
	}
	return utils.Success(c, questions, "Questions fetched successfully")
}

// @Summary Create a new question
// @Description Create a new question
// @Tags questions
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /forms/{formId}/questions [post]
func (pc *QuestionController) Create(c *fiber.Ctx) error {
	formId := c.Params("formId")

	if formId == "" {
		return utils.BadRequest(c, "Form ID is required")
	}

	var questionDto types.QuestionDto
	if err := c.BodyParser(&questionDto); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}
	questionDto.FormID = formId

	if err := pc.validator.Struct(&questionDto); err != nil {
		return utils.BadRequest(c, "Validation failed: "+err.Error())
	}
	question, err := pc.questionService.Create(c, &questionDto)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to create question")
	}
	return utils.Created(c, question, "Question created successfully")
}

// @Summary Update a question
// @Description Update an existing question by ID
// @Tags questions
// @Accept json
// @Produce json
// @Param questionId path string true "Question ID"
// @Success 200 {object} map[string]interface{}
// @Router /questions/{questionId} [patch]
func (pc *QuestionController) Update(c *fiber.Ctx) error {
	questionId := c.Params("questionId")

	if questionId == "" {
		return utils.BadRequest(c, "Question ID is required")
	}

	var questionDto types.QuestionDto
	if err := c.BodyParser(&questionDto); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}
	questionDto.ID = questionId

	if err := pc.validator.Struct(&questionDto); err != nil {
		return utils.BadRequest(c, "Validation failed: "+err.Error())
	}
	question, err := pc.questionService.Update(c, &questionDto)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to update question")
	}
	return utils.Success(c, question, "Question updated successfully")
}

// @Summary Delete a question
// @Description Delete a question by ID
// @Tags questions
// @Accept json
// @Produce json
// @Param questionId path string true "Question ID"
// @Success 200 {object} map[string]interface{}
// @Router /questions/{id} [delete]
func (pc *QuestionController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "QuestionController Delete method called",
	})
}

// @Summary Get question options
// @Description Retrieve options for a question
// @Tags questions
// @Accept json
// @Produce json
// @Param questionId path string true "Question ID"
// @Success 200 {object} map[string]interface{}
// @Router /questions/{questionId}/options [get]
func (pc *QuestionController) GetOptions(c *fiber.Ctx) error {
	questionId := c.Params("questionId")
	options, err := pc.questionService.GetOptions(c, questionId)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to fetch options")
	}
	return utils.Success(c, options, "Options fetched successfully")
}

// @Summary Create a new question option
// @Description Create a new option for a question
// @Tags questions
// @Accept json
// @Produce json
// @Param questionId path string true "Question ID"
// @Success 200 {object} map[string]interface{}
// @Router /questions/{questionId}/options [post]
func (pc *QuestionController) CreateOption(c *fiber.Ctx) error {
	questionId := c.Params("questionId")
	var optionDto types.QuestionOptionDto
	if err := c.BodyParser(&optionDto); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}
	optionDto.QuestionID = questionId
	if err := pc.validator.Struct(&optionDto); err != nil {
		return utils.BadRequest(c, "Validation failed: "+err.Error())
	}
	option, err := pc.questionService.CreateOption(c, &optionDto)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to create option")
	}
	return utils.Created(c, option, "Option created successfully")
}

// @Summary Update a question option
// @Description Update an existing option by ID
// @Tags questions
// @Accept json
// @Produce json
// @Param optionId path string true "Option ID"
// @Success 200 {object} map[string]interface{}
// @Router /questions/options/{optionId} [patch]
func (pc *QuestionController) UpdateOption(c *fiber.Ctx) error {
	optionId := c.Params("optionId")
	var optionDto types.QuestionOptionDto
	if err := c.BodyParser(&optionDto); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}
	optionDto.ID = optionId
	if err := pc.validator.Struct(&optionDto); err != nil {
		return utils.BadRequest(c, "Validation failed: "+err.Error())
	}
	option, err := pc.questionService.UpdateOption(c, &optionDto)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to update option")
	}
	return utils.Success(c, option, "Option updated successfully")
}

// @Summary Delete a question option
// @Description Delete an option by ID
// @Tags questions
// @Accept json
// @Produce json
// @Param optionId path string true "Option ID"
// @Success 200 {object} map[string]interface{}
// @Router /questions/options/{optionId} [delete]
func (pc *QuestionController) DeleteOption(c *fiber.Ctx) error {
	optionId := c.Params("optionId")
	err := pc.questionService.DeleteOption(c, optionId)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to delete option")
	}
	return utils.NoContent(c)
}
