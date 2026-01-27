package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProjectController struct {
	validator      *validator.Validate
	projectService *services.ProjectService
}

func NewProjectController(service *services.ProjectService) *ProjectController {
	return &ProjectController{
		validator:      validator.New(),
		projectService: service,
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
	pagination := &types.PaginationQuery{
		Page:   c.QueryInt("page", 1),
		Limit:  c.QueryInt("limit", 10),
		Search: c.Query("search", ""),
		SortBy: c.Query("sortBy", ""),
		Order:  c.Query("order", ""),
	}
	teamId := c.Query("teamId", "")
	projects, totalCount, err := pc.projectService.Get(c, pagination, true, teamId)
	if err != nil {
		return utils.InternalServerError(c, "Failed to fetch projects")
	}
	return utils.Success(c, projects, "Projects fetched successfully", totalCount)
}

// @Summary Get a project by ID
// @Description Retrieve a project by its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param projectId path string true "Project ID"
// @Success 200 {object} map[string]interface{}
// @Router /projects/{projectId} [get]
func (pc *ProjectController) GetById(c *fiber.Ctx) error {
	projectId := c.Params("projectId")
	project, err := pc.projectService.GetById(c, projectId)
	if err != nil {
		return utils.NotFound(c, "Project not found")
	}
	return utils.Success(c, project, "Project fetched successfully")
}

// @Summary Create a new project
// @Description Create a new project
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /projects [post]
func (pc *ProjectController) Create(c *fiber.Ctx) error {
	projectDto := new(types.ProjectDto)
	if err := c.BodyParser(projectDto); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}

	if err := pc.validator.Struct(projectDto); err != nil {
		return utils.BadRequest(c, "Validation failed")
	}

	createdProject, err := pc.projectService.Create(c, projectDto)
	if err != nil {
		return utils.BadRequest(c, "Failed to create project")
	}

	return utils.Success(c, createdProject, "Project created successfully", 0)
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
	projectDto := new(types.ProjectDto)
	if err := c.BodyParser(projectDto); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}

	if err := pc.validator.Struct(projectDto); err != nil {
		return utils.BadRequest(c, "Validation failed")
	}

	ok, err := pc.projectService.Update(c, &types.ProjectDto{
		ID:          c.Params("id"),
		Name:        projectDto.Name,
		Description: projectDto.Description,
		TeamID:      projectDto.TeamID,
	})
	if err != nil || !ok {
		return utils.BadRequest(c, "Failed to update project")
	}
	return utils.Success(c, nil, "Project updated successfully")
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
	projectId := c.Params("id")

	if projectId == "" {
		return utils.BadRequest(c, "Project ID is required")
	}

	ok, err := pc.projectService.Delete(c, projectId)
	if err != nil || !ok {
		return utils.BadRequest(c, "Failed to delete project")
	}
	return utils.Success(c, nil, "Project deleted successfully")
}
