package types

import "github.com/google/uuid"

type Product struct {
	Base
	Name  string
	Value float64
}

type ProductResponse struct {
	Message Product `json:"message"`
	Success bool    `json:"success"`
}

type ListProductResponse struct {
	Message []Product `json:"message"`
	Success bool      `json:"success"`
}

func (e Product) GetID() uuid.UUID {
	return e.ID
}
