package types

import "github.com/google/uuid"

type RolePermission struct {
	PermissionID uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoleID       uuid.UUID `gorm:"type:uuid;primaryKey"`

	Permission Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE"`
	Role       Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
}
