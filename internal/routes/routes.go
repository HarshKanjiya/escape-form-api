package routes

import (
	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/controllers"
	"github.com/HarshKanjiya/escape-form-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, cfg *config.Config) {

	// Initialize services
	// userService := services.NewUserService()
	teamService := services.NewTeamService(cfg)
	projectService := services.NewProjectService(cfg)
	formService := services.NewFormService(cfg)
	questionService := services.NewQuestionService(cfg)
	edgeService := services.NewEdgeService(cfg)
	dashService := services.NewDashService(cfg)

	// Initialize controllers
	// userController := controllers.NewUserController(userService, cfg)
	teamController := controllers.NewTeamController(teamService)
	projectController := controllers.NewProjectController(projectService)
	formController := controllers.NewFormController(formService)
	questionController := controllers.NewQuestionController(questionService)
	edgeController := controllers.NewEdgeController(edgeService)
	dashController := controllers.NewDashController(dashService)

	// API v1 routes
	api := app.Group("/api/v1")

	// Health check endpoint (no authentication required)
	api.Get("/health", controllers.HealthCheck)

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
		teams.Get("/", teamController.Get)
		teams.Post("/", teamController.Create)
		teams.Patch("/:id", teamController.Update)
		teams.Delete("/:id", teamController.Delete)
	}

	projects := api.Group("/projects")
	{
		projects.Get("/", projectController.Get)
		projects.Post("/", projectController.Create)
		projects.Patch("/:id", projectController.Update)
		projects.Delete("/:id", projectController.Delete)
	}

	forms := api.Group("/forms")
	{
		forms.Get("/", formController.Get)
		forms.Post("/", formController.Create)
		forms.Get("/:id", formController.GetById)
		forms.Patch("/:id/status", formController.UpdateStatus)
		forms.Delete("/:id/status", formController.UpdateStatus)
	}

	questions := api.Group("/questions")
	{
		questions.Get("/", questionController.Get)
		questions.Post("/", questionController.Create)
		questions.Patch("/:id", questionController.Update)
		questions.Delete("/:id", questionController.Delete)
	}

	edges := api.Group("/edges")
	{
		edges.Get("/", edgeController.Get)
		edges.Post("/", edgeController.Create)
		edges.Patch("/:id", edgeController.Update)
		edges.Delete("/:id", edgeController.Delete)
	}

	dash := api.Group("/dashboard")
	{
		dash.Get("/:formId/analytics", dashController.GetAnalytics)
		dash.Get("/:formId/questions", dashController.GetQuestions)
		dash.Get("/:formId/responses", dashController.GetResponses)
		dash.Patch("/:formId/security", dashController.UpdateSecurity)
		dash.Patch("/:formId/settings", dashController.UpdateSettings)

		// Password management
		dash.Get("/:formId/passwords", dashController.GetPasswords)
		dash.Post("/:formId/passwords", dashController.CreatePasswords)
		dash.Patch("/:formId/passwords/:passwordId", dashController.UpdatePasswords)
		dash.Delete("/:formId/passwords/:passwordId", dashController.DeletePasswords)
	}

}
