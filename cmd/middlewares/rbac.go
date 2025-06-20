package middlewares

import (
	"context"
	"errors"

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

func RequirePermission(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		emailRaw := c.Locals("email")
		if emailRaw == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthenticated user",
			})
		}

		email, ok := emailRaw.(string)
		if !ok || email == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "invalid email in context",
			})
		}

		var user types.User
		err := infraestructure.Db.WithContext(c.Context()).Where("email = ?", email).First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "user not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal error while fetching user",
			})
		}

		ok, err = HasPermission(c.Context(), user.ID, permission)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error",
			})
		}

		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "permission denied",
			})
		}

		return c.Next()
	}
}
