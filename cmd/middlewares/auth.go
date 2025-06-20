package middlewares

import (
	"strings"

	"github.com/WelintonJunior/identity-access-management-go/cmd/auth"
	"github.com/gofiber/fiber/v2"
)

func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "authentication token not provided",
			})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid token format",
			})
		}

		tokenStr := tokenParts[1]

		email, err := auth.VerifyToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}

		c.Locals("email", email)
		return c.Next()
	}
}
