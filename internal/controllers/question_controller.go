package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type QuestionController struct {
	validator       *validator.Validate
	questionService services.IQuestionService
}

func NewQuestionController(service services.IQuestionService) *QuestionController {
	return &QuestionController{
		validator:       validator.New(),
		questionService: service,
	}
}

// @Summary Get all questions
// @Description Retrieve a list of questions
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {array} models.Question
// @Router /forms/{formId}/questions [get]
func (pc *QuestionController) GetQuestions(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")

	if formId == "" {
		return errors.BadRequest("Form ID is required")
	}
	questions, err := pc.questionService.GetQuestions(c.Context(), userId, formId)
	if err != nil {
		return err
	}
	return utils.Success(c, questions, "Questions fetched successfully")
}

// @Summary Create a new question
// @Description Create a new question
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param body body types.QuestionRequest true "Question data"
// @Success 200 {object} models.Question
// @Router /forms/{formId}/questions [post]
func (pc *QuestionController) CreateQuestion(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")

	if formId == "" {
		return errors.BadRequest("Form ID is required")
	}

	var questionDto types.QuestionRequest
	if err := c.BodyParser(&questionDto); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	if err := pc.validator.Struct(&questionDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}
	question, err := pc.questionService.CreateQuestion(c.Context(), userId, formId, &questionDto)
	if err != nil {
		return err
	}
	return utils.Created(c, question, "Question created successfully")
}

// @Summary Update a question
// @Description Update an existing question by ID
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param questionId path string true "Question ID"
// @Param body body types.QuestionRequest true "Question data"
// @Success 200 {object} types.ResponseObj
// @Router /forms/{formId}/questions/{questionId} [patch]
func (pc *QuestionController) UpdateQuestion(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	questionId := c.Params("questionId")
	formId := c.Params("formId")

	if questionId == "" || formId == "" {
		return errors.BadRequest("Question ID and Form ID are required")
	}

	var questionDto map[string]interface{}
	if err := c.BodyParser(&questionDto); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	err := pc.questionService.UpdateQuestion(c.Context(), userId, formId, questionId, &questionDto)
	if err != nil {
		return err
	}
	return utils.Success(c, nil, "Question updated successfully")
}

// @Summary Delete a question
// @Description Delete a question by ID
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param questionId path string true "Question ID"
// @Success 200 {object} types.ResponseObj
// @Router /forms/{formId}/questions/{questionId} [delete]
func (pc *QuestionController) DeleteQuestion(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	questionId := c.Params("questionId")
	formId := c.Params("formId")

	if questionId == "" || formId == "" {
		return errors.BadRequest("Question ID and Form ID are required")
	}

	err := pc.questionService.DeleteQuestion(c.Context(), userId, formId, questionId)
	if err != nil {
		return err
	}
	return utils.Success(c, nil, "Question deleted successfully")
}

// @Summary Get question options
// @Description Retrieve options for a question
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param questionId path string true "Question ID"
// @Success 200 {array} models.QuestionOption
// @Router /forms/{formId}/questions/{questionId}/options [get]
func (pc *QuestionController) GetOptions(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	questionId := c.Params("questionId")
	formId := c.Params("formId")

	if questionId == "" || formId == "" {
		return errors.BadRequest("Question ID and Form ID are required")
	}

	options, err := pc.questionService.GetOptions(c.Context(), userId, formId, questionId)
	if err != nil {
		return err
	}

	return utils.Success(c, options, "Options fetched successfully")
}

// @Summary Create a new question option
// @Description Create a new option for a question
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param questionId path string true "Question ID"
// @Param body body types.QuestionOptionRequest true "Option data"
// @Success 201 {object} types.ResponseObj
// @Router /forms/{formId}/questions/{questionId}/options [post]
func (pc *QuestionController) CreateOption(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	questionId := c.Params("questionId")
	formId := c.Params("formId")

	if questionId == "" || formId == "" {
		return errors.BadRequest("Question ID and Form ID are required")
	}

	var optionDto types.QuestionOptionRequest
	if err := c.BodyParser(&optionDto); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	if err := pc.validator.Struct(&optionDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}

	option, err := pc.questionService.CreateOption(c.Context(), userId, formId, questionId, &optionDto)
	if err != nil {
		return err
	}

	return utils.Created(c, option, "Option created successfully")
}

// @Summary Update a question option
// @Description Update an existing option by ID
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param questionId path string true "Question ID"
// @Param optionId path string true "Option ID"
// @Param body body types.QuestionOptionRequest true "Option data"
// @Success 200 {object} map[string]interface{}
// @Router /forms/{formId}/questions/{questionId}/options/{optionId} [patch]
func (pc *QuestionController) UpdateOption(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	questionId := c.Params("questionId")
	formId := c.Params("formId")
	optionId := c.Params("optionId")

	if questionId == "" || formId == "" || optionId == "" {
		return errors.BadRequest("Question ID, Form ID, and Option ID are required")
	}

	var optionDto types.QuestionOptionRequest
	if err := c.BodyParser(&optionDto); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	if err := pc.validator.Struct(&optionDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}
	err := pc.questionService.UpdateOption(c.Context(), userId, formId, questionId, optionId, &optionDto)
	if err != nil {
		return err
	}
	return utils.Success(c, nil, "Option updated successfully")
}

// @Summary Delete a question option
// @Description Delete an option by ID
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param questionId path string true "Question ID"
// @Param optionId path string true "Option ID"
// @Success 200 {object} map[string]interface{}
// @Router /forms/{formId}/questions/{questionId}/options/{optionId} [delete]
func (pc *QuestionController) DeleteOption(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	questionId := c.Params("questionId")
	formId := c.Params("formId")
	optionId := c.Params("optionId")

	if questionId == "" || formId == "" || optionId == "" {
		return errors.BadRequest("Question ID, Form ID, and Option ID are required")
	}

	err := pc.questionService.DeleteOption(c.Context(), userId, formId, optionId)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to delete option")
	}
	return utils.NoContent(c)
}
