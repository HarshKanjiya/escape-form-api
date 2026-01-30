package middlewares

import (
	"strings"
	"time"

	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupMiddlewares(app *fiber.App, cfg *config.Config) {
	app.Use(recover.New())

	app.Use(Logger())

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(cfg.CORS.Origins, ","),
		AllowMethods:     strings.Join(cfg.CORS.Methods, ","),
		AllowHeaders:     strings.Join(cfg.CORS.Headers, ","),
		AllowCredentials: true,
	}))

	// Rate limiting
	app.Use(RateLimiter(RateLimiterConfig{
		Max:        cfg.RateLimit.Max,
		Expiration: time.Duration(cfg.RateLimit.Expiration) * time.Second,
	}))
}

// ErrorHandler
func ErrorHandler(c *fiber.Ctx, err error) error {
	// AppError
	if appErr, ok := err.(*errors.AppError); ok {
		return c.Status(appErr.StatusCode).JSON(fiber.Map{
			"type":       "error",
			"data":       nil,
			"totalCount": 0,
			"message":    appErr.Message,
		})
	}

	// Fiber built-in HTTP errors (404, etc.)
	if fiberErr, ok := err.(*fiber.Error); ok {
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"type":       "error",
			"data":       nil,
			"totalCount": 0,
			"message":    fiberErr.Message,
		})
	}

	// Unknown / panic / unhandled error
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"type":       "error",
		"data":       nil,
		"totalCount": 0,
		"message":    "Something went wrong",
	})
}
