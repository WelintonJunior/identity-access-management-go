package commons

import (
	"fmt"

	infraestructure "github.com/WelintonJunior/identity-access-management-go/infraestructure/postgres"
	"github.com/google/uuid"
)

type HasID interface {
	GetID() uuid.UUID
}

func CreateRepoRegister[T HasID](register T) (uuid.UUID, error) {
	if err := infraestructure.Db.Create(&register).Error; err != nil {
		return uuid.Nil, err
	}
	return register.GetID(), nil
}

func ListRepoRegisters[T any](filters map[string]interface{}) ([]T, error) {
	var registers []T

	result := infraestructure.Db.Where(filters).Find(&registers)
	if result.Error != nil {
		return nil, result.Error
	}

	return registers, nil
}

func GetRepoRegisterById[T any](id uuid.UUID) (T, error) {
	var register T

	// 	query := infraestructure.Db
	// for _, preload := range preloads {
	// 	query = query.Preload(preload)
	// }

	// result := query.Where("id = ?", id).First(&register)

	result := infraestructure.Db.Where("id = ?", id).First(&register)

	if result.Error != nil {
		var zero T
		return zero, result.Error
	}

	return register, nil
}

func UpdateRepoRegisterById[T any](id uuid.UUID, updatedRegister T) (T, error) {
	var empty T
	result := infraestructure.Db.Model(&empty).Where("id = ?", id).Updates(updatedRegister)

	if result.Error != nil {
		var zero T
		return zero, result.Error
	}

	if result.RowsAffected == 0 {
		var zero T
		return zero, fmt.Errorf("registro com id %d não encontrado", id)
	}

	return updatedRegister, nil
}

func DeleteRepoRegisterById[T any](id uuid.UUID) error {
	var register T
	result := infraestructure.Db.Where("id = ?", id).Delete(&register)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("registro não encontrado para o id: %d", id)
	}

	return nil
}
