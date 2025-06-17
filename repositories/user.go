package repository

import (
	"fmt"

	"github.com/WelintonJunior/identity-access-management-go/commons"
	infraestructure "github.com/WelintonJunior/identity-access-management-go/infraestructure/postgres"
	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/WelintonJunior/identity-access-management-go/validation"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user types.User) error
	FindUserByEmail(email string) (types.User, error)
	ListUsers(filters map[string]interface{}) ([]types.User, error)
	GetUserById(id uint) (types.User, error)
	UpdateUserById(id uint, updatedUser types.User) (types.User, error)
	DeleteUserById(id uint) error
}

type UserGormRepository struct {
	gormDb *gorm.DB
}

func NewUserRepository() *UserGormRepository {
	return &UserGormRepository{gormDb: infraestructure.Db}
}

func (r *UserGormRepository) CreateUser(user types.User) error {
	// Validação dos campos obrigatórios
	if err := validation.ValidateUser(user); err != nil {
		return fmt.Errorf("validação falhou: %w", err)
	}

	// Verifica se já existe um usuário com esse email
	existingUser, err := r.FindUserByEmail(user.Email)
	if err == nil && existingUser.ID != uuid.Nil {
		return fmt.Errorf("usuário com e-mail %s já existe", user.Email)
	}

	// Geração do hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("erro ao gerar hash da senha: %w", err)
	}
	user.Password = string(hashedPassword)

	// Criação do usuário no banco de dados
	if err := r.gormDb.Create(&user).Error; err != nil {
		return fmt.Errorf("erro ao criar usuário no banco: %w", err)
	}

	return nil
}

// Criação das repository manuais
func (r *UserGormRepository) FindUserByEmail(email string) (types.User, error) {
	var user types.User
	result := r.gormDb.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r *UserGormRepository) ListUsers(filters map[string]interface{}) ([]types.User, error) {
	return commons.ListRepoRegisters[types.User](filters)
}

func (r *UserGormRepository) GetUserById(id uuid.UUID) (types.User, error) {
	return commons.GetRepoRegisterById[types.User](id)
}

func (r *UserGormRepository) UpdateUserById(id uuid.UUID, updatedUser types.User) (types.User, error) {
	return commons.UpdateRepoRegisterById[types.User](id, updatedUser)
}

func (r *UserGormRepository) DeleteUserById(id uuid.UUID) error {
	return commons.DeleteRepoRegisterById[types.User](id)
}
