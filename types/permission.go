package types

type Permission struct {
	Base
	Name        string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Description string `gorm:"type:varchar(100);"`
}
