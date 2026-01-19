package controllers

import (
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
	pagination := &types.PaginationQuery{
		Page:   c.QueryInt("page", 1),
		Limit:  c.QueryInt("limit", 10),
		Search: c.Query("search", ""),
		SortBy: c.Query("sortBy", ""),
		Order:  c.Query("order", ""),
	}
	projectId := c.Query("projectId", "")
	if projectId == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "projectId is required",
		})
	}
	forms, total := fc.formService.Get(c, pagination, true, projectId)
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
	return c.JSON(fiber.Map{
		"message": "FormController Create method called",
	})
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
	return c.JSON(fiber.Map{
		"message": "FormController GetById method called",
	})
}

// @Summary Update form status
// @Description Update the status of a form by ID
// @Tags forms
// @Accept json
// @Produce json
// @Param id path string true "Form ID"
// @Success 200 {object} map[string]interface{}
// @Router /forms/{id}/status [patch]
// @Router /forms/{id}/status [delete]
func (pc *FormController) UpdateStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "FormController UpdateStatus method called",
	})
}

func (pc *FormController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "FormController Update method called",
	})
}

func (pc *FormController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "FormController Delete method called",
	})
}
