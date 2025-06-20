package types

import (
	"time"

	"github.com/google/uuid"
)

type LoginAttempt struct {
	Base
	FailedLoginAttempts int `gorm:"default:0"`
	LockoutExpiresAt    *time.Time
	UserID              uuid.UUID `gorm:"type:uuid;primaryKey"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (e LoginAttempt) GetID() uuid.UUID {
	return e.ID
}
