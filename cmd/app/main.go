package main

import (
	"log"
	"time"

	"github.com/HarshKanjiya/escape-form-api/docs"
	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/database"
	"github.com/HarshKanjiya/escape-form-api/internal/middlewares"
	"github.com/HarshKanjiya/escape-form-api/internal/routes"
	"github.com/HarshKanjiya/escape-form-api/internal/storage"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func startServer(app *fiber.App, cfg *config.Config) {
	log.Printf("Starting %s server on port %s", cfg.App.Name, cfg.App.Port)
	log.Printf("API Documentation: http://localhost:%s/swagger/index.html", cfg.App.Port)
	log.Printf("Health Check: http://localhost:%s/api/v1/health", cfg.App.Port)
	log.Fatal(app.Listen(":" + cfg.App.Port))
}

func main() {
	// Initialize Swagger docs
	docs.SwaggerInfo.Title = "EscapeForm API"
	docs.SwaggerInfo.Description = "A scalable REST API built with Fiber framework for user management and authentication"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:7578"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	clerk.SetKey(cfg.Clerk.SecretKey)

	// Setup logger
	zerolog.TimeFieldFormat = time.RFC3339
	if cfg.Logging.Level == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if cfg.Logging.Level == "info" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Connect to AWS S3
	if err := storage.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to AWS S3:", err)
	}

	// Run migrations
	// if err := database.AutoMigrate(&models.User{}); err != nil {
	// 	log.Fatal("Failed to run migrations:", err)
	// }

	app := fiber.New(fiber.Config{
		AppName:               cfg.App.Name,
		ErrorHandler:          middlewares.ErrorHandler,
		DisableStartupMessage: true,
	})

	middlewares.SetupMiddlewares(app, cfg)

	routes.SetupRoutes(app, cfg)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Start server in a goroutine
	go startServer(app, cfg)
	select {}
}
