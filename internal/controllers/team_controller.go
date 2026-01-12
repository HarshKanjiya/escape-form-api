package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type TeamController struct {
}

func NewTeamController(*services.TeamService, *config.Config) *TeamController {
	return &TeamController{}
}

func (tc *TeamController) Get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TeamController Get method called",
	})
}

func (ts *TeamController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TeamController Get method called",
	})
}

func (tc *TeamController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TeamController Get method called",
	})
}

func (tc *TeamController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TeamController Get method called",
	})
}
