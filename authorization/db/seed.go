package db

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
	"hospital-system/authorization/models"
)

var attendingPermissionsMap = map[string]string{
	"ADMISSIONS":                     "WRITE",
	"PATIENTS:HISTORY":               "READ",
	"PATIENTS:INFO":                  "WRITE",
	"PATIENTS:VITALS":                "WRITE",
	"PATIENTS:DIAGNOSIS":             "WRITE",
	"PATIENTS:CONSULTS":              "WRITE",
	"PATIENTS:TRANSFER":              "WRITE",
	"PATIENTS:DISCHARGE":             "WRITE",
	"PATIENTS:MEDICATIONS:PRESCRIBE": "WRITE",
	"PATIENTS:MEDICATIONS:GIVE":      "WRITE",
	"PATIENTS:LABS:ORDER":            "WRITE",
	"PATIENTS:LABS:RESULT":           "READ",
	"PATIENTS:IMAGING:ORDER":         "WRITE",
	"PATIENTS:IMAGING:RESULT":        "READ",
	"PATIENTS:LOGS":                  "READ",
	"PATIENTS:TEAM":                  "WRITE",
}

var residentPermissionsMap = map[string]string{
	"ADMISSIONS":                "READ",
	"PATIENTS:INFO":             "READ",
	"PATIENTS:VITALS":           "WRITE",
	"PATIENTS:DIAGNOSIS":        "READ",
	"PATIENTS:MEDICATIONS:GIVE": "WRITE",
	"PATIENTS:LABS:RESULT":      "READ",
	"PATIENTS:IMAGING:RESULT":   "READ",
}

var nursePermissionsMap = map[string]string{
	"ADMISSIONS":                "READ",
	"PATIENTS:INFO":             "READ",
	"PATIENTS:VITALS":           "WRITE",
	"PATIENTS:DIAGNOSIS":        "READ",
	"PATIENTS:MEDICATIONS:GIVE": "WRITE",
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
				Permissions: []byte(`{}`),
			},
		}).Error; err != nil {
			slog.Error(fmt.Sprintf("Failed to create predefined roles: %s", err.Error()))
			return err
		}

		// Add predefined teams
		if err = tx.Create([]models.Team{
			{
				ID:          "GENERAL",
				Name:        "Internal Medicine",
				Permissions: []byte(`{"PATIENTS": "READ"}`),
			},
			{
				ID:          "CARDIO",
				Name:        "Cardiology",
				Permissions: []byte(`{"PATIENTS": "READ"}`),
			},
			{
				ID:          "NEURO",
				Name:        "Neurology",
				Permissions: []byte(`{"PATIENTS": "READ"}`),
			},
			{
				ID:          "ORTHO",
				Name:        "Orthopedics",
				Permissions: []byte(`{"PATIENTS": "READ"}`),
			},
			{
				ID:          "OB-GYN",
				Name:        "Obstetrics and Gynecology",
				Permissions: []byte(`{"PATIENTS": "READ"}`),
			},
			{
				ID:          "PEDS",
				Name:        "Pediatrics",
				Permissions: []byte(`{"PATIENTS": "READ"}`),
			},
			{
				ID:          "ONCOLOGY",
				Name:        "Oncology",
				Permissions: []byte(`{"PATIENTS": "READ"}`),
			},
			{
				ID:          "PSYCH",
				Name:        "Psychiatry",
				Permissions: []byte(`{"PATIENTS": "READ"}`),
			},
			{
				ID:          "UROLOGY",
				Name:        "Urology",
				Permissions: []byte(`{"PATIENTS": "READ"}`),
			},
			{
				ID:          "ER",
				Name:        "Emergency Medicine",
				Permissions: []byte(`{"INTAKE": "WRITE", "PATIENTS": "READ"}`),
			},
			{
				ID:          "LAB",
				Name:        "Laboratory",
				Permissions: []byte(`{"LABS": "WRITE"}`),
			},
			{
				ID:          "RADIOLOGY",
				Name:        "Radiology",
				Permissions: []byte(`{"IMAGING": "WRITE"}`),
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
