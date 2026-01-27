package middlewares

import (
	"log"
	"strings"

	"github.com/HarshKanjiya/escape-form-api/pkg/utils"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/gofiber/fiber/v2"
)

func ClerkAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or invalid token",
			})
		}
		sessionToken := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwt.Verify(c.Context(), &jwt.VerifyParams{
			Token: sessionToken,
		})
		if err != nil {
			log.Printf("JWT verification failed: %v", err)
			return utils.Unauthorized(c, "Invalid or expired token")
		}
		// log.Print(claims)
		c.Locals("user_claims", claims)
		c.Locals("user_id", claims.Subject)

		return c.Next()
	}
}
