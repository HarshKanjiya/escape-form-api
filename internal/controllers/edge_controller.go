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
	edgeService *services.EdgeService
}

func NewEdgeController(service *services.EdgeService) *EdgeController {
	return &EdgeController{
		validator:   validator.New(),
		edgeService: service,
	}
}

// @Summary Get all edges
// @Description Retrieve a list of edges
// @Tags edges
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /edges [get]
func (ec *EdgeController) Get(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	formId := c.Query("formId")
	if formId == "" {
		return errors.BadRequest("formId is required")
	}
	edges, err := ec.edgeService.Get(c, formId)
	if err != nil {
		return err
	}
	return utils.Success(c, edges, "Edges fetched successfully")
}

// @Summary Create a new edge
// @Description Create a new edge
// @Tags edges
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /edges [post]
func (ec *EdgeController) Create(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	var edgeDto types.EdgeDto
	if err := c.BodyParser(&edgeDto); err != nil {
		return errors.BadRequest("Invalid request body")
	}
	if err := ec.validator.Struct(&edgeDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}
	edge, err := ec.edgeService.Create(c, &edgeDto)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to create edge")
	}
	return utils.Created(c, edge, "Edge created successfully")
}

// @Summary Update an edge
// @Description Update an existing edge by ID
// @Tags edges
// @Accept json
// @Produce json
// @Param id path string true "Edge ID"
// @Success 200 {object} map[string]interface{}
// @Router /edges/{id} [patch]
func (ec *EdgeController) Update(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	id := c.Params("id")
	var edgeDto types.EdgeDto
	if err := c.BodyParser(&edgeDto); err != nil {
		return errors.BadRequest("Invalid request body")
	}
	edgeDto.ID = id
	if err := ec.validator.Struct(&edgeDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}
	edge, err := ec.edgeService.Update(c, &edgeDto)
	if err != nil {
		return err
	}
	return utils.Success(c, edge, "Edge updated successfully")
}

// @Summary Delete an edge
// @Description Delete an edge by ID
// @Tags edges
// @Accept json
// @Produce json
// @Param id path string true "Edge ID"
// @Success 200 {object} map[string]interface{}
// @Router /edges/{id} [delete]
func (ec *EdgeController) Delete(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	id := c.Params("id")
	err := ec.edgeService.Delete(c, id)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to delete edge")
	}
	return utils.NoContent(c)
}
