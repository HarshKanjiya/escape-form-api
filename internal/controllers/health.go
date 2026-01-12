package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// @Router /api/v1/health [get]
func HealthCheck(c *fiber.Ctx) error {
	// Simple health check
	return c.JSON(fiber.Map{
		"status": "healthy",
	})
}
