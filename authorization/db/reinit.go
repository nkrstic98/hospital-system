package db

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
	"hospital-system/authorization/models"
)

func ReinitDatabase() error {
	err := DB.Migrator().DropTable(
		&models.Actor{},
		&models.Role{},
		&models.Team{},
		&models.Resource{},
	)
	if err != nil {
		return err
	}

	if err = DB.AutoMigrate(
		&models.Actor{},
		&models.Role{},
		&models.Team{},
		&models.Resource{},
	); err != nil {
		slog.Error(fmt.Sprintf("Failed to reinit database, error %s", err.Error()))
		return err
	}

	err = DB.Transaction(func(tx *gorm.DB) error {
		// Per default, add ADMIN Role
		adminRole := models.Role{
			ID:          "ADMIN",
			Name:        "Administrator",
			Permissions: []byte(`{"EMPLOYEES":"WRITE", "PATIENTS":"WRITE"}`),
		}
		result := tx.Create(&adminRole)
		if result.Error != nil {
			slog.Error(fmt.Sprintf("Failed to create admin role: %s", result.Error.Error()))
			return result.Error
		}

		// Per default, create ADMIN User
		result = tx.Create(&models.Actor{
			ID:     uuid.MustParse("a9c76cb4-7e7e-4bc1-9562-4da7eb5128ce"),
			RoleID: adminRole.ID,
		})
		if result.Error != nil {
			slog.Error(fmt.Sprintf("Failed to create admin user: %s", result.Error.Error()))
			return result.Error
		}

		return nil
	})
	if err != nil {
		return err
	}

	slog.Info("👍 Database reinit completed successfully!")
	return nil
}
