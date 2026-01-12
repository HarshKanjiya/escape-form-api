package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type EdgeController struct {
}

func NewEdgeController(*services.EdgeService, *config.Config) *EdgeController {
	return &EdgeController{}
}

func (ec *EdgeController) Get(c *fiber.Ctx) error {
	// Placeholder for Get method implementation
	return c.JSON(fiber.Map{
		"message": "EdgeController Get method called",
	})
}

func (ec *EdgeController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "EdgeController Create method called",
	})
}

func (ec *EdgeController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "EdgeController Update method called",
	})
}

func (ec *EdgeController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "EdgeController Delete method called",
	})
}
