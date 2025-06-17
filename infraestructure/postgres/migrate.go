package infraestructure

import (
	"errors"

	"github.com/WelintonJunior/identity-access-management-go/types"
	"gorm.io/gorm"
)

type PostgresMigrateService struct {
	db *gorm.DB
}

func NewPostgresMigrateService(db *gorm.DB) (*PostgresMigrateService, error) {
	var service PostgresMigrateService

	if db == nil {
		return nil, errors.New("No gorm db passes")
	}

	service.db = db
	return &service, nil
}

func (r *PostgresMigrateService) MigrateApply() error {
	return r.db.AutoMigrate(
		&types.User{},
		&types.AuditLog{},
		&types.Permission{},
		&types.RefreshToken{},
		&types.Role{},
		&types.RolePermission{},
		&types.UserRole{},
	)
}

func (r *PostgresMigrateService) MigrateRevert() error {
	return r.db.Migrator().DropTable(
		&types.User{},
		&types.AuditLog{},
		&types.Permission{},
		&types.RefreshToken{},
		&types.Role{},
		&types.RolePermission{},
		&types.UserRole{},
	)
}
