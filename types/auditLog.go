package types

import "github.com/google/uuid"

type AuditLog struct {
	Base
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL"`

	Action    string `gorm:"type:varchar(255);not null"`
	IpAddress string `gorm:"type:varchar(45)"`
	UserAgent string `gorm:"type:text"`
}
