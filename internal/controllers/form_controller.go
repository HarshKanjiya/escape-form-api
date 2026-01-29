package controllers

import (
	"log"

	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FormController struct {
	validator   *validator.Validate
	formService *services.FormService
}

func NewFormController(service *services.FormService) *FormController {
	return &FormController{
		validator:   validator.New(),
		formService: service,
	}
}

// @Summary Get all forms
// @Description Retrieve a list of forms
// @Tags forms
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /forms [get]
func (fc *FormController) Get(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	pagination := &types.PaginationQuery{
		Page:   c.QueryInt("page", 1),
		Limit:  c.QueryInt("limit", 10),
		Search: c.Query("search", ""),
		SortBy: c.Query("sortBy", ""),
		Order:  c.Query("order", ""),
	}
	projectId := c.Query("projectId", "")
	if projectId == "" {
		return utils.BadRequest(c, "projectId is required")
	}
	forms, total, err := fc.formService.Get(c.Context(), userId, pagination, true, projectId)
	if err != nil {
		return err
	}
	return utils.Success(c, forms, "Forms fetched successfully", total)
}

// @Summary Create a new form
// @Description Create a new form
// @Tags forms
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /forms [post]
func (fc *FormController) Create(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	var formDto types.CreateFormDto
	if err := c.BodyParser(&formDto); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}
	newForm, err := fc.formService.Create(c.Context(), userId, &formDto)
	if err != nil {
		return err
	}
	return utils.Success(c, newForm, "Form created successfully")
}

// @Summary Get a form by ID
// @Description Retrieve a form by its ID
// @Tags forms
// @Accept json
// @Produce json
// @Param id path string true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Router /forms/{id} [get]
func (pc *FormController) GetById(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	formId := c.Params("id")

	if formId == "" {
		return utils.BadRequest(c, "Form ID is required")
	}

	form, err := pc.formService.GetById(c.Context(), userId, formId)
	if err != nil {
		return err
	}
	return utils.Success(c, form, "Form fetched successfully")
}

// @Summary Update form status
// @Description Update the status of a form by ID
// @Tags forms
// @Accept json
// @Produce json
// @Param id path string true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Router /forms/{id}/status [patch]
func (pc *FormController) UpdateStatus(c *fiber.Ctx) error {
	formId := c.Params("id")
	if formId == "" {
		return utils.BadRequest(c, "Form ID is required")
	}

	var req struct {
		Action models.FormStatus `json:"action"`
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}

	userId := c.Locals("user_id").(string)

	updatedForm, err := pc.formService.UpdateStatus(
		c.Context(),
		userId,
		formId,
		req.Action,
	)

	if err != nil {
		log.Printf("qqqqqqqqqqqqqqqq %s", err.Error())
		return err
	}

	return utils.Success(c, updatedForm, "Form status updated successfully")
}

// @Summary Delete a form
// @Description Delete a form by its ID
// @Tags forms
// @Accept json
// @Produce json
// @Param id path string true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Router /forms/{id} [delete]
func (pc *FormController) Delete(c *fiber.Ctx) error {
	return utils.Success(c, nil, "Form deleted successfully")
}

func (pc *FormController) UpdateSequence(c *fiber.Ctx) error {
	return utils.Success(c, nil, "Form deleted successfully")
}
