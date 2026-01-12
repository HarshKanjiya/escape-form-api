package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type DashController struct {
}

func NewDashController(*services.DashService, *config.Config) *DashController {
	return &DashController{}
}

func (pc *DashController) GetAnalytics(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController Get method called",
	})
}

func (pc *DashController) GetQuestions(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController Get method called",
	})
}

func (pc *DashController) GetResponses(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController Get method called",
	})
}

func (pc *DashController) GetPasswords(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController Get method called",
	})
}

func (pc *DashController) UpdatePasswords(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController UpdatePasswords method called",
	})
}

func (pc *DashController) CreatePasswords(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController UpdatePasswords method called",
	})
}

func (pc *DashController) DeletePasswords(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController DeletePasswords method called",
	})
}

func (pc *DashController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController Create method called",
	})
}

func (pc *DashController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController Update method called",
	})
}

func (pc *DashController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController Delete method called",
	})
}

func (pc *DashController) UpdateSecurity(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController UpdateSecurity method called",
	})
}

func (pc *DashController) UpdateSettings(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "DashController UpdateSettings method called",
	})
}
