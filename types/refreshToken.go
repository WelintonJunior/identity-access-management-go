package types

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	Token  string    `gorm:"type:text;primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

	ExpiresAt time.Time `gorm:"type:timestamp;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()"`
}
