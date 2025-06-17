package middlewares

import (
	"net/http"
	"strings"

	"github.com/WelintonJunior/identity-access-management-go/cmd/auth"
	"github.com/gofiber/fiber/v2"
)

func JwtAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Authorization header missing",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Authorization header format must be Bearer {token}",
			})
		}

		token := parts[1]

		email, err := auth.VerifyToken(token)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid or expired token",
			})
		}

		c.Locals("userEmail", email)

		return c.Next()
	}
}
