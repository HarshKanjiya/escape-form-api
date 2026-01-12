package main

import (
	"log"

	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"github.com/HarshKanjiya/escape-form-api/internal/database"
	"github.com/HarshKanjiya/escape-form-api/internal/models"
	"gorm.io/gen"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Get the DB instance
	db := database.DB

	// Initialize GORM Gen
	g := gen.NewGenerator(gen.Config{
		OutPath:           "internal/query",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	// Use the database connection
	g.UseDB(db)

	// Apply all models
	g.ApplyBasic(
		models.Team{},
		models.Project{},
		models.Form{},
		models.Question{},
		models.Edge{},
		models.Response{},
		models.User{},
		models.Plan{},
		models.Transaction{},
		models.Coupon{},
		models.Feature{},
		models.ActivePassword{},
		models.QuestionOption{},
	)

	// Generate the code
	g.Execute()
}
