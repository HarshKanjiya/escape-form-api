package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TeamController struct {
	validator   *validator.Validate
	teamService services.ITeamService
}

func NewTeamController(service services.ITeamService) *TeamController {
	return &TeamController{
		validator:   validator.New(),
		teamService: service,
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

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	pagination := &types.PaginationQuery{
		Page:   c.QueryInt("page", 1),
		Limit:  c.QueryInt("limit", 10),
		Search: c.Query("search", ""),
		SortBy: c.Query("sortBy", ""),
		Order:  c.Query("order", ""),
	}
	teams, total := tc.teamService.Get(c, pagination, true)
	return utils.Success(c, teams, "Teams fetched successfully", total)
}

// @Summary Create a new team
// @Description Create a new team
// @Tags teams
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /teams [post]
func (tc *TeamController) Create(c *fiber.Ctx) error {

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	teamDto := new(types.TeamDto)
	if err := c.BodyParser(teamDto); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	if err := tc.validator.Struct(teamDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}

	createdTeam, err := tc.teamService.Create(c, teamDto)
	if err != nil {
		return err
	}

	return utils.Success(c, createdTeam, "Team created successfully", 0)
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

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	teamDto := new(types.TeamDto)
	if err := c.BodyParser(teamDto); err != nil {
		return errors.BadRequest("Invalid request body")
	}

	if err := tc.validator.Struct(teamDto); err != nil {
		return errors.BadRequest("Validation failed: " + err.Error())
	}

	ok, err := tc.teamService.Update(c, &types.TeamDto{
		ID:   c.Params("id"),
		Name: teamDto.Name,
	})
	if err != nil || !ok {
		return err
	}
	return utils.Success(c, nil, "Team updated successfully")
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

	userId, ok := utils.GetUserId(c)
	if ok == false {
		return errors.Unauthorized("")
	}

	teamId := c.Params("id")

	if teamId == "" {
		return errors.BadRequest("Team ID is required")
	}

	ok, err := tc.teamService.Delete(c, teamId)
	if err != nil || !ok {
		return err
	}
	return utils.Success(c, nil, "Team deleted successfully")
}
