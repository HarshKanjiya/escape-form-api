package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type DashController struct {
	validator   *validator.Validate
	dashService *services.DashService
}

func NewDashController(service *services.DashService) *DashController {
	return &DashController{
		validator:   validator.New(),
		dashService: service,
	}
}

// @Summary Get form analytics
// @Description Retrieve analytics for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Router /dashboard/{formId}/analytics [get]
func (pc *DashController) GetAnalytics(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	if formId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Form ID is required"})
	}

	analytics, err := pc.dashService.GetAnalytics(c, formId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.Success(c, analytics, "Analytics fetched successfully")
}

// @Summary Get form questions
// @Description Retrieve questions for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Router /dashboard/{formId}/questions [get]
func (pc *DashController) GetQuestions(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	if formId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Form ID is required"})
	}

	questions, err := pc.dashService.GetQuestions(c, formId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.Success(c, questions, "Questions fetched successfully")
}

// @Summary Get form responses
// @Description Retrieve responses for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Router /dashboard/{formId}/responses [get]
func (pc *DashController) GetResponses(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	if formId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Form ID is required"})
	}

	responses, err := pc.dashService.GetResponses(c, formId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.Success(c, responses, "Responses fetched successfully")
}

// @Summary Get form passwords
// @Description Retrieve passwords for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Router /dashboard/{formId}/passwords [get]
func (pc *DashController) GetPasswords(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	if formId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Form ID is required"})
	}

	passwords, err := pc.dashService.GetPasswords(c, formId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.Success(c, passwords, "Passwords fetched successfully")
}

// @Summary Update form passwords
// @Description Update passwords for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param passwordId path string true "Password ID"
// @Success 200 {object} map[string]interface{}
// @Router /dashboard/{formId}/passwords/{passwordId} [patch]
func (pc *DashController) UpdatePasswords(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	passwordId := c.Params("passwordId")
	if formId == "" || passwordId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Form ID and Password ID are required"})
	}

	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	password, err := pc.dashService.UpdatePassword(c, formId, passwordId, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.Success(c, password, "Password updated successfully")
}

// @Summary Create form passwords
// @Description Create passwords for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Router /dashboard/{formId}/passwords [post]
func (pc *DashController) CreatePasswords(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	if formId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Form ID is required"})
	}

	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	password, err := pc.dashService.CreatePassword(c, formId, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.Success(c, password, "Password created successfully")
}

// @Summary Delete form passwords
// @Description Delete passwords for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param passwordId path string true "Password ID"
// @Success 200 {object} map[string]interface{}
// @Router /dashboard/{formId}/passwords/{passwordId} [delete]
func (pc *DashController) DeletePasswords(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	passwordId := c.Params("passwordId")
	if formId == "" || passwordId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Form ID and Password ID are required"})
	}

	err := pc.dashService.DeletePassword(c, formId, passwordId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.Success(c, nil, "Password deleted successfully")
}

func (pc *DashController) Create(c *fiber.Ctx) error {
	return utils.Success(c, nil, "DashController Create method called")
}

func (pc *DashController) Update(c *fiber.Ctx) error {
	return utils.Success(c, nil, "DashController Update method called")
}

func (pc *DashController) Delete(c *fiber.Ctx) error {
	return utils.Success(c, nil, "DashController Delete method called")
}

// @Summary Update form security
// @Description Update security settings for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Router /dashboard/{formId}/security [patch]
func (pc *DashController) UpdateSecurity(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	if formId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Form ID is required"})
	}

	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	form, err := pc.dashService.UpdateSecurity(c, formId, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.Success(c, form, "Security updated successfully")
}

// @Summary Update form settings
// @Description Update settings for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Router /dashboard/{formId}/settings [patch]
func (pc *DashController) UpdateSettings(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	if formId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Form ID is required"})
	}

	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	form, err := pc.dashService.UpdateSettings(c, formId, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.Success(c, form, "Settings updated successfully")
}
