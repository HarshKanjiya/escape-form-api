package controllers

import (
	"log"

	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FormController struct {
	validator   *validator.Validate
	formService services.IFormService
}

func NewFormController(service services.IFormService) *FormController {
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

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	projectId := c.Query("projectId", "")
	if projectId == "" {
		return errors.BadRequest("projectId is required")
	}

	pagination := &types.PaginationQuery{
		Page:   c.QueryInt("page", 1),
		Limit:  c.QueryInt("limit", 10),
		Search: c.Query("search", ""),
		SortBy: c.Query("sortBy", ""),
		Order:  c.Query("order", ""),
	}

	status := c.Query("status", "")

	forms, totalCount, err := fc.formService.Get(c.Context(), userId, pagination, projectId, status)
	if err != nil {
		log.Printf("--------------------- %v", err)
		return err
	}
	return utils.Success(c, forms, "Forms fetched successfully", totalCount)
}

// @Summary Create a new form
// @Description Create a new form
// @Tags forms
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /forms [post]
func (fc *FormController) Create(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	var createFormDto types.CreateFormRequest
	if err := c.BodyParser(&createFormDto); err != nil {
		log.Println("Error parsing body:", err)
		return errors.BadRequest("Invalid request body")
	}
	if err := fc.validator.Struct(&createFormDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}

	form, err := fc.formService.Create(c.Context(), userId, createFormDto.ProjectID, &createFormDto)
	if err != nil {
		return err
	}
	return utils.Success(c, form, "Form created successfully")
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

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("id")

	if formId == "" {
		return errors.BadRequest("Form ID is required")
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

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("id", "")
	if formId == "" {
		return errors.BadRequest("Form ID is required")
	}

	var statusDto types.UpdateFormStatusRequest
	if err := c.BodyParser(&statusDto); err != nil {
		log.Println("Error parsing body:", err)
		return errors.BadRequest("Invalid request body")
	}
	if err := pc.validator.Struct(&statusDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}
	err := pc.formService.UpdateStatus(c.Context(), userId, formId, statusDto.Status)
	if err != nil {
		return err
	}
	return utils.Success(c, nil, "Form status updated successfully")
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

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("id", "")
	if formId == "" {
		return errors.BadRequest("Form ID is required")
	}
	err := pc.formService.Delete(c.Context(), userId, formId)
	if err != nil {
		return err
	}
	return utils.Success(c, nil, "Form deleted successfully")
}

func (pc *FormController) UpdateSequence(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	if formId == "" {
		return errors.BadRequest("Form ID is required")
	}

	var sequenceDto types.UpdateSequenceRequest
	if err := c.BodyParser(&sequenceDto); err != nil {
		log.Println("Error parsing body:", err)
		return errors.BadRequest("Invalid request body")
	}
	if err := pc.validator.Struct(&sequenceDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}

	err := pc.formService.UpdateSequence(c.Context(), userId, formId, sequenceDto.Sequence)
	if err != nil {
		return err
	}
	return utils.Success(c, nil, "Form deleted successfully")
}
