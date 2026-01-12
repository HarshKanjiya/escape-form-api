package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TeamController struct {
	validator   *validator.Validate
	teamService *services.TeamService
}

func NewTeamController(*services.TeamService) *TeamController {
	return &TeamController{
		validator: validator.New(),
	}
}

// @Summary Get all teams
// @Description Retrieve a list of teams
// @Tags teams
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /teams [get]
func (tc *TeamController) Get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TeamController Get method called",
	})
}

// @Summary Create a new team
// @Description Create a new team
// @Tags teams
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /teams [post]
func (ts *TeamController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TeamController Get method called",
	})
}

// @Summary Update a team
// @Description Update an existing team by ID
// @Tags teams
// @Accept json
// @Produce json
// @Param id path string true "Team ID"
// @Success 200 {object} map[string]interface{}
// @Router /teams/{id} [patch]
func (tc *TeamController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TeamController Get method called",
	})
}

// @Summary Delete a team
// @Description Delete a team by ID
// @Tags teams
// @Accept json
// @Produce json
// @Param id path string true "Team ID"
// @Success 200 {object} map[string]interface{}
// @Router /teams/{id} [delete]
func (tc *TeamController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TeamController Get method called",
	})
}
