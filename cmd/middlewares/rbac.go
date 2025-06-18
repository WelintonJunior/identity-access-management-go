package middlewares

import (
	"context"
	"errors"
	"strings"

	"github.com/WelintonJunior/identity-access-management-go/cmd/auth"
	infraestructure "github.com/WelintonJunior/identity-access-management-go/infraestructure/postgres"
	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func HasPermission(ctx context.Context, userID uuid.UUID, permissionName string) (bool, error) {
	db := infraestructure.Db.WithContext(ctx)

	var roles []types.Role
	err := db.
		Model(&types.Role{}).
		Joins("JOIN user_roles ur ON ur.role_id = roles.id").
		Where("ur.user_id = ?", userID).
		Find(&roles).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	if len(roles) == 0 {
		return false, nil
	}

	roleIDs := make([]string, len(roles))
	for i, r := range roles {
		roleIDs[i] = r.ID.String()
	}

	var count int64
	err = db.Model(&types.RolePermission{}).
		Joins("JOIN permissions p ON p.id = role_permissions.permission_id").
		Where("role_permissions.role_id IN ?", roleIDs).
		Where("p.name = ?", permissionName).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "token de autenticação não fornecido",
			})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "formato do token inválido",
			})
		}

		tokenStr := tokenParts[1]

		email, err := auth.VerifyToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "token inválido ou expirado",
			})
		}

		c.Locals("email", email)
		return c.Next()
	}
}

func RequirePermission(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		emailRaw := c.Locals("email")
		if emailRaw == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "usuário não autenticado",
			})
		}

		email, ok := emailRaw.(string)
		if !ok || email == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "email inválido no contexto",
			})
		}

		var user types.User
		err := infraestructure.Db.WithContext(c.Context()).Where("email = ?", email).First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "usuário não encontrado",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "erro interno ao buscar usuário",
			})
		}

		ok, err = HasPermission(c.Context(), user.ID, permission)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "erro interno no servidor",
			})
		}

		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "permissão negada",
			})
		}

		return c.Next()
	}
}
