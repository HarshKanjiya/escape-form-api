package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary Health check
// @Description Check the health status of the API
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/health [get]
func HealthCheck(c *fiber.Ctx) error {
	// Simple health check
	return c.JSON(fiber.Map{
		"status": "healthy",
	})
}
