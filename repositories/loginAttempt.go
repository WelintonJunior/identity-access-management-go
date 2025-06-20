package repository

import (
	"errors"

	"github.com/WelintonJunior/identity-access-management-go/commons"
	infraestructure "github.com/WelintonJunior/identity-access-management-go/infraestructure/postgres"
	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginAttemptRepository interface {
	CreateLoginAttempts(loginAttempt types.LoginAttempt) (uuid.UUID, error)
	Save(loginAttempt *types.LoginAttempt) error
	FindLoginAttemptsByUserID(userID uuid.UUID) (*types.LoginAttempt, error)
	ListLoginAttempts(filters map[string]interface{}) ([]types.LoginAttempt, error)
	GetLoginAttemptById(id uint) (types.LoginAttempt, error)
	UpdateLoginAttemptById(id uint, updatedLoginAttempt types.LoginAttempt) (types.LoginAttempt, error)
	DeleteLoginAttemptById(id uint) error
}

type LoginAttemptGormRepository struct {
	gormDb *gorm.DB
}

func NewLoginAttemptRepository() *LoginAttemptGormRepository {
	return &LoginAttemptGormRepository{gormDb: infraestructure.Db}
}

func (r *LoginAttemptGormRepository) CreateLoginAttempts(loginAttempt types.LoginAttempt) (uuid.UUID, error) {
	return commons.CreateRepoRegister[types.LoginAttempt](loginAttempt)
}

func (r *LoginAttemptGormRepository) Save(loginAttempt *types.LoginAttempt) error {
	var existing types.LoginAttempt
	err := r.gormDb.
		Where("user_id = ?", loginAttempt.UserID).
		First(&existing).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.gormDb.Create(loginAttempt).Error
	} else if err != nil {
		return err
	}

	return r.gormDb.Model(&existing).Updates(map[string]interface{}{
		"failed_login_attempts": loginAttempt.FailedLoginAttempts,
		"lockout_expires_at":    loginAttempt.LockoutExpiresAt,
	}).Error
}

func (r *LoginAttemptGormRepository) FindLoginAttemptsByUserID(userID uuid.UUID) (*types.LoginAttempt, error) {
	var loginAttempt types.LoginAttempt

	err := r.gormDb.
		Where("user_id = ?", userID).
		First(&loginAttempt).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &loginAttempt, nil
}

func (r *LoginAttemptGormRepository) ListLoginAttempts(filters map[string]interface{}) ([]types.LoginAttempt, error) {
	return commons.ListRepoRegisters[types.LoginAttempt](filters)
}

func (r *LoginAttemptGormRepository) GetLoginAttemptById(id uuid.UUID) (types.LoginAttempt, error) {
	return commons.GetRepoRegisterById[types.LoginAttempt](id)
}

func (r *LoginAttemptGormRepository) UpdateLoginAttemptById(id uuid.UUID, updatedLoginAttempt types.LoginAttempt) (types.LoginAttempt, error) {
	return commons.UpdateRepoRegisterById[types.LoginAttempt](id, updatedLoginAttempt)
}

func (r *LoginAttemptGormRepository) DeleteLoginAttemptById(id uuid.UUID) error {
	return commons.DeleteRepoRegisterById[types.LoginAttempt](id)
}
