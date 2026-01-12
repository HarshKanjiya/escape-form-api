package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProjectController struct {
	validator   *validator.Validate
	teamService *services.TeamService
}

func NewProjectController(*services.ProjectService) *ProjectController {
	return &ProjectController{
		validator: validator.New(),
	}
}

// @Summary Get all projects
// @Description Retrieve a list of projects
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /projects [get]
func (pc *ProjectController) Get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ProjectController Get method called",
	})
}

// @Summary Create a new project
// @Description Create a new project
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /projects [post]
func (pc *ProjectController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ProjectController Create method called",
	})
}

// @Summary Update a project
// @Description Update an existing project by ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Router /projects/{id} [patch]
func (pc *ProjectController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ProjectController Update method called",
	})
}

// @Summary Delete a project
// @Description Delete a project by ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Router /projects/{id} [delete]
func (pc *ProjectController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ProjectController Delete method called",
	})
}
