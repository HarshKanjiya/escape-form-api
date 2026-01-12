package database

import (
	"fmt"
	"log"

	"github.com/HarshKanjiya/escape-form-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// database connection instance
var DB *gorm.DB

func Connect(cfg *config.Config) error {
	dsn := cfg.GetDSN()

	gormConfig := &gorm.Config{}
	if cfg.App.Env == "prod" {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	DB = db
	log.Println("Connected to database successfully")
	return nil
}

// close the database connection
func Close() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err == nil {
			sqlDB.Close()
			log.Println("Database connection closed")
		}
	}
}

// AutoMigrate runs automatic migrations for the provided models
func AutoMigrate(models ...interface{}) error {
	if DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	log.Println("Running database migrations...")
	if err := DB.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}
