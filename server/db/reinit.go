package db

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
	"hospital-system/server/models"
	"hospital-system/server/utils"
)

func ReinitDatabase() error {
	err := DB.Migrator().DropTable(
		&models.User{},
		&models.Patient{},
		&models.Admission{},
		&models.Lab{},
	)
	if err != nil {
		return err
	}

	if err = DB.AutoMigrate(
		&models.User{},
		&models.Patient{},
		&models.Admission{},
		&models.Lab{},
	); err != nil {
		slog.Error(fmt.Sprintf("Failed to reinit database, error %s", err.Error()))
		return err
	}

	hashedPassword, _ := utils.HashPassword("admin123")

	if err = DB.Transaction(func(tx *gorm.DB) error {
		// Per default, add Admin user
		admin := models.User{
			ID:                           uuid.MustParse("a9c76cb4-7e7e-4bc1-9562-4da7eb5128ce"),
			Firstname:                    "John",
			Lastname:                     "Reynolds",
			NationalIdentificationNumber: "0101990640024",
			Username:                     "admin",
			Email:                        "admin@zmajmc.com",
			Password:                     hashedPassword,
		}
		result := tx.Create(&admin)
		if result.Error != nil {
			slog.Error(fmt.Sprintf("Failed to create admin employee: %s", result.Error.Error()))
			return result.Error
		}

		return nil
	}); err != nil {
		return err
	}

	slog.Info("👍 Database reinit completed successfully!")
	return nil
}
