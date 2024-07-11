package db

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
	"hospital-system/authorization/models"
)

var attendingPermissionsMap = map[string]string{
	"PATIENTS:INFO":               "WRITE",
	"PATIENTS:VITALS":             "WRITE",
	"PATIENTS:DIAGNOSIS":          "WRITE",
	"PATIENTS:TRANSFER":           "WRITE",
	"PATIENTS:DISCHARGE":          "WRITE",
	"PATIENTS:MEDICINE:PRESCRIBE": "WRITE",
	"PATIENTS:MEDICINE:GIVE":      "WRITE",
	"PATIENTS:LABS:ORDER":         "WRITE",
	"PATIENTS:LABS:RESULT":        "WRITE",
	"PATIENTS:IMAGING:ORDER":      "WRITE",
	"PATIENTS:IMAGING:RESULT":     "WRITE",
}

var residentPermissionsMap = map[string]string{
	"PATIENTS:INFO":           "READ",
	"PATIENTS:VITALS":         "WRITE",
	"PATIENTS:DIAGNOSIS":      "READ",
	"PATIENTS:MEDICINE:GIVE":  "WRITE",
	"PATIENTS:LABS:RESULT":    "READ",
	"PATIENTS:IMAGING:RESULT": "READ",
}

var nursePermissionsMap = map[string]string{
	"PATIENTS:INFO":          "READ",
	"PATIENTS:VITALS":        "WRITE",
	"PATIENTS:DIAGNOSIS":     "READ",
	"PATIENTS:MEDICINE:GIVE": "WRITE",
}

var technicianPermissionsMap = map[string]string{
	"PATIENTS:INFO":      "READ",
	"PATIENTS:DIAGNOSIS": "READ",
	"PATIENTS:LABS":      "WRITE",
	"PATIENTS:IMAGING":   "WRITE",
}

func SeedDatabase() error {
	if !(DB.Migrator().HasTable(&models.Role{}) && DB.Migrator().HasTable(&models.Team{})) {
		slog.Error("Database tables are not created, execute database reinit first")
		return fmt.Errorf("database is not initialized")
	}

	attendingPermissionsSerialized, err := json.Marshal(attendingPermissionsMap)
	if err != nil {
		slog.Error("Failed to serialize attending permissions")
		return err
	}

	residentPermissionsSerialized, err := json.Marshal(residentPermissionsMap)
	if err != nil {
		slog.Error("Failed to serialize resident permissions")
		return err
	}

	nursePermissionsSerialized, err := json.Marshal(nursePermissionsMap)
	if err != nil {
		slog.Error("Failed to serialize nurse permissions")
		return err
	}

	technicianPermissionsSerialized, err := json.Marshal(technicianPermissionsMap)
	if err != nil {
		slog.Error("Failed to serialize technician permissions")
		return err
	}

	if err = DB.Transaction(func(tx *gorm.DB) error {
		// Add predefined roles
		if err = tx.Create([]models.Role{
			{
				ID:          "ATTENDING",
				Name:        "Attending Physician",
				Permissions: attendingPermissionsSerialized,
			},
			{
				ID:          "RESIDENT",
				Name:        "Resident Doctor",
				Permissions: residentPermissionsSerialized,
			},
			{
				ID:          "NURSE",
				Name:        "Nurse",
				Permissions: nursePermissionsSerialized,
			},
			{
				ID:          "TECHNICIAN",
				Name:        "Technician",
				Permissions: technicianPermissionsSerialized,
			},
		}).Error; err != nil {
			slog.Error(fmt.Sprintf("Failed to create predefined roles: %s", err.Error()))
			return err
		}

		// Add predefined teams
		if err = tx.Create([]models.Team{
			{
				ID:   "GENERAL",
				Name: "Internal Medicine",
			},
			{
				ID:   "CARDIO",
				Name: "Cardiology",
			},
			{
				ID:   "NEURO",
				Name: "Neurology",
			},
			{
				ID:   "ORTHO",
				Name: "Orthopedics",
			},
			{
				ID:   "OB-GYN",
				Name: "Obstetrics and Gynecology",
			},
			{
				ID:   "PEDS",
				Name: "Pediatrics",
			},
			{
				ID:   "ONCOLOGY",
				Name: "Oncology",
			},
			{
				ID:   "PSYCH",
				Name: "Psychiatry",
			},
			{
				ID:   "UROLOGY",
				Name: "Urology",
			},
			{
				ID:   "RADIOLOGY",
				Name: "Radiology",
			},
		}).Error; err != nil {
			slog.Error(fmt.Sprintf("Failed to create predefined teams: %s", err.Error()))
			return err
		}

		return nil
	}); err != nil {
		slog.Error("Failed to seed the database")
		return err
	}

	return nil
}
