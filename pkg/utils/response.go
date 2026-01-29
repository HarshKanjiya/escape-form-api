package utils

import (
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, data interface{}, message string, total ...int) error {
	totalCount := 0
	if len(total) > 0 {
		totalCount = total[0]
	}
	return c.Status(fiber.StatusOK).JSON(types.ResponseObj{
		Type:       "success",
		Message:    message,
		Data:       data,
		TotalCount: totalCount,
	})
}

func Created(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusCreated).JSON(types.ResponseObj{
		Type:    "success",
		Message: message,
		Data:    data,
	})
}

func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

func Error(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(types.ResponseObj{
		Type:    "error",
		Message: message,
	})
}

func BadRequest(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusBadRequest, message)
}

func MapErrors(c *fiber.Ctx, errors interface{}, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(types.ResponseObj{
		Type:    "error",
		Message: message,
		Data:    errors,
	})
}

func Unauthorized(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Unauthorized access"
	}
	return Error(c, fiber.StatusUnauthorized, message)
}

func Forbidden(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Access forbidden"
	}
	return Error(c, fiber.StatusForbidden, message)
}

func NotFound(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Resource not found"
	}
	return Error(c, fiber.StatusNotFound, message)
}

func Conflict(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusConflict, message)
}

func InternalServerError(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Internal server error"
	}
	return Error(c, fiber.StatusInternalServerError, message)
}

func ValidationError(c *fiber.Ctx, errors interface{}) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(types.ResponseObj{
		Type:    "error",
		Message: "Validation failed",
		Data:    errors,
	})
}
