package controllers

import (
	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type FormController struct {
}

func NewFormController(*services.FormService, *config.Config) *FormController {
	return &FormController{}
}

func (pc *FormController) Get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "FormController Get method called",
	})
}

func (pc *FormController) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "FormController Get method called",
	})
}

func (pc *FormController) GetById(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "FormController GetById method called",
	})
}

func (pc *FormController) UpdateStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "FormController UpdateStatus method called",
	})
}

func (pc *FormController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "FormController Update method called",
	})
}

func (pc *FormController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "FormController Delete method called",
	})
}
