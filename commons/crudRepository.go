package commons

import (
	"fmt"

	infraestructure "github.com/WelintonJunior/identity-access-management-go/infraestructure/postgres"
	"github.com/google/uuid"
)

type HasID interface {
	GetID() uuid.UUID
}

func CreateRepoRegister[T HasID](record T) (uuid.UUID, error) {
	if err := infraestructure.Db.Create(&record).Error; err != nil {
		return uuid.Nil, err
	}
	return record.GetID(), nil
}

func ListRepoRegisters[T any](filters map[string]interface{}) ([]T, error) {
	var records []T

	if err := infraestructure.Db.Where(filters).Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func GetRepoRegisterById[T any](id uuid.UUID) (T, error) {
	var record T

	if err := infraestructure.Db.Where("id = ?", id).First(&record).Error; err != nil {
		var zero T
		return zero, err
	}

	return record, nil
}

func UpdateRepoRegisterById[T any](id uuid.UUID, updated T) (T, error) {
	var empty T

	result := infraestructure.Db.Model(&empty).Where("id = ?", id).Updates(updated)
	if result.Error != nil {
		var zero T
		return zero, result.Error
	}

	if result.RowsAffected == 0 {
		var zero T
		return zero, fmt.Errorf("record with id %s not found", id)
	}

	return updated, nil
}

func DeleteRepoRegisterById[T any](id uuid.UUID) error {
	var record T

	result := infraestructure.Db.Where("id = ?", id).Delete(&record)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found for id: %s", id)
	}

	return nil
}
