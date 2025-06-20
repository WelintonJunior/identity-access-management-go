package infraestructure

import (
	"errors"
	"log"
	"time"

	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/google/uuid"
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
		&types.Product{},
		&types.LoginAttempt{},
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
		&types.Product{},
		&types.LoginAttempt{},
	)
}

func Seed(db *gorm.DB) error {
	var count int64
	db.Model(&types.User{}).Count(&count)
	if count > 0 {
		log.Println("Seed: data already exists in the database. Skipping seeding.")
		return nil
	}

	adminRole := types.Role{
		Base: types.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
		Name:        "admin",
		Description: "System administrator",
	}
	userRole := types.Role{
		Base: types.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
		Name:        "user",
		Description: "Default user",
	}

	if err := db.Create(&[]types.Role{adminRole, userRole}).Error; err != nil {
		return err
	}

	readPermission := types.Permission{
		Base: types.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
		Name:        "read",
		Description: "Permission to read data",
	}
	writePermission := types.Permission{
		Base: types.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
		Name:        "write",
		Description: "Permission to write data",
	}
	deletePermission := types.Permission{
		Base: types.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
		Name:        "delete",
		Description: "Permission to delete data",
	}
	adminPermission := types.Permission{
		Base: types.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
		Name:        "admin",
		Description: "Permission for administrative access",
	}

	if err := db.Create(&[]types.Permission{
		readPermission, writePermission, deletePermission, adminPermission,
	}).Error; err != nil {
		return err
	}

	rolePermissions := []types.RolePermission{
		{RoleID: adminRole.ID, PermissionID: readPermission.ID},
		{RoleID: adminRole.ID, PermissionID: writePermission.ID},
		{RoleID: adminRole.ID, PermissionID: deletePermission.ID},
		{RoleID: adminRole.ID, PermissionID: adminPermission.ID},
		{RoleID: userRole.ID, PermissionID: readPermission.ID},
	}

	if err := db.Create(&rolePermissions).Error; err != nil {
		return err
	}

	adminUser := types.User{
		Base: types.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
		FullName: "Admin User",
		Email:    "admin@admin.com",
		Password: "$2a$10$UfwNso7PvDUOWAgkMrhkMe1fi16zYHoJIZ/HvWKURQdKWzzL.Xh8G",
		IsActive: true,
	}
	normalUser := types.User{
		Base: types.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
		FullName: "Normal User",
		Email:    "user@user.com",
		Password: "$2a$10$UfwNso7PvDUOWAgkMrhkMe1fi16zYHoJIZ/HvWKURQdKWzzL.Xh8G",
		IsActive: true,
	}

	if err := db.Create(&[]types.User{adminUser, normalUser}).Error; err != nil {
		return err
	}

	userRoles := []types.UserRole{
		{UserID: adminUser.ID, RoleID: adminRole.ID},
		{UserID: normalUser.ID, RoleID: userRole.ID},
	}

	if err := db.Create(&userRoles).Error; err != nil {
		return err
	}

	log.Println("Seed: initial data inserted successfully.")
	return nil
}
