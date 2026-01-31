package utils

import (
	"github.com/HarshKanjiya/escape-form-api/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, data interface{}, message string, total ...int64) error {
	totalCount := int64(0)
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

// func BadRequest(c *fiber.Ctx, message string) error {
// 	return Error(c, fiber.StatusBadRequest, message)
// }

// func MapErrors(c *fiber.Ctx, errors interface{}, message string) error {
// 	return c.Status(fiber.StatusBadRequest).JSON(types.ResponseObj{
// 		Type:    "error",
// 		Message: message,
// 		Data:    errors,
// 	})
// }

// func Unauthorized(c *fiber.Ctx, message string) error {
// 	if message == "" {
// 		message = "Unauthorized access"
// 	}
// 	return Error(c, fiber.StatusUnauthorized, message)
// }

// func Conflict(c *fiber.Ctx, message string) error {
// 	return Error(c, fiber.StatusConflict, message)
// }

// func ValidationError(c *fiber.Ctx, errors interface{}) error {
// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(types.ResponseObj{
// 		Type:    "error",
// 		Message: "Validation failed",
// 		Data:    errors,
// 	})
// }
