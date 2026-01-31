package middlewares

import (
	"strings"

	"github.com/HarshKanjiya/escape-form-api/pkg/errors"
	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func JWTMiddleware(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return errors.Unauthorized("Missing authorization token")
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return errors.Unauthorized("Invalid authorization header format")
		}

		token := parts[1]

		// Validate token
		claims, err := utils.ValidateToken(token, jwtSecret)
		if err != nil {
			log.Error().Err(err).Msg("Token validation failed")
			return errors.Unauthorized("Invalid or expired token")
		}

		// Store user info in context for use in handlers
		c.Locals("userID", claims.UserID)
		c.Locals("userEmail", claims.Email)
		c.Locals("userRole", claims.Role)

		return c.Next()
	}
}

// extract the user ID from the Fiber context
func GetUserID(c *fiber.Ctx) uint {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return 0
	}
	return userID
}

// extract the user email from the Fiber context
func GetUserEmail(c *fiber.Ctx) string {
	email, ok := c.Locals("userEmail").(string)
	if !ok {
		return ""
	}
	return email
}

// extract the user role from the Fiber context
func GetUserRole(c *fiber.Ctx) string {
	role, ok := c.Locals("userRole").(string)
	if !ok {
		return ""
	}
	return role
}
