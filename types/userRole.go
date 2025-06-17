package types

import "github.com/google/uuid"

type UserRole struct {
	UserID uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoleID uuid.UUID `gorm:"type:uuid;primaryKey"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Role Role `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
}
