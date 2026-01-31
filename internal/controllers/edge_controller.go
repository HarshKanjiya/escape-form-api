package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type EdgeController struct {
	validator   *validator.Validate
	edgeService services.IEdgeService
}

func NewEdgeController(service services.IEdgeService) *EdgeController {
	return &EdgeController{
		validator:   validator.New(),
		edgeService: service,
	}
}

// @Summary Get all edges for a form
// @Description Retrieve a list of edges for the specified form
// @Tags edges
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {array} types.EdgeDto
// @Router /forms/{formId}/edges [get]
func (ec *EdgeController) Get(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")

	edges, err := ec.edgeService.Get(c.Context(), userId, formId)
	if err != nil {
		return err
	}
	return utils.Success(c, edges, "Edges fetched successfully")
}

// @Summary Create a new edge
// @Description Create a new edge for the specified form
// @Tags edges
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param edge body types.CreateEdgeRequest true "Edge creation data"
// @Success 201 {object} types.EdgeDto
// @Router /forms/{formId}/edges [post]
func (ec *EdgeController) Create(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")

	var edgeDto types.CreateEdgeRequest
	if err := c.BodyParser(&edgeDto); err != nil {
		return errors.BadRequest("Invalid request body")
	}
	if err := ec.validator.Struct(&edgeDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}
	edge, err := ec.edgeService.Create(c.Context(), userId, formId, &edgeDto)
	if err != nil {
		return err
	}
	return utils.Created(c, edge, "Edge created successfully")
}

// @Summary Update an edge
// @Description Update an existing edge by ID
// @Tags edges
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param id path string true "Edge ID"
// @Param edge body types.UpdateEdgeRequest true "Edge update data"
// @Success 200 {object} map[string]interface{}
// @Router /forms/{formId}/edges/{id} [patch]
func (ec *EdgeController) Update(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	edgeId := c.Params("id")
	if edgeId == "" {
		return errors.BadRequest("edgeId is required")
	}

	var edgeDto types.UpdateEdgeRequest
	if err := c.BodyParser(&edgeDto); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	if err := ec.validator.Struct(&edgeDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}
	err := ec.edgeService.Update(c.Context(), userId, formId, edgeId, &edgeDto)
	if err != nil {
		return err
	}
	return utils.Success(c, nil, "Edge updated successfully")
}

// @Summary Delete an edge
// @Description Delete an edge by ID
// @Tags edges
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param id path string true "Edge ID"
// @Success 200 {object} map[string]interface{}
// @Router /forms/{formId}/edges/{id} [delete]
func (ec *EdgeController) Delete(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Params("formId")
	edgeId := c.Params("id")
	if edgeId == "" {
		return errors.BadRequest("edgeId is required")
	}

	err := ec.edgeService.Delete(c.Context(), userId, formId, edgeId)
	if err != nil {
		return err
	}
	return utils.Success(c, nil, "Edge deleted successfully")
}
