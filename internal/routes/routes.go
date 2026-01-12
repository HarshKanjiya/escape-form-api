package routes

import (
	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/controllers"
	"github.com/HarshKanjiya/escape-form-api/internal/handlers"
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, cfg *config.Config) {

	// Initialize services
	// userService := services.NewUserService()
	teamService := services.NewTeamService()

	// Initialize controllers
	// userController := controllers.NewUserController(userService, cfg)
	teamController := controllers.NewTeamController(teamService, cfg)

	// API v1 routes
	api := app.Group("/api/v1")

	// Health check endpoint (no authentication required)
	api.Get("/health", handlers.HealthCheck)

	// Public auth routes (no JWT required)
	// auth := api.Group("/auth")
	// {
	// 	auth.Post("/register", userController.Register)
	// 	auth.Post("/login", userController.Login)
	// }

	// Protected user routes (JWT required)
	// users := api.Group("/users", middlewares.JWTMiddleware(cfg.JWT.Secret))
	teams := api.Group("/teams")
	{
		// teams.Get("/", teamController.GetAll)
		// teams.Post("/", teamController.Create)
		// teams.Get("/:id", teamController.GetByID)
		// teams.Put("/:id", teamController.Update)
		// teams.Delete("/:id", teamController.Delete)
	}
	// {
	// 	// All authenticated users can access these
	// 	users.Get("/", userController.GetAll)
	// 	users.Get("/search", userController.Search)
	// 	users.Get("/:id", userController.GetByID)

	// 	// Only admins can modify users
	// 	users.Post("/", userController.Register)
	// 	users.Put("/:id", userController.Update)
	// 	users.Delete("/:id", userController.Delete)
	// }
}
