package middlewares

import (
	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gofiber/fiber/v2"
)

func ClerkAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := string(c.Request().Header.Peek("Authorization"))
		if authHeader == "" {
			return c.SendStatus(fiber.StatusNotFound)
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			return c.SendStatus(fiber.StatusNotFound)
		}

		var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := clerk.SessionClaimsFromContext(r.Context())
			if !ok {
				c.SendStatus(fiber.StatusNotFound)
				return
			}

			c.Locals("clerk_session", claims)
			c.Next()
		})
		// mw := clerkhttp.WithHeaderAuthorization()(handler)

		// req := c.Request().Ctx().Value(fiber.HeaderAuthorizationKey)
		// _ = req

		return nil
	}
}
