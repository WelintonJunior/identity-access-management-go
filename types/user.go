package types

type User struct {
	Base
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	FullName string `gorm:"type:varchar(100);not null"`
	IsActive bool
}
