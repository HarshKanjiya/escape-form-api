package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type DashController struct {
	validator   *validator.Validate
	dashService services.IDashService
}

func NewDashController(service services.IDashService) *DashController {
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
// @Success 200 {object} types.FormAnalytics
// @Router /dashboard/{formId}/analytics [get]
func (pc *DashController) GetAnalytics(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId", "")
	if formId == "" {
		return errors.BadRequest("Form ID is required")
	}

	analytics, err := pc.dashService.GetAnalytics(c.Context(), userId, formId)
	if err != nil {
		return err
	}

	return utils.Success(c, analytics, "Analytics fetched successfully")
}

// @Summary Get form questions
// @Description Retrieve questions for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {array} models.Question
// @Router /dashboard/{formId}/questions [get]
func (pc *DashController) GetQuestions(c *fiber.Ctx) error {

	formId := c.Params("formId", "")
	if formId == "" {
		return errors.BadRequest("Form ID is required")
	}

	questions, err := pc.dashService.GetQuestions(c.Context(), formId)
	if err != nil {
		return err
	}

	return utils.Success(c, questions, "Questions fetched successfully")
}

// @Summary Get form responses
// @Description Retrieve responses for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {array} models.Response
// @Router /dashboard/{formId}/responses [get]
func (pc *DashController) GetResponses(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId", "")
	if formId == "" {
		return errors.BadRequest("Form ID is required")
	}

	responses, err := pc.dashService.GetResponses(c.Context(), userId, formId)
	if err != nil {
		return err
	}

	return utils.Success(c, responses, "Responses fetched successfully")
}

// @Summary Get form passwords
// @Description Retrieve passwords for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {array} models.ActivePassword
// @Router /dashboard/{formId}/passwords [get]
func (pc *DashController) GetPasswords(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId", "")
	if formId == "" {
		return errors.BadRequest("Form ID is required")
	}

	passwords, err := pc.dashService.GetPasswords(c.Context(), userId, formId)
	if err != nil {
		return err
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
// @Param body body types.PasswordRequest true "Password update data"
// @Success 200 {object} models.ActivePassword
// @Router /dashboard/{formId}/passwords/{passwordId} [patch]
func (pc *DashController) UpdatePasswords(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId", "")
	passwordId := c.Params("passwordId", "")
	if formId == "" || passwordId == "" {
		return errors.BadRequest("Form ID and Password ID are required")
	}

	var body types.PasswordRequest
	if err := c.BodyParser(&body); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	err := pc.dashService.UpdatePassword(c.Context(), userId, formId, passwordId, body)
	if err != nil {
		return err
	}

	return utils.Success(c, nil, "Password updated successfully")
}

// @Summary Create form passwords
// @Description Create passwords for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param body body types.PasswordRequest true "Password creation data"
// @Success 200 {object} models.ActivePassword
// @Router /dashboard/{formId}/passwords [post]
func (pc *DashController) CreatePasswords(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId", "")
	if formId == "" {
		return errors.BadRequest("projectId is required")
	}

	var body types.PasswordRequest
	if err := c.BodyParser(&body); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	password, err := pc.dashService.CreatePassword(c.Context(), userId, formId, body)
	if err != nil {
		return err
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

	formId := c.Params("formId", "")
	passwordId := c.Params("passwordId", "")
	if formId == "" || passwordId == "" {
		return errors.BadRequest("Form ID and Password ID are required")
	}

	err := pc.dashService.DeletePassword(c.Context(), userId, formId, passwordId)
	if err != nil {
		return err
	}

	return utils.Success(c, nil, "Password deleted successfully")
}

// @Summary Update form security
// @Description Update security settings for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param body body types.UpdateSecurityRequest true "Security settings"
// @Success 200 {object} interface{}
// @Router /dashboard/{formId}/security [patch]
func (pc *DashController) UpdateSecurity(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId", "")
	if formId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Form ID is required"})
	}

	var body types.UpdateSecurityRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := pc.dashService.UpdateSecurity(c.Context(), userId, formId, &body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.Success(c, nil, "Security updated successfully")
}

// @Summary Update form settings
// @Description Update settings for a form
// @Tags dashboard
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param body body types.UpdateSettingsRequest true "Form settings"
// @Success 200 {object} interface{}
// @Router /dashboard/{formId}/settings [patch]
func (pc *DashController) UpdateSettings(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId", "")
	if formId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Form ID is required"})
	}

	var body types.UpdateSettingsRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := pc.dashService.UpdateSettings(c.Context(), userId, formId, &body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.Success(c, nil, "Settings updated successfully")
}
