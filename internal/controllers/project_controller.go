package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ProjectController struct {
}

func NewProjectController(*services.ProjectService, *config.Config) *ProjectController {
	return &ProjectController{}
}

func (pc *ProjectController) Get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ProjectController Get method called",
	})
}

func (pc *ProjectController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ProjectController Create method called",
	})
}

func (pc *ProjectController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ProjectController Update method called",
	})
}

func (pc *ProjectController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ProjectController Delete method called",
	})
}
