package types

import "github.com/google/uuid"

type User struct {
	Base
	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null" json:"-"`
	FullName string `json:"full_name" gorm:"type:varchar(100);not null"`
	IsActive bool
}

type UserResponse struct {
	Message User `json:"message"`
	Success bool `json:"success"`
}

type ListUserResponse struct {
	Message []User `json:"message"`
	Success bool   `json:"success"`
}

func (e User) GetID() uuid.UUID {
	return e.ID
}
