package middlewares

import (
	"strings"

	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
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
	// app.Use(RateLimiter(RateLimiterConfig{
	// 	Max:        cfg.RateLimit.Max,
	// 	Expiration: time.Duration(cfg.RateLimit.Expiration) * time.Second,
	// }))
}

// ErrorHandler
func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if appErr, ok := err.(*utils.AppError); ok {
		code = appErr.StatusCode
		message = appErr.Message
	} else if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"type":       "error",
		"data":       nil,
		"message":    message,
		"totalItems": 0,
	})
}
