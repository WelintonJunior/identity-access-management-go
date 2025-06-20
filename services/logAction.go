package services

import (
	infraestructure "github.com/WelintonJunior/identity-access-management-go/infraestructure/postgres"
	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func LogAction(c *fiber.Ctx, userID uuid.UUID, action string) error {
	ip := c.IP()
	userAgent := c.Get("User-Agent")

	auditLog := types.AuditLog{
		UserID:    userID,
		Action:    action,
		IpAddress: ip,
		UserAgent: userAgent,
	}

	return infraestructure.Db.WithContext(c.Context()).Create(&auditLog).Error
}
