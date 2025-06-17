package repository

import (
	"github.com/WelintonJunior/identity-access-management-go/commons"
	infraestructure "github.com/WelintonJunior/identity-access-management-go/infraestructure/postgres"
	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProducts(product types.Product) (uuid.UUID, error)
	ListProducts(filters map[string]interface{}) ([]types.Product, error)
	GetProductById(id uint) (types.Product, error)
	UpdateProductById(id uint, updatedProduct types.Product) (types.Product, error)
	DeleteProductById(id uint) error
}

type ProductGormRepository struct {
	gormDb *gorm.DB
}

func NewProductRepository() *ProductGormRepository {
	return &ProductGormRepository{gormDb: infraestructure.Db}
}

func (r *ProductGormRepository) CreateProducts(product types.Product) (uuid.UUID, error) {
	return commons.CreateRepoRegister[types.Product](product)
}

func (r *ProductGormRepository) ListProducts(filters map[string]interface{}) ([]types.Product, error) {
	return commons.ListRepoRegisters[types.Product](filters)
}

func (r *ProductGormRepository) GetProductById(id uuid.UUID) (types.Product, error) {
	return commons.GetRepoRegisterById[types.Product](id)
}

func (r *ProductGormRepository) UpdateProductById(id uuid.UUID, updatedProduct types.Product) (types.Product, error) {
	return commons.UpdateRepoRegisterById[types.Product](id, updatedProduct)
}

func (r *ProductGormRepository) DeleteProductById(id uuid.UUID) error {
	return commons.DeleteRepoRegisterById[types.Product](id)
}
