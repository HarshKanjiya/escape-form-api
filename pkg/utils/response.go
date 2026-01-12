package utils

import (
	"github.com/gofiber/fiber/v2"
)

// Response represents a standard API response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Meta    interface{} `json:"meta,omitempty"` // For pagination, etc.
}

// Success sends a successful response
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SuccessWithMeta sends a successful response with metadata (e.g., pagination)
func SuccessWithMeta(c *fiber.Ctx, message string, data interface{}, meta interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// Created sends a 201 Created response
func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// NoContent sends a 204 No Content response
func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

// Error sends an error response
func Error(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(Response{
		Success: false,
		Error:   message,
	})
}

// BadRequest sends a 400 Bad Request response
func BadRequest(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusBadRequest, message)
}

// Unauthorized sends a 401 Unauthorized response
func Unauthorized(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Unauthorized access"
	}
	return Error(c, fiber.StatusUnauthorized, message)
}

// Forbidden sends a 403 Forbidden response
func Forbidden(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Access forbidden"
	}
	return Error(c, fiber.StatusForbidden, message)
}

// NotFound sends a 404 Not Found response
func NotFound(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Resource not found"
	}
	return Error(c, fiber.StatusNotFound, message)
}

// Conflict sends a 409 Conflict response
func Conflict(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusConflict, message)
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Internal server error"
	}
	return Error(c, fiber.StatusInternalServerError, message)
}

// ValidationError sends a 422 Unprocessable Entity response for validation errors
func ValidationError(c *fiber.Ctx, errors interface{}) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(Response{
		Success: false,
		Error:   "Validation failed",
		Data:    errors,
	})
}
