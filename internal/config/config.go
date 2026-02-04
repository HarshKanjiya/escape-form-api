package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// global configuration
type Config struct {
	App       AppConfig
	Database  DatabaseConfig
	JWT       JWTConfig
	CORS      CORSConfig
	RateLimit RateLimitConfig
	Logging   LoggingConfig
	Clerk     ClerkConfig
	AWS       AWSConfig
}

// application-level configuration
type AppConfig struct {
	Name string
	Env  string
	Port string
	Host string
}

// database connection configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// JWT authentication configuration
type JWTConfig struct {
	Secret        string
	Expiry        time.Duration
	RefreshExpiry time.Duration
}

// CORS configuration
type CORSConfig struct {
	Origins []string
	Methods []string
	Headers []string
}

// rate limiting configuration
type RateLimitConfig struct {
	Max        int
	Expiration int
}

// logging configuration
type LoggingConfig struct {
	Level string
}

// Clerk Config
type ClerkConfig struct {
	SecretKey string
}

type AWSConfig struct {
	AccessKey  string
	SecretKey  string
	Region     string
	BucketName string
	EndPoint   string
}

// automatically loads the correct .env file based on APP_ENV
func Load() (*Config, error) {

	// Load environment variables from file if exists
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	envFile := ".env." + env
	if err := godotenv.Load(envFile); err != nil {
		godotenv.Load()
	}

	cfg := &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "EscapeForm API"),
			Env:  getEnv("APP_ENV", "local"),
			Port: getEnv("APP_PORT", "3000"),
			Host: getEnv("APP_HOST", "localhost"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "escape_form"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:        getEnv("JWT_SECRET", "default-secret-change-this"),
			Expiry:        parseDuration(getEnv("JWT_EXPIRY", "24h")),
			RefreshExpiry: parseDuration(getEnv("JWT_REFRESH_EXPIRY", "168h")),
		},
		CORS: CORSConfig{
			Origins: splitAndTrim(getEnv("CORS_ORIGINS", "http://localhost:3000,http://localhost:8080,http://127.0.0.1:3000,http://127.0.0.1:8080,https://escform.com,https://dashboard.escform.com,https://form.escform.com")),
			Methods: splitAndTrim(getEnv("CORS_METHODS", "GET,POST,PUT,DELETE,PATCH,OPTIONS")),
			Headers: splitAndTrim(getEnv("CORS_HEADERS", "Content-Type,Authorization")),
		},
		RateLimit: RateLimitConfig{
			Max:        getEnvAsInt("RATE_LIMIT_MAX", 100),
			Expiration: getEnvAsInt("RATE_LIMIT_EXPIRATION", 60),
		},
		Logging: LoggingConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
		Clerk: ClerkConfig{
			SecretKey: getEnv("CLERK_SK", "sk_test_XXX"),
		},
		AWS: AWSConfig{
			AccessKey:  getEnv("AWS_ACCESS_KEY", ""),
			SecretKey:  getEnv("AWS_SECRET_KEY", ""),
			Region:     getEnv("AWS_REGION", "auto"),
			BucketName: getEnv("AWS_BUCKET_NAME", ""),
			EndPoint:   getEnv("AWS_ENDPOINT", ""),
		},
	}

	return cfg, nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

// getEnvAsInt retrieves an environment variable as integer or returns a default value
func getEnvAsInt(key string, defaultVal int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultVal
}

// parseDuration parses a duration string, returns 24h on error
func parseDuration(durationStr string) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 24 * time.Hour
	}
	return duration
}

// splitAndTrim splits a comma-separated string and trims whitespace
func splitAndTrim(s string) []string {
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// GetDSN returns the database DSN string for PostgreSQL
func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.Database.Host,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.Port,
		c.Database.SSLMode,
	)
}
